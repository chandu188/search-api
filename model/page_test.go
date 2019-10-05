package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage(t *testing.T) {
	sku := "NG1244-O12"
	brand := "Nike"
	p := &Product{
		Sku:   sku,
		Brand: brand,
	}
	prods := &PageOfProducts{
		TotalItems: 1,
		PageNumber: 1,
		PageSize:   10,
		Products:   []*Product{p},
	}

	_, err := json.Marshal(prods)
	assert.Nil(t, err)
}
