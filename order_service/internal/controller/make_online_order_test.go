package controller

import (
	"context"
	"store-bpel/order_service/config"
	"store-bpel/order_service/schema"
	"testing"
)

func Test_orderServiceController_CreateOnlineOrder(t *testing.T) {
	type args struct {
		request *schema.MakeOnlineOrderRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should create online order successfully",
			args: args{
				request: &schema.MakeOnlineOrderRequest{
					GoodsList: []*schema.OrderGoodsRequest{
						{
							GoodsId:   "goods-1",
							Size:      "XL",
							Color:     "red",
							Quantity:  3,
							UnitPrice: 5000,
							Price:     15000,
							Tax:       0.1,
							Discount:  0.3,
						},
					},
					TransactionDate: "2023-01-01",
					TotalPrice:      15000,
					CustomerId:      "customer-1",
					PaymentMethod:   "momo",
					ShipFee:         1000,
					ExpectedDate:    "2023-01-05",
					NameReceiver:    "HTTN",
					EmailReceiver:   "httn@gmail.com",
					PhoneReceiver:   "0123456789",
					Address: &schema.Address{
						Street:   "NTH",
						Ward:     "5",
						District: "BT",
						Province: "HCMC",
					},
				},
			},
		},
	}

	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &orderServiceController{
				cfg:          cfg,
				repository:   testRepository,
				goodsAdapter: testGoods,
				kafkaAdapter: testKafka,
			}
			if err := c.CreateOnlineOrder(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("CreateOnlineOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
