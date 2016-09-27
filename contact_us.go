package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/LastZactionHero/contact_us/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	// Connect to database
	db = dbConnect()
	dbInit()

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

func dbConnect() *gorm.DB {
	dbUser := os.Getenv("CONTACT_US_DB_USER")
	dbPass := os.Getenv("CONTACT_US_DB_PASS")
	dbName := os.Getenv("CONTACT_US_DB_NAME")
	connectStr := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)
	dbc, err := gorm.Open("mysql", connectStr)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}
	return dbc
}

func dbInit() {
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Skill{})
	db.AutoMigrate(&models.Contractor{})
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
	db.Create(&contact)
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
