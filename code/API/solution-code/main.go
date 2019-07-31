package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	//CODE1: import gorilla mux and handlers packages for http routing and CORS support
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//CODE2: define codedetail struct
type codedetail struct {
	Usecase  string `json:"usecase,omitempty"`
	Rank     int    `json:"rank,omitempty"`
	Compiled bool   `json:"compiled"`
	Homepage string `json:"homepage,omitempty"`
	Download string `json:"download,omitempty"`
	Votes    int    `json:"votes"`
}

//CODE3: define language struct - provides wrapper over codedetail schema for json
type language struct {
	Name   string     `json:"name,omitempty"`
	Detail codedetail `json:"codedetail,omitempty"`
}

//CODE4: declare a languages map for storage of our languages
var languages = make(map[string]*codedetail)

//CODE5: declare createlanguage function - creates and add into languages
func createlanguage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var detail codedetail
	_ = json.NewDecoder(req.Body).Decode(&detail)
	name := strings.ToLower(params["name"])
	languages[name] = &detail

	err := json.NewEncoder(w).Encode(detail)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

//CODE6: declare getlanguages function - returns all current languages
func getlanguages(w http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(w).Encode(languages)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	return
}

//CODE7: declare getlanguagebyname function - returns a language given its name
func getlanguagebyname(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])

	if detail, ok := languages[name]; ok {
		language := language{
			Name:   name,
			Detail: *detail,
		}

		err := json.NewEncoder(w).Encode(language)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		return
	}

	_ = json.NewEncoder(w).Encode("{'result' : 'language not found'}")
}

//CODE8: declare deletelanguagebyname function - deletes a language by name
func deletelanguagebyname(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])
	delete(languages, name)

	err := json.NewEncoder(w).Encode(languages)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	return
}

//CODE9: declare voteonlanguage function - adds a vote to a language given its name
func voteonlanguage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])

	fmt.Println("incoming vote for: " + name)

	if detail, ok := languages[name]; ok {
		detail.Votes++

		language := language{
			Name:   name,
			Detail: *detail,
		}

		//json.NewEncoder(w).Encode("{'ok' : 1}")
		err := json.NewEncoder(w).Encode(language)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		return
	}

	// if not found return empty object with Language structure
	_ = json.NewEncoder(w).Encode("{'result' : 'language not found'}")
}

//CODE10: declare the init function - called automatically when application starts
func init() {
	goDetail := codedetail{
		Usecase:  "system, web, server-side",
		Rank:     16,
		Compiled: true,
		Homepage: "https://golang.org",
		Download: "https://golang.org/dl/",
	}
	javaDetail := codedetail{
		Usecase:  "system, web, server-side",
		Rank:     2,
		Compiled: true,
		Homepage: "https://www.java.com/en/",
		Download: "https://www.java.com/en/download/",
	}
	nodejsDetail := codedetail{
		Usecase:  "system, web, server-side",
		Rank:     30,
		Compiled: false,
		Homepage: "https://nodejs.org/en/",
		Download: "https://nodejs.org/en/download/",
	}
	javascriptDetail := codedetail{
		Usecase:  "web, frontend development",
		Rank:     1,
		Compiled: false,
		Homepage: "https://en.wikipedia.org/wiki/JavaScript",
	}

	languages["go"] = &goDetail
	languages["java"] = &javaDetail
	languages["nodejs"] = &nodejsDetail
	languages["javascript"] = &javascriptDetail
}

func main() {
	fmt.Println("serving on port 8080!!")

	//CODE11: create a new gorilla mux router
	router := mux.NewRouter()

	//CODE12: establish api routes
	router.HandleFunc("/languages/{name}", createlanguage).Methods("POST")
	router.HandleFunc("/languages", getlanguages).Methods("GET")
	router.HandleFunc("/languages/{name}", getlanguagebyname).Methods("GET")
	router.HandleFunc("/languages/{name}", deletelanguagebyname).Methods("DELETE")
	router.HandleFunc("/languages/{name}/vote", voteonlanguage).Methods("GET")

	//CODE13: configured CORS settings for incoming AJAX requests
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET"})

	//CODE14: startup the http server and configure it on port 8080 with the gorilla mux router
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
