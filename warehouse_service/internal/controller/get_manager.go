package controller

import (
	"context"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseServiceController) GetWarehouseManager(ctx context.Context, warehouseId string) (*schema.GetWarehouseManagerResponseData, error) {
	whManager, err := c.repository.GetWarehouseManager(ctx, warehouseId)
	if err != nil {
		return nil, err
	}

	staffDetail, err := c.staffAdapter.GetDetailStaff(ctx, whManager.StaffCode)
	if err != nil {
		return nil, err
	}

	return &schema.GetWarehouseManagerResponseData{
		StaffId:     staffDetail.StaffId,
		StaffName:   staffDetail.StaffName,
		Street:      staffDetail.Street,
		Ward:        staffDetail.Ward,
		District:    staffDetail.District,
		Province:    staffDetail.Province,
		CitizenId:   staffDetail.CitizenId,
		BranchId:    staffDetail.BranchId,
		Hometown:    staffDetail.Hometown,
		Salary:      staffDetail.Salary,
		Birthdate:   staffDetail.Birthdate,
		StartDate:   staffDetail.StartDate,
		Gender:      staffDetail.Gender,
		PhoneNumber: staffDetail.PhoneNumber,
		Status:      staffDetail.Status,
		Email:       staffDetail.Email,
	}, nil
}
