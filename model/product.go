package model

// Product represents a value of
type Product struct {
	Sku   string  `json:"sku"`
	Title string  `json:"title,omitempty"`
	Brand string  `json:"brand"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
	Tags  []Tag   `json:"tags,omitempty"`
}

// Tag represent a tag for product
type Tag string
