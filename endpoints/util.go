package endpoints

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reflect"

	validator "gopkg.in/validator.v2"
)

// OptionsHandler respond to Options method
func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	applyCorsHeader(w, r)
}

func applyCorsHeader(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Origin, X-Auth-Token")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
}

func triggerNotification(title string, email string) {
	key := os.Getenv("CONTACT_US_IFTTT_KEY")
	trigger := os.Getenv("CONTACT_US_IFTTT_TRIGGER")
	url := "https://maker.ifttt.com/trigger/" + trigger + "/with/key/" + key

	type iftttValue struct {
		Title string `json:"value1"`
		Email string `json:"value2"`
	}

	value := iftttValue{Title: title, Email: email}
	valueJSON, err := json.Marshal(value)

	var reader *bytes.Reader
	if err != nil {
		log.Print("Error creating IFTTT JSON")
		log.Print(err.Error())
	} else {
		reader = bytes.NewReader(valueJSON)
	}

	_, err = http.Post(url, "application/json", reader)
	if err != nil {
		log.Printf("Error posting notification: %s", err)
	}
}

func validate(m interface{}) *[]byte {
	mt := reflect.TypeOf(m)
	if err := validator.Validate(m); err != nil {
		errors, _ := err.(validator.ErrorMap)

		errorContent := map[string]interface{}{}

		for i := 0; i < mt.NumField(); i++ {
			jsonName := mt.Field(i).Tag.Get("json")
			fieldName := mt.Field(i).Name
			errorString := errors[fieldName]
			if errorString != nil {
				errorContent[jsonName] = replaceErrorMessages(errorString)
			}
		}
		errorBody, _ := json.Marshal(errorContent)

		return &errorBody
	}
	return nil
}

func replaceErrorMessages(errorArray validator.ErrorArray) []string {
	var messages []string
	for _, defaultMessage := range errorArray {
		newMessage := replaceErrorMessage(defaultMessage.Error())
		messages = append(messages, newMessage)
	}
	return messages
}

func replaceErrorMessage(message string) string {
	switch {
	case message == "less than min":
		return "is required"
	case message == "greater than max":
		return "is too long"
	default:
		return message
	}
}
