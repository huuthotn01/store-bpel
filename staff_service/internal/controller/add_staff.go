package controller

import (
	"context"
	"encoding/json"
	accountSchema "store-bpel/account_service/schema"
	branchSchema "store-bpel/branch_service/schema"
	"store-bpel/library/kafka_lib"
	"store-bpel/staff_service/internal/repository"
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

		// call account service to add account
		// not provide password, let account service generate
		addAccountRequest := &accountSchema.AddAccountRequest{
			Username: staffId,
			Role:     role,
		}
		addAccReqByte, err := json.Marshal(addAccountRequest)
		if err != nil {
			return err
		}
		err = s.kafkaAdapter.Publish(ctx, kafka_lib.ACCOUNT_SERVICE_TOPIC, addAccReqByte)
		if err != nil {
			return err
		}

		// call branch service to add new staff and current working place
		addBranchStaff := &branchSchema.AddBranchStaffRequest{
			StaffId:  staffId,
			BranchId: request.WorkingPlace,
		}
		addBranchReqByte, err := json.Marshal(addBranchStaff)
		if err != nil {
			return err
		}
		return s.kafkaAdapter.Publish(ctx, kafka_lib.BRANCH_SERVICE_TOPIC, addBranchReqByte)
	})
}
