package controller

import (
	"context"
	"store-bpel/goods_service/schema"
	"testing"
)

func Test_goodsServiceController_CreateGoodsTransaction(t *testing.T) {
	type args struct {
		request         *schema.CreateGoodsTransactionRequest
		transactionType string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should create import transaction successfully",
			args: args{
				request: &schema.CreateGoodsTransactionRequest{
					GoodsCode:  "goods-1",
					GoodsColor: "red",
					GoodsSize:  "XL",
					Quantity:   2,
					From:       "warehouse-1",
					To:         "warehouse-2",
				},
				transactionType: "IMPORT",
			},
		},
		{
			name: "Should create export transaction successfully",
			args: args{
				request: &schema.CreateGoodsTransactionRequest{
					GoodsCode:  "goods-1",
					GoodsColor: "red",
					GoodsSize:  "XL",
					Quantity:   2,
					From:       "warehouse-1",
					To:         "warehouse-2",
				},
				transactionType: "EXPORT",
			},
		},
		{
			name: "Should create transfer transaction successfully",
			args: args{
				request: &schema.CreateGoodsTransactionRequest{
					GoodsCode:  "goods-1",
					GoodsColor: "red",
					GoodsSize:  "XL",
					Quantity:   2,
					From:       "warehouse-1",
					To:         "warehouse-2",
				},
				transactionType: "TRANSFER",
			},
		},
		{
			name: "Should create return manufacturer transaction successfully",
			args: args{
				request: &schema.CreateGoodsTransactionRequest{
					GoodsCode:  "goods-1",
					GoodsColor: "red",
					GoodsSize:  "XL",
					Quantity:   2,
					From:       "warehouse-1",
					To:         "warehouse-2",
				},
				transactionType: "RETURN_MANUFACT",
			},
		},
		{
			name: "Should create customer return transaction successfully",
			args: args{
				request: &schema.CreateGoodsTransactionRequest{
					GoodsCode:  "goods-1",
					GoodsColor: "red",
					GoodsSize:  "XL",
					Quantity:   2,
					From:       "warehouse-1",
					To:         "warehouse-2",
				},
				transactionType: "CUST_RETURN",
			},
		},
		{
			name: "Should return error transaction type invalid",
			args: args{
				request: &schema.CreateGoodsTransactionRequest{
					GoodsCode:  "goods-1",
					GoodsColor: "red",
					GoodsSize:  "XL",
					Quantity:   2,
					From:       "warehouse-1",
					To:         "warehouse-2",
				},
				transactionType: "SOMETHING",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goodsServiceController{
				repository: testRepository,
			}
			if err := c.CreateGoodsTransaction(ctx, tt.args.request, tt.args.transactionType); (err != nil) != tt.wantErr {
				t.Errorf("CreateGoodsTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
