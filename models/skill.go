package models

// Validatable model interface
type Validatable interface {
}

// Skill possessed by a contractor
type Skill struct {
	ID   int64  `json:"id"`
	Name string `json:"name" validate:"min=1"`
}
