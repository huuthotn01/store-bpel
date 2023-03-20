package controller

import (
	"context"
	account_schema "store-bpel/account_service/schema"
	"store-bpel/staff_service/repository"
	"store-bpel/staff_service/schema"
	"strings"

	"gorm.io/gorm"
)

var mappingRoleStringRoleInt = map[string]int{
	"client":                  1,
	"admin - quản trị viên": 2,
	"quản lý chi nhánh":    3,
	"quản lý kho":           4,
	"quản lý hàng":         5,
	"trưởng chi nhánh":      6,
}

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
			Username: request.Email,
			StaffId:  staffId,
		})
		// TODO call to branch service to add new staff and current working place

		// TODO apply async to call to account service to create account
		return s.accountAdapter.CreateAccount(ctx, &account_schema.AddAccountRequest{
			Username: staffId,
			Role:     mappingRoleStringRoleInt[request.Role],
		})
	})
}
