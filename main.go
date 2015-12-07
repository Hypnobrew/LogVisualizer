package main

import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
	"LogVisualizer/models"
)

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(writer http.ResponseWriter, request *http.Request)  {
	page := &models.HomeModel{Title: "Hemma"}
	homeTemplate, err := template.ParseFiles("views/home.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	err = homeTemplate.Execute(writer, page)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}