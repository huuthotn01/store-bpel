package controller

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"store-bpel/staff_service/internal/repository"
	"store-bpel/staff_service/schema"
	"strings"
	"time"
)

func (s *staffServiceController) CreateAddRequest(ctx context.Context, request *schema.CreateAddRequest) error {
	/*
		When receiving CreateAddRequest:
		1. Create request in requests' table, with request_type = 'ADD' and status = 'PENDING'
		2. Create staff in staff table with staff's status = 'PENDING'
		3. Not create account yet.
	*/
	requestId := fmt.Sprintf("add_%s", cast.ToString(time.Now().Unix()))
	staffId := strings.Split(request.Email, "@")[0]
	return s.db.Transaction(func(tx *gorm.DB) error {
		// create add request
		err := s.repository.CreateStaffRequest(ctx, &repository.RequestsModel{
			Id:          requestId,
			RequestDate: time.Now(),
			RequestType: "ADD",
			StaffId:     staffId,
			Status:      "PENDING",
		})
		if err != nil {
			return err
		}
		// create staff in staff table
		return s.repository.AddStaff(ctx, &repository.StaffModel{
			StaffId:       staffId,
			StaffName:     request.Name,
			Province:      request.Province,
			District:      request.District,
			Ward:          request.Ward,
			Street:        request.Street,
			Hometown:      request.Hometown,
			CitizenId:     request.CitizenId,
			Birthdate:     request.Birthdate,
			Phone:         request.Phone,
			StaffPosition: request.Role,
			Gender:        request.Gender,
			Salary:        request.Salary,
			Email:         request.Email,
			Status:        "PENDING",
			BranchId: request.WorkingPlace,
		})
	})
}
