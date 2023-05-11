package controller

import (
	"context"
	"reflect"
	"store-bpel/staff_service/schema"
	"testing"
)

func Test_staffServiceController_GetRequest(t *testing.T) {
	tests := []struct {
		name    string
		want    []*schema.GetRequestResponseData
		wantErr bool
	}{
		{
			name: "Should get request correctly",
			want: []*schema.GetRequestResponseData{
				{
					Id:            "request-1",
					RequestType:   "ADD",
					Status:        "PENDING",
					StaffId:       "staff-2",
					StaffName:     "TNHT",
					Province:      "Binh Duong",
					District:      "Di An",
					Ward:          "Dong Hoa",
					Street:        "Nguyen Du",
					Hometown:      "HCMC",
					CitizenId:     "1234567890",
					StaffPosition: "4",
					Birthdate:     "2001-01-02",
					Salary:        1000000,
					Gender:        "MALE",
					Phone:         "0123456789",
					Email:         "tnht@gmail.com",
					BranchId:      "branch-1",
				},
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &staffServiceController{
				repository: testRepository,
			}
			got, err := s.GetRequest(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
