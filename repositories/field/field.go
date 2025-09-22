package repositories

import (
	"context"
	"errors"
	"fmt"

	errWrap "github.com/thomzes/field-service-booking-app/common/error"
	errConstant "github.com/thomzes/field-service-booking-app/constants/error"
	errField "github.com/thomzes/field-service-booking-app/constants/error/field"
	"github.com/thomzes/field-service-booking-app/domain/dto"
	"github.com/thomzes/field-service-booking-app/domain/models"
	"gorm.io/gorm"
)

type FieldRepository struct {
	db *gorm.DB
}

type IFieldRepository interface {
	FindAllWithPagination(context.Context, *dto.FieldRequestParam) ([]models.Field, int64, error)
	FindAllWithoutPagination(context.Context) ([]models.Field, error)
	FindByUUID(context.Context, string) (*models.Field, error)
	Create(context.Context, *models.Field) (*models.Field, error)
	Update(context.Context, string, *models.Field) (*models.Field, error)
	Delete(context.Context, string) error
}

func NewFieldRepository(db *gorm.DB) FieldRepository {
	return &FieldRepository{db: db}
}

func (f *FieldRepository) FindAllWithPagination(ctx context.Context, param *dto.FieldRequestParam) ([]models.Field, int64, error) {
	var (
		fields []models.Field
		sort   string
		total  int64
	)

	if param.SortColumn != nil {
		sort = fmt.Sprintf("%s %s", *param.SortColumn, *param.SortOrder)
	} else {
		sort = "created_at desc"
	}

	limit := param.Limit
	offset := (param.Page - 1) * limit
	err := f.db.WithContext(ctx).Limit(limit).Offset(offset).Order(sort).Find(&fields).Error
	if err != nil {
		return nil, 0, errWrap.WrapError(errConstant.ErrSQLError)
	}

	err = f.db.WithContext(ctx).Model(&fields).Count(&total).Error
	if err != nil {
		return nil, 0, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return fields, total, nil
}

func (f *FieldRepository) FindAllWithoutPagination(ctx context.Context) ([]models.Field, error) {
	var fields []models.Field
	err := f.db.WithContext(ctx).Find(&fields).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}
	return fields, nil
}

func (f *FieldRepository) FindByUUID(ctx context.Context, uuid string) (*models.Field, error) {
	var field models.Field
	err := f.db.WithContext(ctx).Where("uuid = ?", uuid).Find(&field).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errField.ErrFieldNotFound)
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}
	return &field, nil
}

func (f *FieldRepository) Create(ctx context.Context, req *models.Field) (*models.Field, error) {

}

func (f *FieldRepository) Update(ctx context.Context, uuid string, field *models.Field) (*models.Field, error) {

}

func (f *FieldRepository) Delete(ctx context.Context, uuid string) error {

}
