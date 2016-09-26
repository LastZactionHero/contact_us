package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Skill possessed by a contractor
type Skill struct {
	ID   int64
	Name string `json:"name"`
}

func skillCreateHandler(w http.ResponseWriter, r *http.Request) {
	applyCorsHeader(w, r)
	body, _ := ioutil.ReadAll(r.Body)
	var skill Skill
	err := json.Unmarshal(body, &skill)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db.Create(&skill)
}

func skillIndexHandler(w http.ResponseWriter, r *http.Request) {
	applyCorsHeader(w, r)
	var skills []Skill
	db.Find(&skills)

	skillsByte, err := json.Marshal(skills)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(skillsByte)
	w.Header().Set("Content-Type", "application/json")
}
