package domain

type Product struct {
	ID          uint    `json:"id" gorm:"unique;not null"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"category_id"`
	Size        int     `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}