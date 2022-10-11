package main

import (
	"github.com/luankkobs/goweb/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
