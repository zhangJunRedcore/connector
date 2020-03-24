package main

import (
	"connector/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":8282")
}
