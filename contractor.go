package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LastZactionHero/contact_us/models"
)

// PayloadContractor from API
type PayloadContractor struct {
	ID                int64
	Name              string  `json:"name"`
	City              string  `json:"city"`
	Phone             string  `json:"phone"`
	CurrentlyEmployed bool    `json:"currently_employed"`
	Availability      string  `json:"availability"`
	SkillIDs          []int64 `json:"skills"`
	Projects          string  `json:"projects" sql:"type:text"`
	Github            string  `json:"github"`
	Linkedin          string  `json:"linkedin"`
	Website           string  `json:"website"`
	AnythingElse      string  `json:"anything_else"`
}

func contractorCreateHandler(w http.ResponseWriter, r *http.Request) {
	applyCorsHeader(w, r)
	body, _ := ioutil.ReadAll(r.Body)
	var payload PayloadContractor
	err := json.Unmarshal(body, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var skills []models.Skill
	db.Where("id in (?)", payload.SkillIDs).Find(&skills)

	contractor := models.Contractor{
		Name:              payload.Name,
		City:              payload.City,
		Phone:             payload.Phone,
		CurrentlyEmployed: payload.CurrentlyEmployed,
		Availability:      payload.Availability,
		Projects:          payload.Projects,
		Github:            payload.Github,
		Linkedin:          payload.Linkedin,
		Website:           payload.Website,
		AnythingElse:      payload.AnythingElse,
		Skills:            skills,
	}
	db.Create(&contractor)

	w.WriteHeader(http.StatusCreated)

	triggerNotification()
}
