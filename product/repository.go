package product

import (
	"context"

	"github.com/dynastymasra/privy/domain"
	"github.com/jinzhu/gorm"
)

const (
	TableName = "products"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, product domain.Product) error {
	return r.db.Omit("created_at").Table(TableName).Save(&product).Error
}

func (r *Repository) FindByID(ctx context.Context, id int) (*domain.Product, error) {
	var (
		result domain.Product
		query  = domain.Product{
			ID: id,
		}
	)

	if err := r.db.Table(TableName).Where(query).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) Fetch(ctx context.Context, offset, limit int) ([]domain.Product, error) {
	var result []domain.Product

	err := r.db.Table(TableName).Limit(limit).Offset(offset).Find(&result).Error

	return result, err
}

func (r *Repository) Update(ctx context.Context, product domain.Product) error {
	var (
		query = domain.Product{
			ID: product.ID,
		}
	)
	return r.db.Table(TableName).Where(query).Update(&product).Error
}

func (r *Repository) Delete(ctx context.Context, product domain.Product) error {
	return r.db.Table(TableName).Delete(product).Error
}
