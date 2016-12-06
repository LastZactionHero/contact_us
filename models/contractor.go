package models

// Contractor record
type Contractor struct {
	ID                int64
	Email             string  `json:"email" validate:"min=1"`
	Name              string  `json:"name" validate:"min=1"`
	City              string  `json:"city" validate:"min=1"`
	Phone             string  `json:"phone" validate:"max=255"`
	Speciality        string  `json:"speciality" validate:"min=1"`
	CurrentlyEmployed bool    `json:"currently_employed"`
	Availability      string  `json:"availability" validate:"max=255"`
	Skills            []Skill `json:"skills" gorm:"many2many:contractor_skills;"`
	Projects          string  `json:"projects" sql:"type:text"`
	Twitter           string  `json:"twitter" validate:"max=255"`
	Github            string  `json:"github" validate:"max=255"`
	Linkedin          string  `json:"linkedin" validate:"max=255"`
	Website           string  `json:"website" validate:"max=255"`
	AnythingElse      string  `sql:"type:text"`
	Newsletter        bool    `json:"newsletter"`
}
