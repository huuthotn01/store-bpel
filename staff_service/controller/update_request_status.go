package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/staff_service/schema"
)

func (s *staffServiceController) UpdateRequestStatus(ctx context.Context, request *schema.UpdateRequestStatusRequest, requestId string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// check whether request is add or delete request
		req, err := s.repository.GetStaffRequest(ctx, requestId)
		if err != nil {
			return err
		}
		// update request status
		err = s.repository.UpdateRequestStatus(ctx, request.Status, requestId)
		if err != nil {
			return err
		}
		// in case delete request, delete staff
		if req.RequestType == 2 { // request is delete
			err = s.repository.DeleteStaff(ctx, req.StaffId)
			if err != nil {
				return err
			}
		} else {
			// TODO call account service to create account for user
		}
		return nil
	})
}
