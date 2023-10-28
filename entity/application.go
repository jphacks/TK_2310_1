package entity

import (
	"database/sql/driver"
	"fmt"
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
	switch v := value.(type) {
	case []byte:
		*a = ApplicationStatus(v)
	case string:
		*a = ApplicationStatus(v)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}

func (a ApplicationStatus) Value() (driver.Value, error) {
	return string(a), nil
}

type Application struct {
	UserID    string            `gorm:"type:varchar(255);not null;references:users(id);primaryKey" json:"user_id,omitempty"`
	EventID   string            `gorm:"type:varchar(255);not null;references:events(id);primaryKey" json:"event_id,omitempty"`
	Status    ApplicationStatus `gorm:"type:application_status;not null;default:'participant'" json:"status,omitempty"`
	CreatedAt time.Time         `gorm:"not null;default:current_timestamp" json:"created_at,omitempty"`
	UpdatedAt time.Time         `gorm:"not null;default:current_timestamp" json:"updated_at,omitempty"`
}
