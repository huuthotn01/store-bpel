package controller

import "context"

func (c *orderServiceController) GetBestSellingGoods(ctx context.Context) ([]string, error) {
	goods, err := c.repository.GetBestGoods(ctx)
	if err != nil {
		return nil, err
	}

	return goods, nil
}
