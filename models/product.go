package models

type Category string

const (
	Fryers Category = "Fryers"
	Other  Category = "Other"
)

type Product struct {
	ID           string   `json:"id" bson:"_id" validate:"required"`
	Name         string   `json:"name" bson:"name" validate:"required"`
	Description  string   `json:"description" bson:"description" omitempty:"true"`
	Price        float64  `json:"price" bson:"price" validate:"required"`
	CountInStock int      `json:"countInStock" bson:"countInStock" validate:"required"`
	Category     Category `json:"category" bson:"category"`
	Image        string   `json:"image" bson:"image" omitempty:"true"`
}
