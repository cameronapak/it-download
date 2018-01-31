package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jessebarton/image-get/install"
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
	http.HandleFunc("/image", Image)

	server.ListenAndServe()
}

func Index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
	HandleError(w, err)
}

func Image(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("i")
	install.Install(v)
	err := tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
