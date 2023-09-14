package main

import (
	_ "github.com/2tsumo-hitori/sample-api/config/esclient"
	"github.com/2tsumo-hitori/sample-api/router"
)

func main() {
	router.InitRouter()
}
