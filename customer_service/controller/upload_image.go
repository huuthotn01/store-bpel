package controller

import (
	"context"
	"store-bpel/customer_service/schema"
)

func (c *customerServiceController) UploadImage(ctx context.Context, request *schema.UploadImageRequest) error {
	return c.repository.UpdateCustomerImage(ctx, request.Username, request.ImageUrl)
}
