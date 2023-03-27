package staff_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/staff_service"
)

func (c *staffBffController) GetRequestList(ctx context.Context) ([]*staff_service.GetRequestResponseData, error) {
	requests, err := c.staffAdapter.GetRequestList(ctx)
	if err != nil {
		return nil, err
	}

	respRequests := make([]*staff_service.GetRequestResponseData, 0, len(requests))
	for _, data := range respRequests {
		respRequests = append(respRequests, &staff_service.GetRequestResponseData{
			Id:          data.Id,
			RequestType: data.RequestType,
			RequestDate: data.RequestDate,
			StaffId:     data.StaffId,
			Status:      data.Status,
		})
	}

	return respRequests, nil
}
