package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	server := http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}

	http.HandleFunc("/", Index)
	http.HandleFunc("/images", Images)
	http.HandleFunc("/individual", Individual)

	server.ListenAndServe()
}

func Index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
	HandleError(w, err)
}

func Images(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "images.tmpl.html", nil)
	HandleError(w, err)
}

func Individual(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "individual.tmpl.html", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
