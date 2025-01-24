package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {
	templates = template.Must(templates.ParseGlob("html/templates/*.html"))
	templates = template.Must(templates.ParseGlob("html/pages/*.html"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderHome)
	mux.HandleFunc("/about", renderAboutUs)
	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func renderHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Home",
	}
	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println("Error rendering home page:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderAboutUs(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "About",
	}
	templates.ExecuteTemplate(w, "about.html", data)
}
