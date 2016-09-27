package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LastZactionHero/contact_us/database"
	"github.com/LastZactionHero/contact_us/models"
)

// SkillCreateHandler POST create Skill
func SkillCreateHandler(w http.ResponseWriter, r *http.Request) {
	applyCorsHeader(w, r)
	body, _ := ioutil.ReadAll(r.Body)
	var skill models.Skill
	err := json.Unmarshal(body, &skill)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var existingSkill models.Skill
	database.DB.Where("name = ?", skill.Name).First(&existingSkill)
	if existingSkill.ID > 0 {
		// Already exists
	} else {
		database.DB.Create(&skill)
	}

	w.WriteHeader(http.StatusCreated)
}

// SkillIndexHandler GET skills
func SkillIndexHandler(w http.ResponseWriter, r *http.Request) {
	applyCorsHeader(w, r)
	var skills []models.Skill
	database.DB.Find(&skills)

	skillsByte, err := json.Marshal(skills)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(skillsByte)
	w.Header().Set("Content-Type", "application/json")
}
