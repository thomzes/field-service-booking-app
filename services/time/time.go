package services

import (
	"context"

	"github.com/thomzes/field-service-booking-app/domain/dto"
	"github.com/thomzes/field-service-booking-app/domain/models"
	"github.com/thomzes/field-service-booking-app/repositories"
)

type TimeService struct {
	repository repositories.IRepositoryRegistry
}

type ITimeService interface {
	GetAll(context.Context) ([]dto.TimeResponse, error)
	GetByUUID(context.Context, string) (*dto.TimeResponse, error)
	Create(context.Context, *dto.TimeRequest) (*dto.TimeResponse, error)
}

func NewTimeService(repository repositories.IRepositoryRegistry) ITimeService {
	return &TimeService{repository: repository}
}

func (t *TimeService) GetAll(ctx context.Context) ([]dto.TimeResponse, error) {
	times, err := t.repository.GetTime().FindAll(ctx)
	if err != nil {
		return nil, err
	}

	timeResults := make([]dto.TimeResponse, 0, len(times))
	for _, time := range times {
		timeResults = append(timeResults, dto.TimeResponse{
			UUID:      time.UUID,
			StartTime: time.StartTime,
			EndTime:   time.EndTime,
			CreatedAt: time.CreatedAt,
			UpdatedAt: time.UpdatedAt,
		})
	}

	return timeResults, nil
}

func (t *TimeService) GetByUUID(ctx context.Context, uuid string) (*dto.TimeResponse, error) {
	time, err := t.repository.GetTime().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	response := dto.TimeResponse{
		UUID:      time.UUID,
		StartTime: time.StartTime,
		EndTime:   time.EndTime,
		CreatedAt: time.CreatedAt,
		UpdatedAt: time.UpdatedAt,
	}

	return &response, nil
}

func (t *TimeService) Create(ctx context.Context, request *dto.TimeRequest) (*dto.TimeResponse, error) {
	time := &dto.TimeRequest{
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
	}

	timeResult, err := t.repository.GetTime().Create(ctx, &models.Time{
		StartTime: time.StartTime,
		EndTime:   time.EndTime,
	})
	if err != nil {
		return nil, err
	}

	response := dto.TimeResponse{
		UUID:      timeResult.UUID,
		StartTime: timeResult.StartTime,
		EndTime:   timeResult.EndTime,
		CreatedAt: timeResult.CreatedAt,
		UpdatedAt: timeResult.UpdatedAt,
	}

	return &response, err
}
