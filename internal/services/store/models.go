package delivery

import "time"

type ProductStock struct {
	ProductID string
	Count     string
	UpdateAt  time.Time
}

type Product struct {
	ID         string
	PLU        string
	Price      float64
	Discount   float64
	CategoryID string
}
