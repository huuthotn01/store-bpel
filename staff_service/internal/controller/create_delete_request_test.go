package controller

import (
	"context"
	"testing"
)

func Test_staffServiceController_CreateDeleteRequest(t *testing.T) {
	type args struct {
		staffId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should create delete request",
			args: args{
				staffId: "staff-1",
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
			if err := s.CreateDeleteRequest(ctx, tt.args.staffId); (err != nil) != tt.wantErr {
				t.Errorf("CreateDeleteRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
