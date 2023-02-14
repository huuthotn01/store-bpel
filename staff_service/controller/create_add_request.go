package controller

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"store-bpel/staff_service/repository"
	"store-bpel/staff_service/schema"
	"time"
)

func (s *staffServiceController) CreateAddRequest(ctx context.Context, request *schema.CreateAddRequest) error {
	requestId := fmt.Sprintf("add_%s", cast.ToString(time.Now().Unix()))
	staffId := fmt.Sprintf("staff_%s", cast.ToString(time.Now().Unix()))
	return s.repository.CreateAddRequest(ctx, &repository.StaffModel{
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
	}, &repository.RequestsModel{
		Id:          requestId,
		RequestDate: time.Now(),
		// RequestType: ,
		StaffId: staffId,
		Status:  "PENDING",
	}, request.Email)
}
