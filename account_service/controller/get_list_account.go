package controller

import (
	"context"
	"store-bpel/account_service/schema"
)

func (c *accountServiceController) GetListAccount(ctx context.Context, username string) ([]*schema.GetListAccountResponseData, error) {
	acc, err := c.repository.GetListAccount(ctx, username)
	if err != nil {
		return nil, err
	}
	res := make([]*schema.GetListAccountResponseData, 0, len(acc))
	for _, data := range acc {
		res = append(res, &schema.GetListAccountResponseData{
			Username:    data.Username,
			Role:        data.UserRole,
			IsActivated: data.IsActivated == 1,
			CreatedAt:   data.CreatedAt,
		})
	}
	return res, err
}
