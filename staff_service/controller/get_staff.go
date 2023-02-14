package controller

import (
	"context"
	"github.com/spf13/cast"
	"store-bpel/staff_service/schema"
)

func (s *staffServiceController) GetStaff(ctx context.Context) ([]*schema.GetStaffResponseData, error) {
	staffs, err := s.repository.GetStaff(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*schema.GetStaffResponseData, 0, len(staffs))
	for _, staff := range staffs {
		res = append(res, &schema.GetStaffResponseData{
			StaffId:     staff.StaffId,
			StaffName:   staff.StaffName,
			Street:      staff.Street,
			Ward:        staff.Ward,
			District:    staff.District,
			Province:    staff.Province,
			CitizenId:   staff.CitizenId,
			Role:        cast.ToString(staff.StaffPosition),
			PhoneNumber: staff.Phone,
			Email:       staff.Email,
			StartDate:   staff.StartDate,
			Salary:      staff.Salary,
		})
	}
	return res, nil
}

func (s *staffServiceController) GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error) {
	return nil, nil
}
