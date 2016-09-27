package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LastZactionHero/contact_us/database"
	"github.com/LastZactionHero/contact_us/models"
)

type payloadContractor struct {
	ID                int64
	Email             string  `json:"email"`
	Name              string  `json:"name"`
	City              string  `json:"city"`
	Phone             string  `json:"phone"`
	Specialty         string  `json:"speciality"`
	CurrentlyEmployed bool    `json:"currently_employed"`
	Availability      string  `json:"availability"`
	SkillIDs          []int64 `json:"skills"`
	Projects          string  `json:"projects" sql:"type:text"`
	Twitter           string  `json:"twitter"`
	Github            string  `json:"github"`
	Linkedin          string  `json:"linkedin"`
	Website           string  `json:"website"`
	AnythingElse      string  `json:"anything_else"`
}

// ContractorCreateHandler POST create Contractor
func ContractorCreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ContractorCreateHandler")
	applyCorsHeader(w, r)
	body, _ := ioutil.ReadAll(r.Body)
	log.Print(body)
	var payload payloadContractor

	err := json.Unmarshal(body, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Print(payload)

	var skills []models.Skill
	database.DB.Where("id in (?)", payload.SkillIDs).Find(&skills)

	contractor := models.Contractor{
		Email:             payload.Email,
		Name:              payload.Name,
		City:              payload.City,
		Phone:             payload.Phone,
		Speciality:        payload.Specialty,
		CurrentlyEmployed: payload.CurrentlyEmployed,
		Availability:      payload.Availability,
		Projects:          payload.Projects,
		Twitter:           payload.Twitter,
		Github:            payload.Github,
		Linkedin:          payload.Linkedin,
		Website:           payload.Website,
		AnythingElse:      payload.AnythingElse,
		Skills:            skills,
	}

	log.Print(contractor)
	if errorBody := validate(contractor); errorBody != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(*errorBody)
		return
	}

	database.DB.Create(&contractor)

	w.WriteHeader(http.StatusCreated)

	triggerNotification()
}
