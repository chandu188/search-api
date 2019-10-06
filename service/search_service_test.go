package service

import (
	"errors"
	"testing"

	"github.com/gfg/model"
	"github.com/gfg/storage"
	"github.com/stretchr/testify/assert"
)

type inMemoryStorage struct {
	products []*model.Product
	fail     bool
}

func (ims *inMemoryStorage) GetProducts(params storage.SearchParams) ([]*model.Product, int, error) {
	if ims.fail {
		return nil, 0, errors.New("failed to load products from storage")
	}
	return ims.products, len(ims.products), nil

}

func (ims *inMemoryStorage) AddProduct(prod *model.Product) error {
	if ims.fail {
		return errors.New("failed to add products to storage")
	}
	ims.products = append(ims.products, prod)
	return nil
}

func newProduct(sku string, brand string, price float32, stock int, tags []model.Tag) *model.Product {
	return &model.Product{
		Sku:   sku,
		Brand: brand,
		Price: price,
		Stock: stock,
		Tags:  tags,
	}
}

func TestSearchService(t *testing.T) {
	/*
		Load the storage with some articles/products
	*/
	ims := &inMemoryStorage{}
	ims.products = append(ims.products, newProduct("NG1244-O12", "Nike", 24.56, 10, []model.Tag{"black_shoes"}))
	ims.products = append(ims.products, newProduct("KG1244-O12", "Kopper", 22.56, 1, []model.Tag{"white_bag"}))

	psvc := NewProductService(ims)

	testCases := []struct {
		failStorage bool
		err         bool
	}{
		{
			failStorage: true,
			err:         true,
		},
		{
			failStorage: false,
			err:         false,
		},
	}

	for _, tc := range testCases {
		m := make(map[string]string)
		ims.fail = tc.failStorage
		_, prob := psvc.GetProducts(SvcParams{
			Filters: m,
			From:    1,
			Size:    1,
		})
		assert.Equal(t, tc.err, prob != nil)
	}

}

func TestAddProductService(t *testing.T) {
	/*
		Load the storage with some articles/products
	*/
	ims := &inMemoryStorage{}
	ims.products = append(ims.products, newProduct("NG1244-O12", "Nike", 24.56, 10, []model.Tag{"black_shoes"}))
	ims.products = append(ims.products, newProduct("KG1244-O12", "Kopper", 22.56, 1, []model.Tag{"white_bag"}))

	psvc := NewProductService(ims)

	testCases := []struct {
		failStorage bool
		err         bool
	}{
		{
			failStorage: true,
			err:         true,
		},
		{
			failStorage: false,
			err:         false,
		},
	}

	for _, tc := range testCases {
		p := newProduct("TEST123-O11", "Nike", 34.5, 10, nil)
		ims.fail = tc.failStorage
		prob := psvc.AddProduct(p)
		assert.Equal(t, tc.err, prob != nil)
	}

}
