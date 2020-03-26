package main

import (
	"connector/router"
	"connector/service"
)

func main() {

	r := router.InitRouter()

	service.InitGateway()

	r.Run(":8282")
}
