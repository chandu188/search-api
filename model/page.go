package model

// PageOfProducts is a single page of products
type PageOfProducts struct {

	// number of products that match the search criteria
	TotalItems int `json:"total_items"`

	// number of the page
	PageNumber int `json:"page_number"`

	// number of elements in the page
	PageSize int `json:"page_size"`

	// List of products
	Products []*Product `json:"products"`
}
