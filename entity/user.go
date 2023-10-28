package entity

import "database/sql/driver"

type SexType string

const (
	Male   SexType = "male"
	Female SexType = "female"
	Other  SexType = "other"
)

// 構造体 SexType に値とポインターレシーバーの両方のメソッドがあるのは公式ドキュメントがそうなっていたからです。
// 参照: https://gorm.io/ja_JP/docs/data_types.html

func (s *SexType) Scan(value interface{}) error {
	*s = SexType(value.([]byte))
	return nil
}

func (s SexType) Value() (driver.Value, error) {
	return string(s), nil
}

type User struct {
	ID          string `gorm:"primaryKey;type:varchar(255)" json:"id,omitempty"`
	CompanyID   string `gorm:"type:varchar(255);references:companies(id)" json:"company_id,omitempty"`
	DisplayName string `gorm:"type:varchar(255);not null" json:"display_name,omitempty"`
	Age         int    `gorm:"type:integer" json:"age,omitempty"`
	IconURL     string `gorm:"type:text" json:"icon_url,omitempty"`
}
