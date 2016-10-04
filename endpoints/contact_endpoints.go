package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LastZactionHero/contact_us/database"
	"github.com/LastZactionHero/contact_us/models"
)

// ContactCreateHandler POST create Contact
func ContactCreateHandler(w http.ResponseWriter, r *http.Request) {
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

	triggerNotification("Contact Us", contact.Email)
}
