package domain

import "context"

type Product struct {
	ID          int        `json:"id" gorm:"column:id;primary_key" validate:"omitempty"`
	Name        string     `json:"name" gorm:"c:name;" validate:"required,max=255"`
	Description string     `json:"description" gorm:"column:description" validate:"required"`
	Enable      bool       `json:"enable" gorm:"column:enable" validate:"omitempty"`
	Categories  []Category `json:"categories,omitempty" gorm:"many2many:category_products" validate:"-"`
	CategoryIDs []int      `json:"-" gorm:"-" validate:"omitempty"`
	Images      []Image    `json:"images,omitempty" gorm:"many2many:product_images" validate:"-"`
	ImageIDs    []int      `json:"-" gorm:"-" validate:"omitempty"`
}

type ProductImage struct {
	ProductID int `gorm:"omitempty;product_id"`
	ImageID   int `gorm:"omitempty;image_id"`
}

func (Product) TableName() string {
	return "products"
}

type ProductRepository interface {
	Create(context.Context, Product) (*Product, error)
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
