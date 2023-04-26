package controller

import "context"

func (c *customerServiceController) DeleteImage(ctx context.Context, username string) error {
	return c.repository.UpdateCustomerImage(ctx, username, "")
}
