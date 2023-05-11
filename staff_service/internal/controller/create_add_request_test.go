package controller

import (
	"context"
	"store-bpel/staff_service/schema"
	"testing"
)

func Test_staffServiceController_CreateAddRequest(t *testing.T) {
	type args struct {
		request *schema.CreateAddRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should create add request successfully",
			args: args{
				request: &schema.CreateAddRequest{
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
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &staffServiceController{
				db:         db,
				repository: testRepository,
			}
			if err := s.CreateAddRequest(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("CreateAddRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
