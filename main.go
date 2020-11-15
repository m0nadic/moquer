package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m0nadic/moquer/models"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	printConfig = kingpin.Flag("print", "prints a sample configuration").Default("false").Bool()
	configFile  = kingpin.Flag("config", "specify yaml config file").Default("config.yaml").String()
	adminPort   = kingpin.Flag("admin", "admin server port").Default("8888").Int()
)

func main() {
	kingpin.Parse()
	if *printConfig {
		_ = SampleConfig().PrintYaml(os.Stdout)
		os.Exit(0)
	}


	cfg, _ := models.NewConfigFromFile(*configFile)

	for name, service := range cfg.Services {
		service.SetName(name)
		go service.Start()
	}

	StartAdmin(cfg)
}

func StartAdmin(cfg *models.Config) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(200, cfg)
	})
	r.Run(fmt.Sprintf("0.0.0.0:%d", *adminPort))
}
