package controller

import (
	"context"
	"reflect"
	"testing"
)

func Test_orderServiceController_GetBestSellingGoods(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		{
			name: "Should get best selling goods correctly",
			want: []string{"goods-1", "goods-2"},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &orderServiceController{
				repository: testRepository,
			}
			got, err := c.GetBestSellingGoods(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBestSellingGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBestSellingGoods() got = %v, want %v", got, tt.want)
			}
		})
	}
}
