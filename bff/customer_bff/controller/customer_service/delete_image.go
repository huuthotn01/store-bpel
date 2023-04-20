package customer_service

import "context"

func (c *customerBffController) DeleteImage(ctx context.Context, username string) error {
	return c.customerAdapter.DeleteImage(ctx, username)
}
