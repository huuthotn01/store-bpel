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
	for _, data := range requests {
		respRequests = append(respRequests, &staff_service.GetRequestResponseData{
			Id:            data.Id,
			RequestDate:   data.RequestDate,
			RequestType:   data.RequestType, // ADD or DELETE
			Status:        data.Status,
			StaffId:       data.StaffId,
			StaffName:     data.StaffName,
			Province:      data.Province,
			District:      data.District,
			Ward:          data.Ward,
			Street:        data.Street,
			Hometown:      data.Hometown,
			CitizenId:     data.CitizenId,
			StaffPosition: data.StaffPosition,
			Birthdate:     data.Birthdate,
			StartDate:     data.StartDate,
			Salary:        data.Salary,
			Gender:        data.Gender,
			Phone:         data.Phone,
			Email:         data.Email,
			BranchId:      data.BranchId,
		})
	}

	return respRequests, nil
}
