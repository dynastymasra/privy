package domain

import "context"

type Product struct {
	ID          int        `json:"id" gorm:"column:id;primary_key" validate:"omitempty"`
	Name        string     `json:"name" gorm:"column:name;" validate:"required,max=255"`
	Description string     `json:"description" gorm:"column:description" validate:"required"`
	Enable      bool       `json:"enable" gorm:"column:enable" validate:"required"`
	Categories  []Category `json:"categories,omitempty" gorm:"many2many:category_products" validate:"-"`
	Images      []Image    `json:"images,omitempty" gorm:"many2many:product_images" validate:"-"`
}

func (Product) TableName() string {
	return "products"
}

type ProductRepository interface {
	Create(context.Context, Product) error
	FindByID(context.Context, int) (*Product, error)
	Fetch(context.Context, int, int) ([]Product, error)
	Update(context.Context, Product) error
	Delete(context.Context, Product) error
}

type ProductService interface {
	Create(context.Context, Product) (*Product, error)
	FindByID(context.Context, int) (*Product, error)
	Fetch(context.Context, int, int) ([]Product, error)
	Update(context.Context, int, Product) (*Product, error)
	Delete(context.Context, int) error
}
