package main

import "github.com/2tsumo-hitori/sample-api/router"
import "github.com/2tsumo-hitori/sample-api/config/esclient"

func main() {
	esclient.InitElasticSearch()
	router.InitRouter()
}
