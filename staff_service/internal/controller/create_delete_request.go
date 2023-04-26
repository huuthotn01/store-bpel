package controller

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"store-bpel/staff_service/internal/repository"
	"time"
)

func (s *staffServiceController) CreateDeleteRequest(ctx context.Context, staffId string) error {
	requestId := fmt.Sprintf("del_%s", cast.ToString(time.Now().Unix()))
	return s.repository.CreateStaffRequest(ctx, &repository.RequestsModel{
		Id:          requestId,
		RequestDate: time.Now(),
		RequestType: "DELETE",
		StaffId:     staffId,
		Status:      "PENDING",
	})
}
