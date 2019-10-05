package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	p := &Product{
		Sku:   "NG1244-O12",
		Brand: "Nike",
		Tags:  []Tag{"black_shoes"},
	}
	_, err := json.Marshal(p)
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestProductUnMarshal(t *testing.T) {
	sku := "NG1244-O12"
	brand := "Nike"

	productJSON := fmt.Sprintf(`{
		"sku" : "%s",
		"brand" : "%s",
		"tags" : ["black_shoes"]
	}`, sku, brand)

	var p Product
	err := json.Unmarshal([]byte(productJSON), &p)
	assert.Nil(t, err)
	assert.Equal(t, sku, p.Sku)
	assert.Equal(t, brand, p.Brand)
	assert.Equal(t, 1, len(p.Tags))
}
