package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/staff_service/repository"
	"store-bpel/staff_service/schema"
	"strings"
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
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		err := s.repository.AddStaff(ctx, staffModel)
		if err != nil {
			return err
		}
		return s.repository.CreateAccount(ctx, &repository.AccountModel{
			Username: request.Email,
			StaffId:  staffId,
		})
		// TODO call to branch service to add new staff and current working place
		// TODO call to account service to create account, use async
	})
}
