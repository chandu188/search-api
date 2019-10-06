package service

import (
	"github.com/gfg/model"
	"github.com/gfg/storage"
	log "github.com/sirupsen/logrus"
)

type SvcParams struct {
	Filters    map[string]string
	Query      string
	SortKeys   map[string]bool
	From, Size int
}

//Product represents an interface to get collection of products matching the search criteria
type Product interface {
	GetProducts(params SvcParams) (*model.PageOfProducts, *model.Problem)
	AddProduct(p *model.Product) *model.Problem
}

// NewProductService returns a Product service using the underlying storage
func NewProductService(st storage.Service) Product {
	return &product{
		st: st,
	}
}

type product struct {
	st storage.Service
}

//GetProducts retreives products from the underlying storage service matching the search crieteria
func (p *product) GetProducts(params SvcParams) (*model.PageOfProducts, *model.Problem) {
	sorts := parseSortKeys(params.SortKeys)
	products, count, err := p.st.GetProducts(storage.SearchParams{
		Filters:  params.Filters,
		Query:  params.Query,
		Skip:     params.From,
		Size:     params.Size,
		SortKeys: sorts,
	})
	if err != nil {
		log.WithError(err).Error("error while retrieving data from elastic search")
		return nil, &model.Problem{
			Title:  "service Unavailable",
			Status: 500,
			Detail: "error while retrieving data from elastic search",
		}
	}
	page := &model.PageOfProducts{
		TotalItems: count,
		PageNumber: params.From/params.Size + 1,
		PageSize:   len(products),
		Products:   products,
	}
	return page, nil
}

func (p *product) AddProduct(prod *model.Product) *model.Problem {
	err := p.st.AddProduct(prod)
	if err != nil {
		return &model.Problem{
			Status: 500,
			Detail: err.Error(),
			Title:  "error while getting information from the storage",
		}
	}
	return nil
}

func parseSortKeys(sortKeys map[string]bool) []storage.Sort {
	sorts := make([]storage.Sort, 0)
	for k, v := range sortKeys {
		sorts = append(sorts, storage.Sort{
			Key: k,
			Asc: v,
		})
	}
	return sorts
}
