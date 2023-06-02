package controller

import (
	"gorm.io/gorm"
	"reflect"
	"store-bpel/event_service/config"
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
		want IEventServiceController
	}{
		{
			name: "Should init controller with correct element",
			args: args{
				cfg: &config.Config{},
				db: &gorm.DB{},
			},
			want: &eventServiceController{},
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
