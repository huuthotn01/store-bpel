package controller

import (
	"context"
	"store-bpel/staff_service/schema"
)

func (s *staffServiceController) GetStaffAttendance(ctx context.Context, staffId string) ([]*schema.GetStaffAttendanceResponseData, error) {
	attendance, err := s.repository.GetStaffAttendance(ctx, staffId)
	if err != nil {
		return nil, err
	}
	result := make([]*schema.GetStaffAttendanceResponseData, 0, len(attendance))
	for _, atte := range attendance {
		result = append(result, &schema.GetStaffAttendanceResponseData{
			AttendanceDate: atte.AttendanceDate,
			CheckinTime:    atte.CheckinTime,
			CheckoutTime:   atte.CheckoutTime,
		})
	}
	return result, nil
}
