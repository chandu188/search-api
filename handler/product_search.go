package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gfg/model"
	"github.com/gfg/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ProductHandler represents a handler for handling the various endpoints related to product
type ProductHandler struct {
	service service.Product
}

//NewProductHandler creates a ProductHandler
func NewProductHandler(svc service.Product) *ProductHandler {
	return &ProductHandler{
		service: svc,
	}
}

// GetProducts is handler for handling the products related search queries
func (ph *ProductHandler) GetProducts(ctx *gin.Context) {
	filters := parseKeyValue(ctx.QueryArray("filter"))
	query := ctx.Query("q")

	ss := ctx.QueryArray("sort_by")
	sortKeys := make(map[string]bool)
	for _, e := range ss {
		parts := strings.Split(e, ":")
		var asc bool
		if len(parts) == 2 {
			asc = parts[1] == "asc"
		}
		sortKeys[parts[0]] = asc
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	pp, prob := ph.service.GetProducts(service.SvcParams{
		Filters:  filters,
		Query:    query,
		From:     offset,
		Size:     limit,
		SortKeys: sortKeys,
	})
	if prob != nil {
		ctx.JSON(prob.Status, prob)
		return
	}
	ctx.JSON(200, pp)
}

// AddProduct is handler for adding product
func (ph *ProductHandler) AddProduct(ctx *gin.Context) {
	var prod model.Product
	err := ctx.BindJSON(&prod)
	if err != nil {
		ctx.JSON(400, &model.Problem{
			Detail: "error while parsing product information",
			Status: 400,
			Title:  "Bad Request",
		})
		log.WithError(err).Error("error while converting body to Product")
		return
	}
	prob := ph.service.AddProduct(&prod)
	if prob != nil {
		ctx.JSON(prob.Status, prob)
	}
	ctx.JSON(201, http.NoBody)
}

func parseKeyValue(strs []string) map[string]string {
	filters := make(map[string]string)
	for _, e := range strs {
		parts := strings.Split(e, ":")
		if len(parts) == 2 {
			filters[parts[0]] = parts[1]
		}
	}
	return filters
}
