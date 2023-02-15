package controller

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"store-bpel/staff_service/repository"
	"store-bpel/staff_service/schema"
	"time"
)

func (s *staffServiceController) CreateAddRequest(ctx context.Context, request *schema.CreateAddRequest) error {
	requestId := fmt.Sprintf("add_%s", cast.ToString(time.Now().Unix()))
	staffId := fmt.Sprintf("staff_%s", cast.ToString(time.Now().Unix()))
	return s.db.Transaction(func(tx *gorm.DB) error {
		// create add request
		err := s.repository.CreateStaffRequest(ctx, &repository.RequestsModel{
			Id:          requestId,
			RequestDate: time.Now(),
			RequestType: 1, // 1 for ADD
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
		}, request.Email)
	})
}
