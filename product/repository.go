package product

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

func (r *Repository) Create(ctx context.Context, product domain.Product) (*domain.Product, error) {
	if err := r.db.Omit("created_at").Table(config.TableNameProduct).Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *Repository) FindByID(ctx context.Context, id int) (*domain.Product, error) {
	var (
		result domain.Product
		query  = domain.Product{
			ID: id,
		}
	)

	if err := r.db.Table(config.TableNameProduct).Where(query).Preload("Categories").Preload("Images").First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) Fetch(ctx context.Context, offset, limit int) ([]domain.Product, error) {
	var result []domain.Product

	err := r.db.Table(config.TableNameProduct).Limit(limit).Offset(offset).Preload("Categories").Preload("Images").Find(&result).Error

	return result, err
}

func (r *Repository) Update(ctx context.Context, product domain.Product) error {
	var (
		query = map[string]interface{}{
			"product_id": product.ID,
		}
	)

	txn := r.db.Begin()

	if err := txn.Table(config.TableNameProductImages).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameCategoryProducts).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameProduct).Save(&product).Error; err != nil {
		txn.Rollback()
		return err
	}

	txn.Commit()

	return nil
}

func (r *Repository) Delete(ctx context.Context, product domain.Product) error {
	var (
		query = map[string]interface{}{
			"product_id": product.ID,
		}
	)

	txn := r.db.Begin()

	if err := txn.Table(config.TableNameProductImages).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameCategoryProducts).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameProduct).Delete(&product).Error; err != nil {
		txn.Rollback()
		return err
	}

	txn.Commit()

	return nil
}
