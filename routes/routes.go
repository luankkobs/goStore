package routes

import (
	"github.com/luankkobs/goweb/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
