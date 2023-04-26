package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
)

func (c *warehouseBffController) GetManager(ctx context.Context, request *warehouse_service.GetWarehouseId) (*warehouse_service.GetWarehouseManagerResponseData, error) {
	manager, err := c.warehouseAdapter.GetManager(ctx, request.WarehouseId)
	if err != nil {
		return nil, err
	}

	return &warehouse_service.GetWarehouseManagerResponseData{
		StaffId:     manager.StaffId,
		StaffName:   manager.StaffName,
		Street:      manager.Street,
		Ward:        manager.Ward,
		District:    manager.District,
		Province:    manager.Province,
		CitizenId:   manager.CitizenId,
		BranchId:    manager.BranchId,
		Hometown:    manager.Hometown,
		Salary:      manager.Salary,
		Birthdate:   manager.Birthdate,
		StartDate:   manager.StartDate,
		Gender:      manager.Gender,
		PhoneNumber: manager.PhoneNumber,
		Status:      manager.Status,
		Email:       manager.Email,
	}, nil
}
