package controller

import (
	"context"
	"store-bpel/staff_service/schema"
	"testing"
)

func Test_staffServiceController_UpdateRequestStatus(t *testing.T) {
	type args struct {
		request   *schema.UpdateRequestStatusRequest
		requestId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should update request status correctly, case request approve",
			args: args{
				request: &schema.UpdateRequestStatusRequest{
					Status: "APPROVED",
				},
				requestId: "request-1",
			},
		},
		{
			name: "Should update request status correctly, case request pending",
			args: args{
				request: &schema.UpdateRequestStatusRequest{
					Status: "UNAPPROVED",
				},
				requestId: "request-1",
			},
		},
		{
			name: "Should update request status correctly, case request approve, delete request",
			args: args{
				request: &schema.UpdateRequestStatusRequest{
					Status: "APPROVE",
				},
				requestId: "request-2",
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
			if err := s.UpdateRequestStatus(ctx, tt.args.request, tt.args.requestId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRequestStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
