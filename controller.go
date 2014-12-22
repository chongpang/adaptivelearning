package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)
import "net/http/pprof"

func AttachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
}

// Handle the index page which contains a form only.
func Index(w http.ResponseWriter, r *http.Request) {

	// Just for debugging.
	cwd, _ := os.Getwd()
	fmt.Println(filepath.Join(cwd, "view/index.html"))

	templates, err := template.ParseFiles(filepath.Join(cwd, "view/index.html"))
	if err != nil {
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	templates.ExecuteTemplate(w, "welcome", nil)
}

// Handle the create learning object request.
func CreateLearningObject(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	templates, err := template.ParseFiles(filepath.Join(cwd, "view/created.html"))
	if err != nil {
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	templates.ExecuteTemplate(w, "created", nil)
}

func main() {
	r := mux.NewRouter()
	AttachProfiler(r)

	// Serve the static files here.
	r.PathPrefix("/assets").Handler(http.FileServer(http.Dir("./public/")))
	r.HandleFunc("/createlo", CreateLearningObject)
	r.HandleFunc("/", Index)

	http.ListenAndServe(":8080", r)
}
