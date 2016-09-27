package models

// Contractor record
type Contractor struct {
	ID                int64
	Email             string `validate:"min=1"`
	Name              string `validate:"min=1"`
	City              string `validate:"min=1"`
	Phone             string
	Speciality        string `validate:"min=1"`
	CurrentlyEmployed bool
	Availability      string
	Skills            []Skill `gorm:"many2many:contractor_skills;"`
	Projects          string  `sql:"type:text"`
	Twitter           string
	Github            string
	Linkedin          string
	Website           string
	AnythingElse      string `sql:"type:text"`
}
