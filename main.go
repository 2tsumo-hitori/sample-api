package main

import (
	"sample-api/config/esclient"
	"sample-api/router"
)

func main() {
	esclient.InitElasticSearch()
	router.InitRouter()
}
