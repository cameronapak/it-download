package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

var tpl *template.Template

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
	}

	// get template from views directory
	tpl = template.Must(template.ParseGlob(path.Dir(filename) + "/templates/*"))
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", Index)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := "8080"
	portEnv := os.Getenv("PORT")
	if len(portEnv) > 0 {
		port = portEnv
	}

	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	// listen and serve on 0.0.0.0:8080 by default
	// set environment variable PORT if you want to change port
}

func Index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
	HandleError(w, err)
}
