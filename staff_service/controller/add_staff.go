package controller

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"store-bpel/staff_service/repository"
	"store-bpel/staff_service/schema"
	"time"
)

func (s *staffServiceController) AddStaff(ctx context.Context, request *schema.AddStaffRequest) error {
	staffId := fmt.Sprintf("staff_%s", cast.ToString(time.Now().Unix()))
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
	}
	err := s.repository.AddStaff(ctx, staffModel, request.Email)
	// TODO call to branch service to add new staff and current working place
	// TODO call to account service to create account, use async
	return err
}
