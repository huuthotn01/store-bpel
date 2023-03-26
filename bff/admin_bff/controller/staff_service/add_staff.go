package staff_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/staff_service"
	"store-bpel/staff_service/schema"
)

func (c *staffBffController) AddStaff(ctx context.Context, request *staff_service.AddStaffRequest) error {
	return c.staffAdapter.AddStaff(ctx, &schema.AddStaffRequest{
		Name:         request.Name,
		Birthdate:    request.Birthdate,
		Hometown:     request.Hometown,
		CitizenId:    request.CitizenId,
		Phone:        request.Phone,
		Street:       request.Street,
		Ward:         request.Ward,
		District:     request.District,
		Province:     request.Province,
		WorkingPlace: request.WorkingPlace,
		Role:         request.Role,
		Gender:       request.Gender,
		Salary:       request.Salary,
		Email:        request.Email,
	})
}
