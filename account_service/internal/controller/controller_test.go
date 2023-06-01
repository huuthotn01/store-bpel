package controller

import (
	"gorm.io/gorm"
	"reflect"
	"store-bpel/account_service/config"
	"testing"
)

func TestNewController(t *testing.T) {
	type args struct {
		cfg *config.Config
		db  *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want IAccountServiceController
	}{
		{
			name: "Should init controller with correct elements",
			args: args{
				cfg: &config.Config{
					HttpPort: 14083,
				},
				db: &gorm.DB{},
			},
			want: &accountServiceController{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.cfg, tt.args.db); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewController() = %v, want %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
			}
		})
	}
}
