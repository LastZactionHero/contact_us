package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Contact - contact form entry
type Contact struct {
	ID      int64
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message" sql:"type:text"`
}

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
	db.AutoMigrate(&Contact{})
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	applyCorsHeader(w, r)
}

func contactCreateHandler(w http.ResponseWriter, r *http.Request) {
	applyCorsHeader(w, r)
	body, _ := ioutil.ReadAll(r.Body)
	var contact Contact
	err := json.Unmarshal(body, &contact)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db.Create(&contact)
}

func applyCorsHeader(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Origin, X-Auth-Token")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
}
