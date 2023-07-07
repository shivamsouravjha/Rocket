package server

import (
	rabbitMQ "rocket/helper/rabbitmq"

	"github.com/gin-gonic/gin"
)

func Init() {
	go rabbitMQ.Run("image")
	r := gin.New()
	r.Run(":8080")
}
