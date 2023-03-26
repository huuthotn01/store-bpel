package staff_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/staff_service"
)

func (c *staffBffController) CreateDeleteRequest(ctx context.Context, request *staff_service.CreateDeleteRequest) error {
	return c.staffAdapter.CreateDeleteRequest(ctx, request.StaffId)
}
