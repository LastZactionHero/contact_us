package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/LastZactionHero/contact_us/database"
	"github.com/LastZactionHero/contact_us/endpoints"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Connect to database
	database.DBConnect()
	database.DBInit()

	// Router
	serverPort := os.Getenv("CONTACT_US_PORT")
	r := mux.NewRouter()
	r.HandleFunc("/contact", endpoints.OptionsHandler).Methods("OPTIONS")
	r.HandleFunc("/contact", endpoints.ContactCreateHandler).Methods("POST")

	r.HandleFunc("/skills", endpoints.OptionsHandler).Methods("OPTIONS")
	r.HandleFunc("/skills", endpoints.SkillCreateHandler).Methods("POST")
	r.HandleFunc("/skills", endpoints.SkillIndexHandler).Methods("GET")

	r.HandleFunc("/contractors", endpoints.OptionsHandler).Methods("OPTIONS")
	r.HandleFunc("/contractors", endpoints.ContractorCreateHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
}
