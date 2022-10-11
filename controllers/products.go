package controllers

import (
	"github.com/luankkobs/goweb/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price to float64", err)
		}
		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity to int", err)
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedQuantity)
	}
	http.Redirect(w, r, "/", 301)
}
