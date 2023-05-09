package controller

import (
	"context"
	"gorm.io/gorm"
	"os"
	"store-bpel/event_service/internal/repository"
	"testing"
)

var (
	testRepository TestRepository
)

// MOCK REPOSITORY
type TestRepository interface {
	GetAllEvent(ctx context.Context) ([]*repository.EventModel, error)
	GetAllEventCurrent(ctx context.Context, date int) ([]*repository.EventModel, error)
	GetGoods(ctx context.Context, eventId string) ([]string, error)
	GetEvent(ctx context.Context, eventId string) (*repository.EventModel, error)
	AddEvent(ctx context.Context, data *repository.AddEventData) error
	AddGoods(ctx context.Context, eventId string, listGoods []string) error
	UpdateEvent(ctx context.Context, data *repository.UpdateEventData) error
	DeleteGoods(ctx context.Context, eventId string) error
	DeleteEvent(ctx context.Context, eventId string) error
	GetEventByGoods(ctx context.Context, goodsId string) ([]*repository.EventModel, error)
	UpdateImage(ctx context.Context, eventId string, imageUrl string) error
}

type testRepo struct {
}

func NewTestRepo() TestRepository {
	return &testRepo{}
}

func (t *testRepo) GetAllEvent(ctx context.Context) ([]*repository.EventModel, error) {
	return []*repository.EventModel{
		{
			EventId:   "event-1",
			Name:      "Test Event 1",
			Discount:  0.5,
			StartTime: "2023-01-01",
			EndTime:   "2023-06-01",
		},
		{
			EventId:   "event-2",
			Name:      "Test Event 2",
			Discount:  0.3,
			StartTime: "2023-01-01",
			EndTime:   "2023-01-05",
		},
	}, nil
}

func (t *testRepo) GetAllEventCurrent(ctx context.Context, date int) ([]*repository.EventModel, error) {
	return []*repository.EventModel{
		{
			EventId:  "event-1",
			Name:     "Test Event 1",
			Discount: 0.5,
		},
	}, nil
}

func (t *testRepo) GetGoods(ctx context.Context, eventId string) ([]string, error) {
	if eventId == "event-1" {
		return []string{"goods-1", "goods-2"}, nil
	}
	return []string{"goods-3", "goods-4"}, nil
}

func (t *testRepo) GetEvent(ctx context.Context, eventId string) (*repository.EventModel, error) {
	if eventId == "event-2" {
		return &repository.EventModel{
			EventId:   "event-2",
			Name:      "Test Event 2",
			Discount:  0.3,
			StartTime: "2023-06-06",
			EndTime:   "2023-12-12",
		}, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (t *testRepo) AddEvent(ctx context.Context, data *repository.AddEventData) error {
	return nil
}

func (t *testRepo) AddGoods(ctx context.Context, eventId string, listGoods []string) error {
	return nil
}

func (t *testRepo) UpdateEvent(ctx context.Context, data *repository.UpdateEventData) error {
	return nil
}

func (t *testRepo) DeleteGoods(ctx context.Context, eventId string) error {
	return nil
}

func (t *testRepo) DeleteEvent(ctx context.Context, eventId string) error {
	return nil
}

func (t *testRepo) GetEventByGoods(ctx context.Context, goodsId string) ([]*repository.EventModel, error) {
	return []*repository.EventModel{
		{
			EventId:  "event-1",
			Name:     "Test Event 1",
			Discount: 0.5,
		},
		{
			EventId:  "event-2",
			Name:     "Test Event 2",
			Discount: 0.3,
		},
	}, nil
}

func (t *testRepo) UpdateImage(ctx context.Context, eventId string, imageUrl string) error {
	return nil
}

// Test Main func
func TestMain(m *testing.M) {
	testRepository = NewTestRepo()

	exitVal := m.Run()

	os.Exit(exitVal)
}
