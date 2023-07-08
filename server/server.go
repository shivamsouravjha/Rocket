package server

import (
	rabbitMQ "rocket/helper/rabbitmq"
	"rocket/routes"
)

func Init() {
	go rabbitMQ.Run("image")
	r := routes.NewRouter()
	r.Run(":8080")
}
