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
			Id:          request.Id,
			RequestDate: request.RequestDate,
			RequestType: request.RequestType,
			StaffId:     request.StaffId,
			Status:      request.Status,
		})
	}
	return res, nil
}
