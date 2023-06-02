package controller

import (
	"gorm.io/gorm"
	"reflect"
	"store-bpel/staff_service/config"
	"testing"
)

func TestNewController(t *testing.T) {
	type args struct {
		config *config.Config
		db     *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want IStaffServiceController
	}{
		{
			name: "Should init controller with correct element",
			args: args{
				config: &config.Config{},
				db:     &gorm.DB{},
			},
			want: &staffServiceController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.config, tt.args.db); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}
