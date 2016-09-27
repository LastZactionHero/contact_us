package models

// Contractor record
type Contractor struct {
	ID                int64
	Name              string
	City              string
	Phone             string
	CurrentlyEmployed bool
	Availability      string
	Skills            []Skill `gorm:"many2many:contractor_skills;"`
	Projects          string  `sql:"type:text"`
	Github            string
	Linkedin          string
	Website           string
	AnythingElse      string `sql:"type:text"`
}
