package category

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

func (r *Repository) Create(ctx context.Context, category domain.Category) (*domain.Category, error) {
	if err := r.db.Omit("created_at").Table(config.TableNameCategory).Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *Repository) FindByID(ctx context.Context, id int) (*domain.Category, error) {
	var (
		result domain.Category
		query  = domain.Category{
			ID: id,
		}
	)

	if err := r.db.Table(config.TableNameCategory).Where(query).Preload("Products").First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) Fetch(ctx context.Context, offset, limit int) ([]domain.Category, error) {
	var result []domain.Category

	err := r.db.Table(config.TableNameCategory).Limit(limit).Offset(offset).Preload("Products").Find(&result).Error

	return result, err
}

func (r *Repository) Update(ctx context.Context, category domain.Category) error {
	var (
		query = map[string]interface{}{
			"category_id": category.ID,
		}
	)

	txn := r.db.Begin()

	if err := txn.Table(config.TableNameCategoryProducts).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameCategory).Save(&category).Error; err != nil {
		txn.Rollback()
		return err
	}

	txn.Commit()

	return nil
}

func (r *Repository) Delete(ctx context.Context, category domain.Category) error {
	var (
		query = map[string]interface{}{
			"category_id": category.ID,
		}
	)

	txn := r.db.Begin()

	if err := txn.Table(config.TableNameCategoryProducts).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameCategory).Delete(&category).Error; err != nil {
		txn.Rollback()
		return err
	}

	txn.Commit()

	return nil
}
