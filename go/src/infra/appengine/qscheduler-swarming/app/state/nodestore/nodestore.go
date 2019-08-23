// Copyright 2019 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package nodestore implements a datastore-backed persistent store of qscheduler
// state, that shards state over as many entities as necessary to stay under datastore's
// single-entity size limit, and uses an in-memory cache to avoid unnecessary
// datastore reads.
package nodestore

import (
	"bytes"
	"context"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/luci/common/errors"

	"infra/appengine/qscheduler-swarming/app/state/nodestore/internal/blob"
	"infra/appengine/qscheduler-swarming/app/state/types"
	"infra/qscheduler/qslib/reconciler"
	"infra/qscheduler/qslib/scheduler"
)

var errWrongGeneration = errors.New("wrong generation")

type stateAndGeneration struct {
	state      *blob.QSchedulerPoolState
	generation int64
}

// Operator describes an interface that NodeStore will use for for mutating a
// quotascheduler's state, and persisting any side-effects.
//
// NodeStore will not call any methods of an Operator concurrently.
type Operator interface {
	// Modify is called to modify a quotascheduler state; it may be called
	// more than once, therefore it should not have side effects besides:
	// any side effects other than:
	// - modifying the supplied *types.QScheduler,
	// - side effects that are stored internally to the Operator (e.g. metrics
	// to be used in the Commit or Finish calls).
	//
	// If there are any side effects stored internally to the Operator, they
	// should be reset on each call to Modify.
	Modify(context.Context, *types.QScheduler) error

	// Commit will be called within a datastore transaction, after a successful
	// call to Modify. Commit should be used to persist any transactional
	// side effects of Modify (such as emitting tasks to a task queue).
	Commit(context.Context) error

	// Finish will be called at most once, after a successful call to Commit.
	// This will be called outside of any transactions, and should be used
	// for non-transactional at-most-once side effects, such as incrementing
	// ts_mon counters.
	Finish(context.Context)
}

// New returns a new NodeStore.
func New(qsPoolID string) *NodeStore {
	return &NodeStore{qsPoolID: qsPoolID}
}

// NodeStore is a persistent store for an individual quotascheduler state.
//
// All methods are concurrency-safe.
type NodeStore struct {
	qsPoolID string

	cacheLock sync.RWMutex
	cache     *stateAndGeneration
}

// Create creates a new persistent scheduler entity if one doesn't exist.
func (n *NodeStore) Create(ctx context.Context, timestamp time.Time) error {
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		record := &stateRecord{PoolID: n.qsPoolID}
		exists, err := datastore.Exists(ctx, record)
		if err != nil {
			return err
		}
		if exists.Any() {
			return errors.Reason("entity already exists").Err()
		}

		s := scheduler.New(timestamp)
		p := &blob.QSchedulerPoolState{Scheduler: s.ToProto()}
		nodeIDs, err := writeNodes(ctx, p, timestamp)
		if err != nil {
			return err
		}
		record.NodeIDs = nodeIDs
		return datastore.Put(ctx, record)
	}, &datastore.TransactionOptions{XG: true})
	if err != nil {
		return errors.Annotate(err, "nodestore create").Err()
	}
	return nil
}

// Run runs the given operator.
func (n *NodeStore) Run(ctx context.Context, o Operator) error {
	sg := n.getCached()
	// Fast path; use in-memory cache to avoid reading state from datastore.
	if sg != nil {
		err := n.tryRun(ctx, o, sg)
		switch {
		case err == nil:
			o.Finish(ctx)
			return nil
		case errors.Contains(err, errWrongGeneration):
			// In-memory cache was wrong generation; try slow path.
		default:
			return errors.Annotate(err, "nodestore run").Err()
		}
	}

	for i := 0; i < 10; i++ {
		// Slow path; read full state from datastore, then follow usual modification
		// flow.
		sg, err := n.loadState(ctx)
		if err != nil {
			return errors.Annotate(err, "nodestore run").Err()
		}

		err = n.tryRun(ctx, o, sg)
		switch {
		case err == nil:
			o.Finish(ctx)
			return nil
		case errors.Contains(err, errWrongGeneration):
			// Contention against some other writer; try again.
		default:
			return errors.Annotate(err, "nodestore run").Err()
		}
	}

	return errors.New("nodestore run: too many attempts")
}

// Get returns the current qscheduler state.
func (n *NodeStore) Get(ctx context.Context) (*types.QScheduler, error) {
	sg, err := n.loadState(ctx)
	if err != nil {
		return nil, errors.Annotate(err, "nodestore get").Err()
	}

	return &types.QScheduler{
		Reconciler: reconciler.NewFromProto(sg.state.Reconciler),
		Scheduler:  scheduler.NewFromProto(sg.state.Scheduler),
	}, nil
}

// tryRun attempts to modify and commit the given state, using the given operator.
func (n *NodeStore) tryRun(ctx context.Context, o Operator, sg *stateAndGeneration) error {
	timestamp := time.Now()
	q := &types.QScheduler{
		SchedulerID: n.qsPoolID,
		Reconciler:  reconciler.NewFromProto(sg.state.Reconciler),
		Scheduler:   scheduler.NewFromProto(sg.state.Scheduler),
	}
	if err := o.Modify(ctx, q); err != nil {
		return errors.Annotate(err, "nodestore try").Err()
	}
	p := &blob.QSchedulerPoolState{
		Reconciler: q.Reconciler.ToProto(),
		Scheduler:  q.Scheduler.ToProto(),
	}
	IDs, err := writeNodes(ctx, p, timestamp)
	if err != nil {
		return errors.Annotate(err, "nodestore try").Err()
	}

	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		record := &stateRecord{PoolID: n.qsPoolID}
		if err := datastore.Get(ctx, record); err != nil {
			return err
		}

		if record.Generation != sg.generation {
			// Supplied generation was out of date.
			return errWrongGeneration
		}

		outRecord := &stateRecord{
			PoolID:     n.qsPoolID,
			NodeIDs:    IDs,
			Generation: sg.generation + 1,
		}
		if err := datastore.Put(ctx, outRecord); err != nil {
			return err
		}

		return o.Commit(ctx)
	}, nil)

	if err != nil {
		return errors.Annotate(err, "nodestore try").Err()
	}

	n.setCached(&stateAndGeneration{p, sg.generation + 1})
	return nil
}

func (n *NodeStore) getCached() *stateAndGeneration {
	n.cacheLock.RLock()
	s := n.cache
	n.cacheLock.RUnlock()
	return s
}

func (n *NodeStore) setCached(sg *stateAndGeneration) {
	n.cacheLock.Lock()
	if n.cache == nil || sg.generation > n.cache.generation {
		n.cache = sg
	}
	n.cacheLock.Unlock()
}

func (n *NodeStore) loadState(ctx context.Context) (*stateAndGeneration, error) {
	record := &stateRecord{PoolID: n.qsPoolID}
	if err := datastore.Get(ctx, record); err != nil {
		return nil, errors.Annotate(err, "nodestore load").Err()
	}

	state, err := loadNodes(ctx, record.NodeIDs)
	if err != nil {
		return nil, errors.Annotate(err, "nodestore load").Err()
	}

	return &stateAndGeneration{state: state, generation: record.Generation}, nil
}

type stateRecord struct {
	_kind string `gae:"$kind,stateRecord"`

	// PoolID is the qs pool ID for this record.
	PoolID string `gae:"$id"`

	Generation int64 `gae:",noindex"`

	NodeIDs []string `gae:",noindex"`
}

// stateNode is the datastore entity used to represent a shard of
// quotascheduler state.
//
// TODO(akeshet): Add a cleanup mechanism that removes stale graph entities.
type stateNode struct {
	_kind string `gae:"$kind,stateNode"`

	// ID is a globally unique ID for this entity. Entities are append-only.
	ID string `gae:"$id"`

	// QSchedulerPoolStateDataShard contains this node's shard of the proto-
	// serialized QSchedulerPoolState.
	QSchedulerPoolStateDataShard []byte `gae:",noindex"`
}

// writeNodes writes the given state to as many nodes as necessary, and returns
// their IDs.
func writeNodes(ctx context.Context, state *blob.QSchedulerPoolState, timestamp time.Time) ([]string, error) {
	bytes, err := proto.Marshal(state)
	if err != nil {
		return nil, errors.Annotate(err, "write nodes").Err()
	}

	// TODO(akeshet): Tune this for a good balance between staying safely below
	// upper limit and using fewer shards.
	maxBytes := 900000
	var shards [][]byte
	for offset := 0; offset < len(bytes); offset += maxBytes {
		start := offset
		end := start + maxBytes
		if end > len(bytes) {
			end = len(bytes)
		}
		shards = append(shards, bytes[start:end])
	}

	nodes := make([]interface{}, len(shards))
	IDs := make([]string, len(shards))
	for i, shard := range shards {
		ID := uuid.New().String()
		node := &stateNode{
			ID:                           ID,
			QSchedulerPoolStateDataShard: shard,
		}

		nodes[i] = node
		IDs[i] = ID
	}

	if err := datastore.Put(ctx, nodes...); err != nil {
		return nil, errors.Annotate(err, "write nodes").Err()
	}

	return IDs, nil
}

// loadNodes loads state from the given set of nodes.
func loadNodes(ctx context.Context, nodeIDs []string) (*blob.QSchedulerPoolState, error) {
	nodes := make([]interface{}, len(nodeIDs))
	for i, ID := range nodeIDs {
		nodes[i] = &stateNode{ID: ID}
	}

	if err := datastore.Get(ctx, nodes...); err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	for _, n := range nodes {
		node, ok := n.(*stateNode)
		if !ok {
			return nil, errors.New("load nodes: unexpected node type")
		}

		if _, err := buffer.Write(node.QSchedulerPoolStateDataShard); err != nil {
			return nil, errors.Annotate(err, "load nodes").Err()
		}
	}

	state := blob.QSchedulerPoolState{}
	if err := proto.Unmarshal(buffer.Bytes(), &state); err != nil {
		return nil, err
	}

	return &state, nil
}
