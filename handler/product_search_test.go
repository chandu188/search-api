package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gfg/model"
	"github.com/gfg/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyProductSvc struct {
	fail bool
}

func (d *dummyProductSvc) GetProducts(params service.SvcParams) (*model.PageOfProducts, *model.Problem) {
	if d.fail {
		return nil, &model.Problem{
			Status: 500,
			Detail: "error to test the handler",
		}
	}

	product := &model.Product{
		Sku:   "NIKE123-I12",
		Brand: "Nike",
		Stock: 1,
		Price: 34.56,
	}
	prods := []*model.Product{product}
	return &model.PageOfProducts{
		PageNumber: 1,
		PageSize:   1,
		TotalItems: 1,
		Products:   prods,
	}, nil
}

func (d *dummyProductSvc) AddProduct(*model.Product) *model.Problem {
	if d.fail {
		return &model.Problem{
			Status: 500,
			Detail: "error to test the handler",
		}
	}
	return nil
}

func TestProductSearchHandler(t *testing.T) {
	tests := []struct {
		url     string
		body    string
		resp    int
		failSvc bool
	}{
		{
			url:  "/v1/products/",
			resp: 200,
			body: "sampleOutput",
		},
		{
			url:     "/v1/products/",
			resp:    500,
			body:    "sampleOutput",
			failSvc: true,
		},
	}

	dps := &dummyProductSvc{}
	hdlr := NewProductHandler(dps)
	for _, tc := range tests {
		dps.fail = tc.failSvc
		req := httptest.NewRequest("GET", tc.url, bytes.NewReader([]byte(tc.body)))
		resp := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(resp)
		ctx.Request = req
		hdlr.GetProducts(ctx)
		assert.Equal(t, tc.resp, resp.Result().StatusCode)
	}
}

func TestProductAddHandler(t *testing.T) {
	tests := []struct {
		url     string
		body    string
		resp    int
		failSvc bool
	}{
		{
			url:  "/v1/products/",
			resp: 400,
			body: "sampleOutput",
		},
		{
			url:  "/v1/products/",
			resp: 201,
			body: `{
				"sku":"NIKE123-O12",
				"stock": 10,
				"tags": ["black_shoes"]
			}`,
			failSvc: false,
		},
		{
			url:  "/v1/products/",
			resp: 500,
			body: `{
				"sku":"NIKE123-O12",
				"stock": 10,
				"tags": ["black_shoes"]
			}`,
			failSvc: true,
		},
	}
	dps := &dummyProductSvc{}
	hdlr := NewProductHandler(dps)
	for _, tc := range tests {
		dps.fail = tc.failSvc
		req := httptest.NewRequest("POST", tc.url, bytes.NewReader([]byte(tc.body)))
		resp := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(resp)
		ctx.Request = req
		hdlr.AddProduct(ctx)
		assert.Equal(t, tc.resp, resp.Result().StatusCode)
	}
}
