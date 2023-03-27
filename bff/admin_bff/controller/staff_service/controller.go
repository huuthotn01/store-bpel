package staff_service

import (
	"context"
	"store-bpel/bff/admin_bff/adapter"
	"store-bpel/bff/admin_bff/config"
	"store-bpel/bff/admin_bff/schema/staff_service"
)

type IStaffBffController interface {
	GetStaff(ctx context.Context, request *staff_service.GetStaffRequest) ([]*staff_service.GetStaffResponseData, error)
	GetStaffDetail(ctx context.Context, request *staff_service.GetStaffRequest) (*staff_service.GetStaffResponseData, error)
	GetStaffAttendance(ctx context.Context, request *staff_service.GetStaffAttendanceRequest) ([]*staff_service.GetStaffAttendanceResponseData, error)
	AddStaff(ctx context.Context, request *staff_service.AddStaffRequest) error
	UpdateStaff(ctx context.Context, request *staff_service.UpdateStaffRequest) error
	CreateAddRequest(ctx context.Context, request *staff_service.CreateAddRequest) error
	CreateDeleteRequest(ctx context.Context, request *staff_service.CreateDeleteRequest) error
	UpdateRequestStatus(ctx context.Context, request *staff_service.UpdateRequestStatusRequest) error
	GetRequestList(ctx context.Context) ([]*staff_service.GetRequestResponseData, error)
}

type staffBffController struct {
	cfg          *config.Config
	staffAdapter adapter.IStaffServiceAdapter
}

func NewController(cfg *config.Config) IStaffBffController {
	// init staff adapter
	staffAdapter := adapter.NewStaffAdapter(cfg)

	return &staffBffController{
		cfg:          cfg,
		staffAdapter: staffAdapter,
	}
}
