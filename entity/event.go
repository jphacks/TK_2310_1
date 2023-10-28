package entity

import "time"

type Event struct {
	ID                        string    `gorm:"type:varchar(255);primaryKey" json:"id"`
	Title                     string    `gorm:"type:varchar(255);not null" json:"title,omitempty"`
	HostCompanyID             string    `gorm:"type:varchar(255);not null;references:companies(id)" json:"host_company_id,omitempty"`
	Description               string    `gorm:"type:text;not null" json:"description,omitempty"`
	Address                   string    `gorm:"type:varchar(255);not null" json:"address,omitempty"`
	Latitude                  float64   `gorm:"not null" json:"latitude,omitempty"`
	Longitude                 float64   `gorm:"not null" json:"longitude,omitempty"`
	ParticipantCount          int       `gorm:"not null" json:"participant_count,omitempty"`
	UnitPrice                 int       `gorm:"not null" json:"unit_price,omitempty"`
	WillStartAt               time.Time `gorm:"not null" json:"will_start_at,omitempty"`
	WillCompleteAt            time.Time `gorm:"not null" json:"will_complete_at,omitempty"`
	ApplicationDeadline       time.Time `gorm:"not null" json:"application_deadline,omitempty"`
	Leader                    string    `gorm:"type:varchar(255);references:users(id)" json:"leader,omitempty"`
	StartedAt                 time.Time `json:"-"`
	CompletedAt               time.Time `json:"-"`
	ProofParticipantsImageURL string    `gorm:"type:text" json:"proof_participants_image_url,omitempty"`
	ProofGarbageImageURL      string    `gorm:"type:text" json:"proof_garbage_image_url,omitempty"`
	Report                    string    `gorm:"type:text" json:"report,omitempty"`
}
