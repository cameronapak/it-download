package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/heroku/x/hmetrics/onload"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func main() {
	server := http.Server{
		Addr:        ":" + os.Getenv("PORT"),
		IdleTimeout: time.Hour,
	}

	http.HandleFunc("/", Index)

	server.ListenAndServe()
}

func Index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
	HandleError(w, err)
}
