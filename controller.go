package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"html/template"
	//"io/ioutil"
	"github.com/syabondama/adaptivelearning/models"
	//"log"
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

	// Decode json data
	var m map[string]interface{}
	err := decoder.Decode(&m)
	if err != nil {
		panic(err)
	}

	id, err := models.SaveLODocObj(m)
	// Just for testing
	fmt.Println(m)
	models.CreateGraphNodeAndRelationships(m, id)
	fmt.Fprintf(w, "Your learnig object has been created! Thank you!")

}

// Return a list of learning object id and title
func GetLearningObjectsIds(w http.ResponseWriter, r *http.Request) {
	ids, err := models.GetLearningObjectsIds()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.Write(ids)
	} else {
		w.Write([]byte("Error occurs"))
	}
}

func main() {
	r := mux.NewRouter()
	AttachProfiler(r)

	// Serve the static files here.
	r.PathPrefix("/assets").Handler(http.FileServer(http.Dir("./public/")))
	r.HandleFunc("/createlo", CreateLearningObject)
	r.HandleFunc("/getlolist", GetLearningObjectsIds)
	r.HandleFunc("/", Welcome)
	http.ListenAndServe(":8080", r)
}
