package main

import (
	"net/http"

	"github.com/gfg/handler"
	"github.com/gfg/middleware"
	"github.com/gfg/service"
	"github.com/gfg/storage"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	ESConfig
}

type ESConfig struct {
	Address []string
	Index   string
	Typ     string
}

func main() {

	viper.SetConfigFile("../config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Fatal("error while reading configuration file")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.WithError(err).Panic("error while deserailizing the configuration file")
		log.Debug()
	}

	engine := gin.Default()
	engine.LoadHTMLFiles("../templates/stock.html")
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "stock.html", gin.H{
			"title": "Main website",
		})
	})

	st := storage.NewElasticSearch(cfg.Address, cfg.Index, cfg.Typ)
	svc := service.NewProductService(st)
	h := handler.NewProductHandler(svc)

	v1Group := engine.Group("/api/v1")
	routes := v1Group.Use(middleware.BasicAuth())
	routes.GET("/products", h.GetProducts)
	routes.POST("/products", h.AddProduct)

	engine.Run(cfg.ServerAddress)

}
