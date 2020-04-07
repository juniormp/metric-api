package main

import (
	"github.com/juniormp/metric-api/src/web"
)

func main() {
	router := web.SetupRouter()
	router.Run()
}
