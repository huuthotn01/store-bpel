package customer_service

import (
	"context"
	"store-bpel/bff/customer_bff/schema/customer_service"
	"store-bpel/customer_service/schema"
)

func (c *customerBffController) UploadImage(ctx context.Context, request *customer_service.UploadImageRequest) error {
	return c.customerAdapter.UploadImage(ctx, &schema.UploadImageRequest{
		Username: request.Username,
		ImageUrl: request.ImageUrl,
	})
}
