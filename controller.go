package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"html/template"
	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)
import "net/http/pprof"

type Person struct {
	Name  string
	Phone string
}

func AttachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
}

// Handle the index page which contains a form only.
func Welcome(w http.ResponseWriter, r *http.Request) {

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

	//body, err := ioutil.ReadAll(r.Body)
	//var m map[string]interface{}
	//err = json.Unmarshal(body, &m)
	decoder := json.NewDecoder(r.Body)

	var m map[string]interface{}
	err := decoder.Decode(&m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)

	session, err := mgo.Dial("adaptivelearner:81hocyupang@54.187.83.59/learningobjects")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("learningobjects").C("learningobjects")
	err = c.Insert(m)
	if err != nil {
		panic(err)
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Your learnig object has been created! Thank you!")

}

func main() {
	r := mux.NewRouter()
	AttachProfiler(r)

	// Serve the static files here.
	r.PathPrefix("/assets").Handler(http.FileServer(http.Dir("./public/")))
	r.HandleFunc("/createlo", CreateLearningObject)
	r.HandleFunc("/", Welcome)
	http.ListenAndServe(":8080", r)
}
