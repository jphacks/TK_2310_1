package entity

import "database/sql/driver"

type ParticipationStatus string

const (
	NotCompleted ParticipationStatus = "not_completed"
	Completed    ParticipationStatus = "completed"
)

// 構造体 ParticipationStatus に値とポインターレシーバーの両方のメソッドがあるのは公式ドキュメントがそうなっていたからです。
// 参照: https://gorm.io/ja_JP/docs/data_types.html

func (p *ParticipationStatus) Scan(value interface{}) error {
	*p = ParticipationStatus(value.([]byte))
	return nil
}

func (p ParticipationStatus) Value() (driver.Value, error) {
	return string(p), nil
}

type Participant struct {
	UserID  string              `gorm:"type:varchar(255);primaryKey;references:users(id)"`
	EventID string              `gorm:"type:varchar(255);primaryKey;references:events(id)"`
	Status  ParticipationStatus `gorm:"type:participation_status;not null;default:'not_completed'"`
}
