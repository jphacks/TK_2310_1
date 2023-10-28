package entity

import "time"

type Event struct {
	ID                        string    `gorm:"type:varchar(255);primaryKey"`
	Title                     string    `gorm:"type:varchar(255);not null"`
	HostCompanyID             string    `gorm:"type:varchar(255);not null;references:companies(id)"`
	Description               string    `gorm:"type:text;not null"`
	Address                   string    `gorm:"type:varchar(255);not null"`
	Latitude                  float64   `gorm:"not null"`
	Longitude                 float64   `gorm:"not null"`
	ParticipantCount          int       `gorm:"not null"`
	UnitPrice                 int       `gorm:"not null"`
	WillStartAt               time.Time `gorm:"not null"`
	WillCompleteAt            time.Time `gorm:"not null"`
	ApplicationDeadline       time.Time `gorm:"not null"`
	Leader                    string    `gorm:"type:varchar(255);references:users(id)"`
	StartedAt                 time.Time
	CompletedAt               time.Time
	ProofParticipantsImageURL string `gorm:"type:text"`
	ProofGarbageImageURL      string `gorm:"type:text"`
	Report                    string `gorm:"type:text"`
}
