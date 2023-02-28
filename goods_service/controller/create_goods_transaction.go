package controller

import (
	"context"
	"errors"
	"store-bpel/goods_service/schema"
)

func (c *goodsServiceController) CreateGoodsTransaction(ctx context.Context, request *schema.CreateGoodsTransactionRequest, transactionType string) error {
	if request.From == "" || request.To == "" {
		return errors.New("invalid From and To")
	}

	var err error
	switch transactionType {
	case "IMPORT":
		err = c.handleImport(ctx, request)
	case "EXPORT":
		err = c.handleExport(ctx, request)
	case "TRANSFER":
		err = c.handleWHTransfer(ctx, request)
	case "RETURN_MANUFACT":
		err = c.handleReturnManufact(ctx, request)
	case "CUST_RETURN":
		err = c.handleCustReturn(ctx, request)
	default:
		err = errors.New("invalid transactionType")
	}
	return err
}
