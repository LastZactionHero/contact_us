package models

// Contact - contact form entry
type Contact struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message" sql:"type:text"`
}
