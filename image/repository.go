package image

import (
	"context"

	"github.com/dynastymasra/privy/config"

	"github.com/dynastymasra/privy/domain"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, image domain.Image) (*domain.Image, error) {
	if err := r.db.Omit("created_at").Table(config.TableNameImage).Create(&image).Error; err != nil {
		return nil, err
	}

	return &image, nil
}

func (r *Repository) FindByID(ctx context.Context, id int) (*domain.Image, error) {
	var (
		result domain.Image
		query  = domain.Image{
			ID: id,
		}
	)

	if err := r.db.Table(config.TableNameImage).Where(query).Preload("Products").First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) Fetch(ctx context.Context, offset, limit int) ([]domain.Image, error) {
	var result []domain.Image

	err := r.db.Table(config.TableNameImage).Limit(limit).Offset(offset).Preload("Products").Find(&result).Error

	return result, err
}

func (r *Repository) Update(ctx context.Context, image domain.Image) error {
	var (
		query = map[string]interface{}{
			"image_id": image.ID,
		}
	)

	txn := r.db.Begin()

	if err := txn.Table(config.TableNameProductImages).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameImage).Save(&image).Error; err != nil {
		txn.Rollback()
		return err
	}

	txn.Commit()

	return nil
}

func (r *Repository) Delete(ctx context.Context, image domain.Image) error {
	var (
		query = map[string]interface{}{
			"image_id": image.ID,
		}
	)

	txn := r.db.Begin()

	if err := txn.Table(config.TableNameProductImages).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameImage).Delete(&image).Error; err != nil {
		txn.Rollback()
		return err
	}

	txn.Commit()

	return nil
}
