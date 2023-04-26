package staff_service

import (
	"context"
	"store-bpel/bff/admin_bff/schema/staff_service"
)

func (c *staffBffController) DeleteStaff(ctx context.Context, request *staff_service.CreateDeleteRequest) error {
	return c.staffAdapter.DeleteStaff(ctx, request.StaffId)
}
