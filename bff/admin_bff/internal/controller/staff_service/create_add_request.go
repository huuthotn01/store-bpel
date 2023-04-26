package staff_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/staff_service"
	"store-bpel/staff_service/schema"
)

func (c *staffBffController) CreateAddRequest(ctx context.Context, request *staff_service.CreateAddRequest) error {
	return c.staffAdapter.CreateAddRequest(ctx, &schema.CreateAddRequest{
		Name:         request.Name,
		Birthdate:    request.Birthdate,
		Hometown:     request.Hometown,
		CitizenId:    request.CitizenId,
		Phone:        request.Phone,
		Province:     request.Province,
		District:     request.District,
		Ward:         request.Ward,
		Street:       request.Street,
		WorkingPlace: request.WorkingPlace,
		Role:         request.Role,
		Gender:       request.Gender,
		Salary:       request.Salary,
		Email:        request.Email,
	})
}
