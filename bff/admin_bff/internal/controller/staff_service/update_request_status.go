package staff_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/staff_service"
	"store-bpel/staff_service/schema"
)

func (c *staffBffController) UpdateRequestStatus(ctx context.Context, request *staff_service.UpdateRequestStatusRequest) error {
	return c.staffAdapter.UpdateRequestStatus(ctx, request.RequestId, &schema.UpdateRequestStatusRequest{
		Status: request.Status,
	})
}
