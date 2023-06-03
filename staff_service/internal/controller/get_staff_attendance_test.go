package controller

import (
	"context"
	"database/sql"
	"reflect"
	"store-bpel/staff_service/schema"
	"testing"
	"time"
)

func Test_staffServiceController_GetStaffAttendance(t *testing.T) {
	type args struct {
		staffId string
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.GetStaffAttendanceResponseData
		wantErr bool
	}{
		{
			name: "Should get staff attendance correctly",
			args: args{
				staffId: "staff-1",
			},
			want: []*schema.GetStaffAttendanceResponseData{
				{
					AttendanceDate: "2023-01-01",
					CheckinTime:    time.Date(2023, 1, 1, 9, 0, 0, 0, time.Local),
					CheckoutTime: sql.NullTime{
						Time: time.Date(2023, 1, 1, 18, 3, 0, 0, time.Local),
					},
				},
				{
					AttendanceDate: "2023-01-01",
					CheckinTime:    time.Date(2023, 1, 3, 8, 49, 22, 0, time.Local),
				},
			},
		},
		{
			name: "Should return error when db get staff attendance fails",
			args: args{
				staffId: "invalid-staff",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &staffServiceController{
				repository: testRepository,
			}
			got, err := s.GetStaffAttendance(ctx, tt.args.staffId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaffAttendance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaffAttendance() got = %v, want %v", got, tt.want)
			}
		})
	}
}
