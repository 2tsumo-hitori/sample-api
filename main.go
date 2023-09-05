package main

import (
	"sample-api/router"
)

func main() {
	r := router.InitHandler()

	r.Run(":8080")
}
