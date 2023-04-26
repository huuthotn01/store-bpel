package controller

import (
	"context"
	"store-bpel/warehouse_service/schema"
)

func (c *warehouseServiceController) GetWarehouseStaff(ctx context.Context, warehouseId string) ([]*schema.GetWarehouseStaffResponseData, error) {
	staffInWh, err := c.repository.GetStaffByWarehouse(ctx, warehouseId)
	if err != nil {
		return nil, err
	}

	respStaffs := make([]*schema.GetWarehouseStaffResponseData, 0, len(staffInWh))

	for _, staff := range staffInWh {
		data, err := c.staffAdapter.GetDetailStaff(ctx, staff.StaffCode)
		if err != nil {
			return nil, err
		}

		respStaffs = append(respStaffs, &schema.GetWarehouseStaffResponseData{
			StaffId:     data.StaffId,
			StaffName:   data.StaffName,
			Street:      data.Street,
			Ward:        data.Ward,
			District:    data.District,
			Province:    data.Province,
			CitizenId:   data.CitizenId,
			Role:        data.Role,
			BranchId:    data.BranchId,
			Hometown:    data.Hometown,
			Salary:      data.Salary,
			Birthdate:   data.Birthdate,
			StartDate:   data.StartDate,
			Gender:      data.Gender,
			PhoneNumber: data.PhoneNumber,
			Status:      data.Status,
			Email:       data.Email,
		})
	}

	return respStaffs, nil
}
