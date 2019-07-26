package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	//CODE1: import gorilla mux package for http routing

)

//CODE2: define codedetail struct

//CODE3: define language struct - provides wrapper over codedetail schema for json

//CODE4: declare a languages map for storage of our languages

//CODE5: declare createlanguage function - creates and add into languages

//CODE6: declare getlanguages function - returns all current languages

//CODE7: declare getlanguagebyname function - returns a language given its name

//CODE8: declare deletelanguagebyname function - deletes a language by name

//CODE9: declare voteonlanguage function - adds a vote to a language given its name

//CODE10: declare the init function - called automatically when application starts

func main() {
	fmt.Println("serving on port 8080!!")

	//CODE12: create a new gorilla mux router

	//CODE13: establish api routes

	//CODE14: startup the http server and configure it on port 8080 with gorilla mux router

}
