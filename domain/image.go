package domain

import "context"

type Image struct {
	ID         int       `json:"id,omitempty" gorm:"column:id;primary_key" validate:"omitempty"`
	Name       string    `json:"name" gorm:"column:name;" validate:"required,max=255"`
	File       string    `json:"file" gorm:"column:file;" validate:"required"`
	Enable     bool      `json:"enable" gorm:"column:enable;" validate:"omitempty"`
	Products   []Product `json:"products,omitempty" gorm:"many2many:product_images" validate:"-"`
	ProductIDs []int     `json:"-" gorm:"-" validate:"omitempty"`
}

func (Image) TableName() string {
	return "images"
}

type ImageRepository interface {
	Create(context.Context, Image) (*Image, error)
	FindByID(context.Context, int) (*Image, error)
	Fetch(context.Context, int, int) ([]Image, error)
	Update(context.Context, Image) error
	Delete(context.Context, Image) error
}

type ImageService interface {
	Create(context.Context, Image) (*Image, error)
	FindByID(context.Context, int) (*Image, error)
	Fetch(context.Context, int, int) ([]Image, error)
	Update(context.Context, int, Image) (*Image, error)
	Delete(context.Context, int) error
}
