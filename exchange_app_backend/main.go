package main

import (
	"exchange_app/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(config.AppConfig.App.Port)
	if err != nil {
		return
	}
}
