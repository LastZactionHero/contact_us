package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/LastZactionHero/contact_us/database"
	"github.com/LastZactionHero/contact_us/models"
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
	r.HandleFunc("/contact", optionsHandler).Methods("OPTIONS")
	r.HandleFunc("/contact", contactCreateHandler).Methods("POST")

	r.HandleFunc("/skills", optionsHandler).Methods("OPTIONS")
	r.HandleFunc("/skills", skillCreateHandler).Methods("POST")
	r.HandleFunc("/skills", skillIndexHandler).Methods("GET")

	r.HandleFunc("/contractors", optionsHandler).Methods("OPTIONS")
	r.HandleFunc("/contractors", contractorCreateHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	applyCorsHeader(w, r)
}

func contactCreateHandler(w http.ResponseWriter, r *http.Request) {
	applyCorsHeader(w, r)
	body, _ := ioutil.ReadAll(r.Body)
	var contact models.Contact
	err := json.Unmarshal(body, &contact)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	database.DB.Create(&contact)
	w.WriteHeader(http.StatusCreated)

	triggerNotification()
}

func applyCorsHeader(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Origin, X-Auth-Token")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
}

type smtpTemplateData struct {
	From    string
	To      string
	Subject string
	Body    string
}

func triggerNotification() {
	key := os.Getenv("CONTACT_US_IFTTT_KEY")
	trigger := os.Getenv("CONTACT_US_IFTTT_TRIGGER")
	url := "https://maker.ifttt.com/trigger/" + trigger + "/with/key/" + key
	_, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Printf("Error posting notification: %s", err)
	}
}
