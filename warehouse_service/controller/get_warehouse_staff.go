package controller

import (
	"context"
	"store-bpel/staff_service/schema"
)

func (s *warehouseServiceController) GetWarehouseStaff(ctx context.Context) (*schema.GetStaffResponse, error) {
	staffs, err := s.staffAdapter.GetStaff(ctx)
	if err != nil {
		return nil, err
	}
	return staffs, nil
}
