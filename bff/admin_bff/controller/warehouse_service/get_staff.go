package warehouse_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/warehouse_service"
)

func (c *warehouseBffController) GetStaff(ctx context.Context, request *warehouse_service.GetWarehouseId) ([]*warehouse_service.GetWarehouseStaffResponseData, error) {
	staffs, err := c.warehouseAdapter.GetStaff(ctx, request.WarehouseId)
	if err != nil {
		return nil, err
	}

	respStaffs := make([]*warehouse_service.GetWarehouseStaffResponseData, 0, len(staffs))
	for _, data := range staffs {
		respStaffs = append(respStaffs, &warehouse_service.GetWarehouseStaffResponseData{
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
