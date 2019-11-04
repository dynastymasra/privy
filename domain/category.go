package domain

import "context"

type Category struct {
	ID       int       `json:"id,omitempty" gorm:"column:id;primary_key" validate:"omitempty"`
	Name     string    `json:"name" gorm:"column:name;" validate:"required,max=255"`
	Enable   bool      `json:"enable" gorm:"column:enable;" validate:"required"`
	Products []Product `json:"products,omitempty" gorm:"many2many:category_products" validate:"-"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryRepository interface {
	Create(context.Context, Category) (*Category, error)
	FindByID(context.Context, int) (*Category, error)
	Fetch(context.Context, int, int) ([]Category, error)
	Update(context.Context, Category) error
	Delete(context.Context, Category) error
}

type CategoryService interface {
	Create(context.Context, Category) (*Category, error)
	FindByID(context.Context, int) (*Category, error)
	Fetch(context.Context, int, int) ([]Category, error)
	Update(context.Context, int, Category) (*Category, error)
	Delete(context.Context, int) error
}
