package entity

import (
	"database/sql/driver"
	"time"
)

// ApplicationStatus は申請の状態を表します。
// - participant: 参加者
// - absent: 不参加 (キャンセル)
type ApplicationStatus string

const (
	// ParticipantUser を Participant としていないのは、同パッケージ内の Participant 構造体と名前が被ってしまうためです。
	ParticipantUser ApplicationStatus = "participant"
	Absent          ApplicationStatus = "absent"
)

// 構造体 ApplicationStatus に値とポインターレシーバーの両方のメソッドがあるのは公式ドキュメントがそうなっていたからです。
// 参照: https://gorm.io/ja_JP/docs/data_types.html

func (a *ApplicationStatus) Scan(value interface{}) error {
	*a = ApplicationStatus(value.([]byte))
	return nil
}

func (a ApplicationStatus) Value() (driver.Value, error) {
	return string(a), nil
}

type Application struct {
	ID        string            `gorm:"type:varchar(255);primaryKey"`
	UserID    string            `gorm:"type:varchar(255);not null;references:users(id)"`
	EventID   string            `gorm:"type:varchar(255);not null;references:events(id)"`
	Status    ApplicationStatus `gorm:"type:application_status;not null;default:'participant'"`
	CreatedAt time.Time         `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time         `gorm:"not null;default:current_timestamp"`
}
