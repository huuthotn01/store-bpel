package account_service

import (
	"context"
	"store-bpel/account_service/schema"
	"store-bpel/bff/admin_bff/schema/account_service"
)

func (c *accountBffController) UpdateRole(ctx context.Context, request *account_service.UpdateRoleRequest) error {
	return c.accountAdapter.UpdateRole(ctx, request.Username, &schema.UpdateRoleRequest{
		Role: request.Role,
	})
}
