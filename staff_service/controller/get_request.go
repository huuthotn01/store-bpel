package controller

import (
	"context"
	"store-bpel/staff_service/schema"
)

func (s *staffServiceController) GetRequest(ctx context.Context) ([]*schema.GetRequestResponseData, error) {
	requestsList, err := s.repository.GetListRequest(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*schema.GetRequestResponseData, 0, len(requestsList))
	for _, request := range requestsList {

		res = append(res, &schema.GetRequestResponseData{
			Id:            request.Id,
			RequestDate:   request.RequestDate,
			RequestType:   request.RequestType, // ADD or DELETE
			Status:        request.Status,
			StaffId:       request.StaffId,
			StaffName:     request.StaffName,
			Province:      request.Province,
			District:      request.District,
			Ward:          request.Ward,
			Street:        request.Street,
			Hometown:      request.Hometown,
			CitizenId:     request.CitizenId,
			StaffPosition: request.StaffPosition,
			Birthdate:     request.Birthdate,
			StartDate:     request.StartDate,
			Salary:        request.Salary,
			Gender:        request.Gender,
			Phone:         request.Phone,
			Email:         request.Email,
			BranchId:      request.BranchId,
		})
	}
	return res, nil
}
