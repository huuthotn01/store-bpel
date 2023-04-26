package controller

import (
	"context"
	"store-bpel/account_service/internal/repository"
	"store-bpel/account_service/schema"
	"strings"
)

func (c *accountServiceController) GetListAccount(ctx context.Context, username string) ([]*schema.GetListAccountResponseData, error) {
	acc, err := c.repository.GetListAccount(ctx, username)
	if err != nil {
		return nil, err
	}
	res := make([]*schema.GetListAccountResponseData, 0, len(acc))
	for _, data := range acc {
		if data.UserRole == 1 {
			err = c.appendCustomerData(ctx, data, &res)
			if err != nil {
				return nil, err
			}
		} else {
			err = c.appendStaffData(ctx, data, &res)
			if err != nil {
				return nil, err
			}
		}
	}
	return res, err
}

func (c *accountServiceController) appendStaffData(ctx context.Context, data *repository.AccountModel, result *[]*schema.GetListAccountResponseData) error {
	staffId := strings.Split(data.Username, "@")[0]
	staffData, err := c.staffAdapter.GetDetailStaff(ctx, staffId)
	if err != nil {
		return err
	}

	*result = append(*result, &schema.GetListAccountResponseData{
		Username:    data.Username,
		Id:          staffData.StaffId,
		Role:        data.UserRole,
		PhoneNumber: staffData.PhoneNumber,
		Email:       staffData.Email,
		Name:        staffData.StaffName,
		IsActivated: data.IsActivated == 1,
		CreatedAt:   data.CreatedAt,
	})
	return nil
}

func (c *accountServiceController) appendCustomerData(ctx context.Context, data *repository.AccountModel, result *[]*schema.GetListAccountResponseData) error {
	cust, err := c.customerAdapter.GetCustomer(ctx, data.Username)
	if err != nil {
		return err
	}

	*result = append(*result, &schema.GetListAccountResponseData{
		Username:    data.Username,
		Id:          data.Username,
		Role:        data.UserRole,
		PhoneNumber: cust.Phone,
		Email:       cust.Email,
		Name:        cust.Name,
		IsActivated: data.IsActivated == 1,
		CreatedAt:   data.CreatedAt,
	})
	return nil
}
