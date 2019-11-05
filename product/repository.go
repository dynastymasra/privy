package product

import (
	"context"

	"github.com/dynastymasra/privy/config"

	"github.com/dynastymasra/privy/domain"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, product domain.Product) (*domain.Product, error) {
	var images, categories []interface{}

	txn := r.db.Begin()

	if err := txn.Omit("created_at").Table(config.TableNameProduct).Create(&product).Error; err != nil {
		txn.Rollback()
		return nil, err
	}

	for _, image := range product.ImageIDs {
		images = append(images, domain.ProductImage{
			ImageID:   image,
			ProductID: product.ID,
		})
	}

	for _, category := range product.CategoryIDs {
		categories = append(categories, domain.CategoryProduct{
			CategoryID: category,
			ProductID:  product.ID,
		})
	}

	if err := gormbulk.BulkInsert(txn.Table(config.TableNameProductImages), images, 3000); err != nil {
		txn.Rollback()
		return nil, err
	}

	if err := gormbulk.BulkInsert(txn.Table(config.TableNameCategoryProducts), categories, 3000); err != nil {
		txn.Rollback()
		return nil, err
	}

	txn.Commit()

	return &product, nil
}

func (r *Repository) FindByID(ctx context.Context, id int) (*domain.Product, error) {
	result := domain.Product{ID: id}

	if err := r.db.Table(config.TableNameProduct).Preload("Categories").Preload("Images").First(&result).Error; err != nil {
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
		images, categories []interface{}
		query              = map[string]interface{}{
			"product_id": product.ID,
		}
	)

	if notFound := r.db.Table(config.TableNameProduct).First(&product).RecordNotFound(); notFound {
		return gorm.ErrRecordNotFound
	}

	txn := r.db.Begin()

	if err := txn.Table(config.TableNameProductImages).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameCategoryProducts).Where(query).Delete(nil).Error; err != nil {
		txn.Rollback()
		return err
	}

	if err := txn.Table(config.TableNameProduct).Where(domain.Product{ID: product.ID}).Update(&product).Error; err != nil {
		txn.Rollback()
		return err
	}

	for _, image := range product.ImageIDs {
		images = append(images, domain.ProductImage{
			ImageID:   image,
			ProductID: product.ID,
		})
	}

	for _, category := range product.CategoryIDs {
		categories = append(categories, domain.CategoryProduct{
			CategoryID: category,
			ProductID:  product.ID,
		})
	}

	if err := gormbulk.BulkInsert(txn.Table(config.TableNameProductImages), images, 3000); err != nil {
		txn.Rollback()
		return err
	}

	if err := gormbulk.BulkInsert(txn.Table(config.TableNameCategoryProducts), categories, 3000); err != nil {
		txn.Rollback()
		return err
	}

	txn.Commit()

	return nil
}

func (r *Repository) Delete(ctx context.Context, product domain.Product) error {
	if notFound := r.db.Table(config.TableNameProduct).First(&product).RecordNotFound(); notFound {
		return gorm.ErrRecordNotFound
	}

	return r.db.Table(config.TableNameProduct).Delete(&product).Error
}
