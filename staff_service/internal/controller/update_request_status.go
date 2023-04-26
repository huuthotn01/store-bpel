package controller

import (
	"context"
	"gorm.io/gorm"
	"store-bpel/staff_service/internal/repository"
	"store-bpel/staff_service/schema"
)

func (s *staffServiceController) UpdateRequestStatus(ctx context.Context, request *schema.UpdateRequestStatusRequest, requestId string) error {
	/*
		When receiving UpdateRequestStatus:
		- Update request's status in requests table to request.Status
		- If request is ADD request:
			+ If APPROVED:
				1. Change staff's status in staff table to APPROVED.
				2. Create staff's account in account table.
				3. Call account service to create account in account service.
			+ If UNAPPROVED: delete staff in staff table by removing.
		- If request is DELETE request:
			+ If APPROVED: delete staff in staff table by updating staff's status to DELETED.
			+ If UNAPPROVED: nothing changes.
	*/
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
		if request.Status == "APPROVED" { // in case request is approved, do corresponding actions
			// in case delete request, delete staff
			if req.RequestType == "DELETE" { // request is delete
				// delete staff by updating status to DELETED
				return s.repository.DeleteStaffUpdateStatus(ctx, req.StaffId)
			} else { // request is ADD
				staff, err := s.repository.GetStaffDetail(ctx, req.StaffId)
				if err != nil {
					return err
				}
				// change staff's status
				staff.Status = "APPROVED"
				err = s.repository.UpdateStaff(ctx, staff)
				// create account for staff
				return s.repository.CreateAccount(ctx, &repository.AccountModel{
					StaffId:  req.StaffId,
					Username: staff.Email,
				})
				// TODO call account service to create account for user
			}
		} else if req.RequestType == "ADD" {
			// in case request is not APPROVED => UNAPPROVED, only need to update ADD request
			// update staff, delete from staff table
			return s.repository.DeleteStaffRemove(ctx, req.StaffId)
		}
		return nil
	})
}
