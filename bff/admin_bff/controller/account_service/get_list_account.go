package account_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/account_service"
)

func (c *accountBffController) GetListAccount(ctx context.Context, request *account_service.GetListAccountRequest) ([]*account_service.GetListAccountResponseData, error) {
	accounts, err := c.accountAdapter.GetListAccount(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	resp := make([]*account_service.GetListAccountResponseData, 0, len(accounts))
	for _, acc := range accounts {
		resp = append(resp, &account_service.GetListAccountResponseData{
			Username:    acc.Username,
			Role:        acc.Role,
			IsActivated: acc.IsActivated,
			CreatedAt:   acc.CreatedAt,
		})
	}
	return resp, nil
}
