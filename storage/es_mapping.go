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
					"type":"keyword",
					"copy_to":"all"
				},
				"title":{
					"type":"text",
					"copy_to":"all"
				},
				"brand":{
					"type":"keyword",
					"copy_to":"all"
				},
				"tags":{
					"type":"keyword",
					"copy_to":"all"
				},
				"price":{
					"type":"double",
					"copy_to":"all"
				},
				"stock":{
					"type":"integer", 
					"copy_to":"all"
				},
				"all": {
					"type": "text"
				}
			}
		}
	}
}`
