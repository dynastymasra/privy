package domain

type Category struct {
	ID       int       `json:"id,omitempty" gorm:"column:id;primary_key" validate:"omitempty"`
	Name     string    `json:"name" gorm:"column:name;" validate:"required,max=255"`
	Enable   bool      `json:"enable" gorm:"column:enable;" validate:"required"`
	Products []Product `json:"products,omitempty" gorm:"many2many:category_products" validate:"-"`
}

func (Category) TableName() string {
	return "categories"
}
