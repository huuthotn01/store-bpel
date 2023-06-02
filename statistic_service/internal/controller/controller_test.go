package controller

import (
	"gorm.io/gorm"
	"reflect"
	"store-bpel/statistic_service/config"
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
		want IStatisticServiceController
	}{
		{
			name: "Should init controller with correct elements",
			args: args{
				cfg: &config.Config{},
				db:  &gorm.DB{},
			},
			want: &statisticServiceController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.cfg, tt.args.db); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}
