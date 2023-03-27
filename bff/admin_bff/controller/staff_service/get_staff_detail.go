package staff_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/staff_service"
)

func (c *staffBffController) GetStaffDetail(ctx context.Context, request *staff_service.GetStaffRequest) (*staff_service.GetStaffResponseData, error) {
	staffs, err := c.staffAdapter.GetStaffDetail(ctx, request.StaffId)
	if err != nil {
		return nil, err
	}
	data := staffs[0]

	return &staff_service.GetStaffResponseData{
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
	}, nil
}
