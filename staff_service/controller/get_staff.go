package controller

import (
	"context"
	"store-bpel/staff_service/schema"
)

func (s *staffServiceController) GetStaff(ctx context.Context, staffName, staffId string) ([]*schema.GetStaffResponseData, error) {
	staffs, err := s.repository.GetStaff(ctx, staffName, staffId)
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
			Role:        staff.StaffPosition,
			PhoneNumber: staff.Phone,
			StartDate:   staff.StartDate,
			Salary:      staff.Salary,
			Status:      staff.Status,
			Birthdate:   staff.Birthdate,
			Hometown:    staff.Hometown,
			BranchId:    staff.BranchId,
		})
	}
	return res, nil
}

func (s *staffServiceController) GetDetailStaff(ctx context.Context, staffId string) (*schema.GetStaffResponseData, error) {
	staff, err := s.repository.GetStaffDetail(ctx, staffId)
	if err != nil {
		return nil, err
	}
	return &schema.GetStaffResponseData{
		StaffId:     staffId,
		StaffName:   staff.StaffName,
		Street:      staff.Street,
		Ward:        staff.Ward,
		District:    staff.District,
		Province:    staff.Province,
		CitizenId:   staff.CitizenId,
		Role:        staff.StaffPosition,
		PhoneNumber: staff.Phone,
		StartDate:   staff.StartDate,
		Salary:      staff.Salary,
		Status:      staff.Status,
		Birthdate:   staff.Birthdate,
		Hometown:    staff.Hometown,
		BranchId:    staff.BranchId,
	}, nil
}
