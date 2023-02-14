package controller

import (
	"context"
	"store-bpel/staff_service/repository"
	"store-bpel/staff_service/schema"
)

func (s *staffServiceController) UpdateStaff(ctx context.Context, request *schema.UpdateStaffRequest, staffId string) error {
	err := s.repository.UpdateStaff(ctx, &repository.StaffModel{
		StaffId:       staffId,
		StaffName:     request.Name,
		Province:      request.Province,
		District:      request.District,
		Ward:          request.Ward,
		Street:        request.Street,
		CitizenId:     request.CitizenId,
		Phone:         request.Phone,
		Birthdate:     request.Birthdate,
		Hometown:      request.Hometown,
		Salary:        request.Salary,
		StaffPosition: request.Role,
		Gender:        request.Gender,
	})
	if err != nil {
		return err
	}
	if request.WorkingPlace != "" {
		// TODO update to branch service
	}
	return nil
}
