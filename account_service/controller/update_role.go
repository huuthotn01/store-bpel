package controller

import (
	"context"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) UpdateRole(ctx context.Context, username string, request *schema.UpdateRoleRequest) error {
	return c.repository.UpdateRole(ctx, username, request.Role)
}
