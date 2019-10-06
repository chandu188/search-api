package storage

import "github.com/gfg/model"

// SearchParams provides criteria for the search
type SearchParams struct {
	Skip     int
	Size     int
	Filters  map[string]string
	Query  string
	SortKeys []Sort
}

//Sort sepecifies the sort keys for the sorting the result of search query
type Sort struct {
	Key string
	Asc bool
}

// Service provides interface to get the products matching the search criteria
type Service interface {
	GetProducts(params SearchParams) ([]*model.Product, int, error)
	AddProduct(prod *model.Product) error
}
