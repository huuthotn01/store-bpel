package controller

import (
	"context"
	"store-bpel/staff_service/repository"
	"store-bpel/staff_service/schema"
)

func (s *staffServiceController) AddStaff(ctx context.Context, request *schema.AddStaffRequest) error {
	staffModel := &repository.StaffModel{
		StaffId:   request.Username,
		StaffName: request.Name,
		Province:  request.Province,
		District:  request.District,
		Ward:      request.Ward,
		Street:    request.Street,
		CitizenId: request.CitizenId,
	}
	err := s.repository.AddStaff(ctx, staffModel)
	return err
}
