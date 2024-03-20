package models

type Category string

const (
	CategoryFryers Category = "Fryers"
	CategoryOther  Category = "Other"
)

// @MongoDB Product
type Product struct {
	ID           string   `json:"id,omitempty" bson:"_id,omitempty" validate:"omitempty"`
	Name         string   `json:"name" bson:"name" validate:"required"`
	Description  *string  `json:"description" bson:"description" omitempty:"true"`
	Price        float64  `json:"price" bson:"price" validate:"required"`
	CountInStock int      `json:"countInStock" bson:"countInStock" validate:"required"`
	Category     Category `json:"category" bson:"category"`
	Image        *string  `json:"image" bson:"image" omitempty:"true"`
}
