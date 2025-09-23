package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	errWrap "github.com/thomzes/field-service-booking-app/common/error"
	errConstant "github.com/thomzes/field-service-booking-app/constants/error"
	errTime "github.com/thomzes/field-service-booking-app/constants/error/time"

	"github.com/thomzes/field-service-booking-app/domain/models"
	"gorm.io/gorm"
)

type TimeRepository struct {
	db *gorm.DB
}

type ITimeRepository interface {
	FindAll(context.Context) ([]models.Time, error)
	FindByUUID(context.Context, string) (*models.Time, error)
	FindByID(context.Context, string) (*models.Time, error)
	Create(context.Context, *models.Time) (*models.Time, error)
}

func NewTimeRepository(db *gorm.DB) ITimeRepository {
	return &TimeRepository{db: db}
}

func (t *TimeRepository) FindAll(ctx context.Context) ([]models.Time, error) {
	var times []models.Time
	err := t.db.WithContext(ctx).Find(&times).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return times, nil
}

func (t *TimeRepository) FindByUUID(ctx context.Context, uuid string) (*models.Time, error) {
	var time models.Time
	err := t.db.WithContext(ctx).Where("uuid = ?", uuid).First(&time).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errTime.ErrTimeNotFound)
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &time, nil
}

func (t *TimeRepository) FindByID(ctx context.Context, id string) (*models.Time, error) {
	var time models.Time
	err := t.db.WithContext(ctx).Where("id = ?", id).First(&time).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errTime.ErrTimeNotFound)
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &time, nil
}

func (t *TimeRepository) Create(ctx context.Context, req *models.Time) (*models.Time, error) {
	time := models.Time{
		UUID:      uuid.New(),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	err := t.db.WithContext(ctx).Create(&time).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &time, err
}
