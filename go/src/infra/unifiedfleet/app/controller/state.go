// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package controller

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
	"go.chromium.org/luci/common/logging"
	crimson "go.chromium.org/luci/machine-db/api/crimson/v1"

	ufspb "infra/unifiedfleet/api/v1/models"
	"infra/unifiedfleet/app/model/datastore"
	"infra/unifiedfleet/app/model/state"
	"infra/unifiedfleet/app/util"
)

// ImportStates imports states of UFS resources.
func ImportStates(ctx context.Context, machines []*crimson.Machine, racks []*crimson.Rack, hosts []*crimson.PhysicalHost, vms []*crimson.VM, vlans []*crimson.VLAN, kvms []*crimson.KVM, switches []*crimson.Switch, pageSize int) (*datastore.OpResults, error) {
	states := make([]*ufspb.StateRecord, 0)
	logging.Infof(ctx, "collecting states of machines")
	for _, m := range machines {
		resourceName := util.AddPrefix(util.MachineCollection, m.GetName())
		states = append(states, &ufspb.StateRecord{
			ResourceName: resourceName,
			State:        util.ToState(m.GetState()),
			User:         util.DefaultImporter,
		})
	}
	logging.Infof(ctx, "collecting states of racks")
	for _, m := range racks {
		resourceName := util.AddPrefix(util.RackCollection, m.GetName())
		states = append(states, &ufspb.StateRecord{
			ResourceName: resourceName,
			State:        util.ToState(m.GetState()),
			User:         util.DefaultImporter,
		})
	}
	logging.Infof(ctx, "collecting states of hosts")
	for _, m := range hosts {
		resourceName := util.AddPrefix(util.HostCollection, m.GetName())
		states = append(states, &ufspb.StateRecord{
			ResourceName: resourceName,
			State:        util.ToState(m.GetState()),
			User:         util.DefaultImporter,
		})
	}
	logging.Infof(ctx, "collecting states of vms")
	for _, vm := range vms {
		resourceName := util.AddPrefix(util.VMCollection, vm.GetName())
		states = append(states, &ufspb.StateRecord{
			ResourceName: resourceName,
			State:        util.ToState(vm.GetState()),
			User:         util.DefaultImporter,
		})
	}
	logging.Infof(ctx, "collecting states of vlans")
	for _, vlan := range vlans {
		resourceName := util.GetBrowserLabName(util.Int64ToStr(vlan.GetId()))
		states = append(states, &ufspb.StateRecord{
			ResourceName: resourceName,
			State:        util.ToState(vlan.GetState()),
			User:         util.DefaultImporter,
		})
	}
	logging.Infof(ctx, "collecting states of kvms")
	for _, kvm := range kvms {
		resourceName := util.AddPrefix(util.KVMCollection, kvm.GetName())
		states = append(states, &ufspb.StateRecord{
			ResourceName: resourceName,
			State:        util.ToState(kvm.GetState()),
			User:         util.DefaultImporter,
		})
	}
	logging.Infof(ctx, "collecting states of switches")
	for _, sw := range switches {
		resourceName := util.AddPrefix(util.SwitchCollection, sw.GetName())
		states = append(states, &ufspb.StateRecord{
			ResourceName: resourceName,
			State:        util.ToState(sw.GetState()),
			User:         util.DefaultImporter,
		})
	}

	deleteNonExistingStates(ctx, states, pageSize)
	allRes := make(datastore.OpResults, 0)
	logging.Infof(ctx, "Importing %d states", len(states))
	for i := 0; ; i += pageSize {
		end := util.Min(i+pageSize, len(states))
		res, err := state.ImportStateRecords(ctx, states[i:end])
		allRes = append(allRes, *res...)
		if err != nil {
			return &allRes, err
		}
		if i+pageSize >= len(states) {
			break
		}
	}
	return &allRes, nil
}

func deleteNonExistingStates(ctx context.Context, states []*ufspb.StateRecord, pageSize int) (*datastore.OpResults, error) {
	resMap := make(map[string]bool)
	for _, r := range states {
		resMap[r.GetResourceName()] = true
	}
	resp, err := state.GetAllStates(ctx)
	if err != nil {
		return nil, err
	}
	var toDelete []string
	for _, sr := range resp.Passed() {
		s := sr.Data.(*ufspb.StateRecord)
		// Skip deleting os hosts' state
		if strings.HasPrefix(s.GetResourceName(), "hosts/chromeos") {
			continue
		}
		if _, ok := resMap[s.GetResourceName()]; !ok {
			toDelete = append(toDelete, s.GetResourceName())
		}
	}
	logging.Infof(ctx, "Deleting %d non-existing states", len(toDelete))
	return deleteByPage(ctx, toDelete, pageSize, state.DeleteStates), nil
}

// UpdateState updates state record for a resource.
func UpdateState(ctx context.Context, stateRecord *ufspb.StateRecord) (*ufspb.StateRecord, error) {

	// TODO(eshwarn): Remove this code block - once all state data is migrated to os namespace
	// First - update in os namespace
	newCtx, err := util.SetupDatastoreNamespace(ctx, util.OSNamespace)
	if err != nil {
		logging.Debugf(ctx, "UpdateState - Failed to set os namespace in context", err)
	} else {
		// Update in os namespace
		sros := proto.Clone(stateRecord).(*ufspb.StateRecord)
		_, err = state.UpdateStateRecord(newCtx, sros)
		if err != nil {
			logging.Errorf(ctx, "UpdateState - Failed to update in os namespace", err)
		}
	}

	// Second - update in default namespace
	newCtx, err = util.SetupDatastoreNamespace(ctx, "")
	if err != nil {
		logging.Debugf(ctx, "UpdateState - Failed to set default namespace in context", err)
	} else {
		// Update in default namespace
		srdefault := proto.Clone(stateRecord).(*ufspb.StateRecord)
		_, err = state.UpdateStateRecord(newCtx, srdefault)
		if err != nil {
			logging.Errorf(ctx, "UpdateState - Failed to update in default namespace", err)
		}
	}
	// TODO(eshwarn): Remove this code block - once all state data is migrated to os namespace

	// Third - Update in whatever namespace provided by client and return
	return state.UpdateStateRecord(ctx, stateRecord)
}

// GetState returns state record for a resource.
func GetState(ctx context.Context, resourceName string) (*ufspb.StateRecord, error) {
	// First try to find in os namespace, if not find in default namespace
	// TODO(eshwarn): Remove this - once all state data is migrated to os namespace
	newCtx, err := util.SetupDatastoreNamespace(ctx, util.OSNamespace)
	if err != nil {
		logging.Debugf(ctx, "GetState - Failed to set os namespace in context", err)
		return state.GetStateRecord(ctx, resourceName)
	}
	record, err := state.GetStateRecord(newCtx, resourceName)
	if err == nil {
		return record, err
	}

	// default namespace
	newCtx, err = util.SetupDatastoreNamespace(ctx, "")
	if err != nil {
		logging.Debugf(ctx, "GetState - Failed to set default namespace in context", err)
		return state.GetStateRecord(ctx, resourceName)
	}

	return state.GetStateRecord(newCtx, resourceName)
}
