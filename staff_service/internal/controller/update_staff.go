package controller

import (
	"context"
	accountSchema "store-bpel/account_service/schema"
	"store-bpel/staff_service/internal/repository"
	"store-bpel/staff_service/schema"
	"strconv"
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

	role, err := strconv.Atoi(request.Role)
	if err == nil {
		err = s.accountAdapter.UpdateRole(ctx, staffId, &accountSchema.UpdateRoleRequest{
			Role: role,
		})
		return err
	}

	if request.WorkingPlace != "" {
		// TODO update to branch service
	}
	return nil
}
