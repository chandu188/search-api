package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gfg/model"
	elastic "gopkg.in/olivere/elastic.v6"
)

type elasticSearch struct {
	client *elastic.Client
	index  string
	typ    string
}

// NewElasticSearch returns a Storage service backed by elastic Search
func NewElasticSearch(addrs []string, index string, typ string) Service {
	ctx := context.Background()

	client, err := elastic.NewSimpleClient(elastic.SetURL(addrs...))
	if err != nil {
		panic(err)
	}

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	return &elasticSearch{client: client, index: index, typ: typ}
}

func (es *elasticSearch) GetProducts(params SearchParams) ([]*model.Product, int, error) {
	ctx := context.Background()

	bQuery := elastic.NewBoolQuery()
	for key, val := range params.Filters {
		tQuery := elastic.NewTermQuery(key, val)
		bQuery.Filter(tQuery)
	}

	for key, val := range params.Queries {
		mQuery := elastic.NewMatchQuery(key, val)
		bQuery.Must(mQuery)
	}

	searchSvc := es.client.Search(es.index).Type(es.typ).
		Query(bQuery).From(params.Skip).Size(params.Size)

	for _, s := range params.SortKeys {
		searchSvc = searchSvc.Sort(s.Key, s.Asc)
	}

	res, err := searchSvc.Do(ctx)

	if err != nil {
		return nil, 0, err
	}

	products := make([]*model.Product, 0)
	// Iterate through results
	for _, hit := range res.Hits.Hits {
		var p model.Product
		err := json.Unmarshal(*hit.Source, &p)
		if err != nil {
			fmt.Printf("error while deserializing json: %s", err.Error())
		}
		products = append(products, &p)
	}

	return products, int(res.Hits.TotalHits), nil
}

func (es *elasticSearch) AddProduct(prod *model.Product) error {
	res, err := es.client.Index().
		Index(es.index).Type(es.typ).
		BodyJson(prod).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Indexed product %s to index %s, type %s\n", res.Id, res.Index, res.Type)
	return nil
}
