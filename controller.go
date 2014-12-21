package main

import (
	//    "fmt"
	"github.com/gorilla/mux"
	//    "math"
	"html/template"
	"net/http"
	//    "path/filepath"
	"os"
)
import "net/http/pprof"

func AttachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	//    for i := 0; i < 1000000; i++ {
	//        math.Pow(36, 89)
	//    }
	//    fmt.Fprint(w, "Hello!")
	os.Getwd()
	//    fmt.Println( filepath.Join( cwd, "../view/templates/base.html" ) )
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}
	templates := template.Must(template.New("").Funcs(funcMap).ParseFiles("view/templates/base.html",
		"view/templates/view.html"))
	dat := struct {
		Title string
		Body  string
	}{
		Title: "Hello, Welcome to Xueduoduo!",
		Body:  "Welcome to <b>Xueduoduo</b>!",
	}
	err := templates.ExecuteTemplate(w, "base", dat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	AttachProfiler(r)
	r.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8080", r)
}
