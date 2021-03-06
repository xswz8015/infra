// Code generated by svcdec; DO NOT EDIT.

package fleet

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedInventory struct {
	// Service is the service to decorate.
	Service InventoryServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(ctx context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(ctx context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedInventory) DeployDut(ctx context.Context, req *DeployDutRequest) (rsp *DeployDutResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeployDut", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeployDut(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeployDut", rsp, err)
	}
	return
}

func (s *DecoratedInventory) RedeployDut(ctx context.Context, req *RedeployDutRequest) (rsp *RedeployDutResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "RedeployDut", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.RedeployDut(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "RedeployDut", rsp, err)
	}
	return
}

func (s *DecoratedInventory) GetDeploymentStatus(ctx context.Context, req *GetDeploymentStatusRequest) (rsp *GetDeploymentStatusResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetDeploymentStatus", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetDeploymentStatus(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetDeploymentStatus", rsp, err)
	}
	return
}

func (s *DecoratedInventory) DeleteDuts(ctx context.Context, req *DeleteDutsRequest) (rsp *DeleteDutsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteDuts", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteDuts(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteDuts", rsp, err)
	}
	return
}

func (s *DecoratedInventory) BalancePools(ctx context.Context, req *BalancePoolsRequest) (rsp *BalancePoolsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "BalancePools", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.BalancePools(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "BalancePools", rsp, err)
	}
	return
}

func (s *DecoratedInventory) ResizePool(ctx context.Context, req *ResizePoolRequest) (rsp *ResizePoolResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ResizePool", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ResizePool(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ResizePool", rsp, err)
	}
	return
}

func (s *DecoratedInventory) RemoveDutsFromDrones(ctx context.Context, req *RemoveDutsFromDronesRequest) (rsp *RemoveDutsFromDronesResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "RemoveDutsFromDrones", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.RemoveDutsFromDrones(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "RemoveDutsFromDrones", rsp, err)
	}
	return
}

func (s *DecoratedInventory) AssignDutsToDrones(ctx context.Context, req *AssignDutsToDronesRequest) (rsp *AssignDutsToDronesResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "AssignDutsToDrones", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.AssignDutsToDrones(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "AssignDutsToDrones", rsp, err)
	}
	return
}

func (s *DecoratedInventory) ListServers(ctx context.Context, req *ListServersRequest) (rsp *ListServersResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListServers", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListServers(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListServers", rsp, err)
	}
	return
}

func (s *DecoratedInventory) GetDutInfo(ctx context.Context, req *GetDutInfoRequest) (rsp *GetDutInfoResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetDutInfo", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetDutInfo(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetDutInfo", rsp, err)
	}
	return
}

func (s *DecoratedInventory) GetDroneConfig(ctx context.Context, req *GetDroneConfigRequest) (rsp *GetDroneConfigResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetDroneConfig", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetDroneConfig(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetDroneConfig", rsp, err)
	}
	return
}

func (s *DecoratedInventory) ListRemovedDuts(ctx context.Context, req *ListRemovedDutsRequest) (rsp *ListRemovedDutsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListRemovedDuts", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListRemovedDuts(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListRemovedDuts", rsp, err)
	}
	return
}

func (s *DecoratedInventory) PushInventoryToQueen(ctx context.Context, req *PushInventoryToQueenRequest) (rsp *PushInventoryToQueenResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "PushInventoryToQueen", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.PushInventoryToQueen(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "PushInventoryToQueen", rsp, err)
	}
	return
}

func (s *DecoratedInventory) UpdateDutLabels(ctx context.Context, req *UpdateDutLabelsRequest) (rsp *UpdateDutLabelsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateDutLabels", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateDutLabels(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateDutLabels", rsp, err)
	}
	return
}

func (s *DecoratedInventory) BatchUpdateDuts(ctx context.Context, req *BatchUpdateDutsRequest) (rsp *BatchUpdateDutsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "BatchUpdateDuts", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.BatchUpdateDuts(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "BatchUpdateDuts", rsp, err)
	}
	return
}

func (s *DecoratedInventory) UpdateCachedInventory(ctx context.Context, req *UpdateCachedInventoryRequest) (rsp *UpdateCachedInventoryResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateCachedInventory", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateCachedInventory(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateCachedInventory", rsp, err)
	}
	return
}

func (s *DecoratedInventory) UpdateDeviceConfig(ctx context.Context, req *UpdateDeviceConfigRequest) (rsp *UpdateDeviceConfigResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateDeviceConfig", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateDeviceConfig(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateDeviceConfig", rsp, err)
	}
	return
}

func (s *DecoratedInventory) GetStableVersion(ctx context.Context, req *GetStableVersionRequest) (rsp *GetStableVersionResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetStableVersion", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetStableVersion(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetStableVersion", rsp, err)
	}
	return
}

func (s *DecoratedInventory) SetSatlabStableVersion(ctx context.Context, req *SetSatlabStableVersionRequest) (rsp *SetSatlabStableVersionResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "SetSatlabStableVersion", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.SetSatlabStableVersion(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "SetSatlabStableVersion", rsp, err)
	}
	return
}

func (s *DecoratedInventory) DeleteSatlabStableVersion(ctx context.Context, req *DeleteSatlabStableVersionRequest) (rsp *DeleteSatlabStableVersionResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeleteSatlabStableVersion", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteSatlabStableVersion(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeleteSatlabStableVersion", rsp, err)
	}
	return
}

func (s *DecoratedInventory) DumpStableVersionToDatastore(ctx context.Context, req *DumpStableVersionToDatastoreRequest) (rsp *DumpStableVersionToDatastoreResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DumpStableVersionToDatastore", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DumpStableVersionToDatastore(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DumpStableVersionToDatastore", rsp, err)
	}
	return
}

func (s *DecoratedInventory) ReportInventory(ctx context.Context, req *ReportInventoryRequest) (rsp *ReportInventoryResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ReportInventory", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ReportInventory(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ReportInventory", rsp, err)
	}
	return
}

func (s *DecoratedInventory) UpdateManufacturingConfig(ctx context.Context, req *UpdateManufacturingConfigRequest) (rsp *UpdateManufacturingConfigResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateManufacturingConfig", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateManufacturingConfig(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateManufacturingConfig", rsp, err)
	}
	return
}
