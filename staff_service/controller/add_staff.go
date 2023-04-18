package controller

import (
	"context"
	accountSchema "store-bpel/account_service/schema"
	"store-bpel/staff_service/repository"
	"store-bpel/staff_service/schema"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func (s *staffServiceController) AddStaff(ctx context.Context, request *schema.AddStaffRequest) error {
	staffId := strings.Split(request.Email, "@")[0]
	staffModel := &repository.StaffModel{
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
		Email:         request.Email,
		Status:        "APPROVED",
		BranchId:      request.WorkingPlace,
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		err := s.repository.AddStaff(ctx, staffModel)
		if err != nil {
			return err
		}
		err = s.repository.CreateAccount(ctx, &repository.AccountModel{
			Username: staffId,
			StaffId:  staffId,
		})

		role, err := strconv.Atoi(request.Role)
		if err != nil {

			return err
		}

		s.accountAdapter.CreateAccount(ctx, &accountSchema.AddAccountRequest{
			Username: staffId,
			Password: "123456",
			Role:     role,
		})
		// TODO call to branch service to add new staff and current working place

		return nil
	})
}
