package controller

import (
	"context"
	"store-bpel/staff_service/schema"
	"testing"
)

func Test_staffServiceController_AddStaff(t *testing.T) {
	type args struct {
		request *schema.AddStaffRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should add staff successfully",
			args: args{
				request: &schema.AddStaffRequest{
					Name:         "Huu Tho",
					Birthdate:    "2001-01-01",
					Hometown:     "Hue",
					CitizenId:    "9999999999",
					Phone:        "0123456789",
					Street:       "THT",
					Ward:         "11",
					District:     "10",
					Province:     "HCMC",
					WorkingPlace: "HCMUT",
					Gender:       "MALE",
					Salary:       1000000,
					Role:         "3",
					Email:        "tho@gmail.com",
				},
			},
		},
		{
			name: "Should return error when db add staff fails",
			args: args{
				request: &schema.AddStaffRequest{
					Name:         "Huu Tho",
					Birthdate:    "2001-01-01",
					Hometown:     "Hue",
					CitizenId:    "9999999999",
					Phone:        "0123456789",
					Street:       "THT",
					Ward:         "11",
					District:     "10",
					Province:     "HCMC",
					WorkingPlace: "HCMUT",
					Gender:       "MALE",
					Salary:       1000000,
					Role:         "3",
					Email:        "invalid-staff@gmail.com",
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when request role not valid",
			args: args{
				request: &schema.AddStaffRequest{
					Name:         "Huu Tho",
					Birthdate:    "2001-01-01",
					Hometown:     "Hue",
					CitizenId:    "9999999999",
					Phone:        "0123456789",
					Street:       "THT",
					Ward:         "11",
					District:     "10",
					Province:     "HCMC",
					WorkingPlace: "HCMUT",
					Gender:       "MALE",
					Salary:       1000000,
					Role:         "invalid",
					Email:        "httn@gmail.com",
				},
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &staffServiceController{
				db:           db,
				kafkaAdapter: testKafka,
				repository:   testRepository,
			}
			if err := s.AddStaff(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
