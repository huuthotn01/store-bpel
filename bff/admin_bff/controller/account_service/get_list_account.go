package account_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/account_service"
	"time"
)

func (c *accountBffController) GetListAccount(ctx context.Context, request *account_service.GetListAccountRequest) ([]*account_service.GetListAccountResponseData, error) {
	// TODO call adapter to get real data
	if request.Username != "" {
		return []*account_service.GetListAccountResponseData{
			{
				Username:    "Name filtered",
				Role:        1,
				IsActivated: true,
				CreatedAt:   time.Now(),
			},
		}, nil
	}
	return []*account_service.GetListAccountResponseData{
		{
			Username:    "LVTN",
			Role:        3,
			IsActivated: true,
			CreatedAt:   time.Now(),
		},
		{
			Username:    "Name filtered",
			Role:        1,
			IsActivated: true,
			CreatedAt:   time.Now(),
		},
	}, nil
}
