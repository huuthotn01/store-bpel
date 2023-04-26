package staff_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/staff_service"
)

func (c *staffBffController) GetStaffAttendance(ctx context.Context, request *staff_service.GetStaffAttendanceRequest) ([]*staff_service.GetStaffAttendanceResponseData, error) {
	attendance, err := c.staffAdapter.GetStaffAttendance(ctx, request.StaffId)
	if err != nil {
		return nil, err
	}

	respAttendance := make([]*staff_service.GetStaffAttendanceResponseData, 0, len(attendance))
	for _, data := range attendance {
		respAttendance = append(respAttendance, &staff_service.GetStaffAttendanceResponseData{
			AttendanceDate: data.AttendanceDate,
			CheckoutTime:   data.CheckoutTime,
			CheckinTime:    data.CheckinTime,
		})
	}

	return respAttendance, nil
}
