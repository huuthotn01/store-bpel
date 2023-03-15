package controller

import (
	"context"
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
		staffId := strings.Split(data.Username, "@")[0]
		staffData, err := c.staffAdapter.GetDetailStaff(ctx, staffId)
		if err != nil {
			return nil, err
		}

		res = append(res, &schema.GetListAccountResponseData{
			Username:    data.Username,
			StaffId:     staffData.StaffId,
			Role:        data.UserRole,
			PhoneNumber: staffData.PhoneNumber,
			StartDate:   staffData.StartDate,
			BirthDate:   staffData.Birthdate,
			Street:      staffData.Street,
			Ward:        staffData.Ward,
			District:    staffData.District,
			Province:    staffData.Province,
			Name:        staffData.StaffName,
			IsActivated: data.IsActivated == 1,
			CreatedAt:   data.CreatedAt,
		})
	}
	return res, err
}
