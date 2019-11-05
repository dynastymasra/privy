package category

import (
	"context"

	gormbulk "github.com/t-tiger/gorm-bulk-insert"

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
	var products []interface{}

	txn := r.db.Begin()

	if err := txn.Omit("created_at").Table(config.TableNameCategory).Create(&category).Error; err != nil {
		txn.Rollback()
		return nil, err
	}

	for _, product := range category.ProductIDs {
		products = append(products, domain.CategoryProduct{
			CategoryID: category.ID,
			ProductID:  product,
		})
	}

	if err := gormbulk.BulkInsert(txn.Table(config.TableNameCategoryProducts), products, 3000); err != nil {
		txn.Rollback()
		return nil, err
	}

	txn.Commit()

	return &category, nil
}

func (r *Repository) FindByID(ctx context.Context, id int) (*domain.Category, error) {
	result := domain.Category{ID: id}

	if err := r.db.Table(config.TableNameCategory).Preload("Products").First(&result).Error; err != nil {
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
		products []interface{}
		query    = map[string]interface{}{
			"category_id": category.ID,
		}
	)

	if notFound := r.db.Table(config.TableNameCategory).First(&domain.Category{ID: category.ID}).RecordNotFound(); notFound {
		return gorm.ErrRecordNotFound
	}

	txn := r.db.Begin()

	if err := txn.Table(config.TableNameCategoryProducts).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameCategory).Save(&category).Error; err != nil {
		txn.Rollback()
		return err
	}

	for _, product := range category.ProductIDs {
		products = append(products, domain.CategoryProduct{
			CategoryID: category.ID,
			ProductID:  product,
		})
	}

	if err := gormbulk.BulkInsert(txn.Table(config.TableNameCategoryProducts), products, 3000); err != nil {
		txn.Rollback()
		return err
	}

	txn.Commit()

	return nil
}

func (r *Repository) Delete(ctx context.Context, category domain.Category) error {
	if notFound := r.db.Table(config.TableNameCategory).First(&category).RecordNotFound(); notFound {
		return gorm.ErrRecordNotFound
	}

	return r.db.Table(config.TableNameCategory).Delete(&category).Error
}
