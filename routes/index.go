package routes

import (
	"html/template"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {

	indexTemplate, err := template.New("indexTemplate").ParseFiles("templates/index.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	err = indexTemplate.ExecuteTemplate(writer, "index.html", nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

}
