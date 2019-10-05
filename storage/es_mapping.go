package storage

/*
	Sku   string  `json:"sku"`
	Title string  `json:"title,omitempty"`
	Brand string  `json:"brand"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
	Tags  []Tag   `json:"tags"`

*/
const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"product":{
			"properties":{
				"sku":{
					"type":"keyword"
				},
				"title":{
					"type":"text"
				},
				"brand":{
					"type":"keyword"
				},
				"tags":{
					"type":"keyword"
				},
				"price":{
					"type":"double"
				},
				"stock":{
					"type":"integer"
				}
			}
		}
	}
}`
