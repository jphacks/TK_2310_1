package entity

import "database/sql/driver"

type CompanyCategory string

const (
	IT            CompanyCategory = "it"
	Manufacturing CompanyCategory = "manufacturing"
	Service       CompanyCategory = "service"
	Others        CompanyCategory = "others"
)

// 構造体 CompanyCategory に値とポインターレシーバーの両方のメソッドがあるのは公式ドキュメントがそうなっていたからです。
// 参照: https://gorm.io/ja_JP/docs/data_types.html

func (c *CompanyCategory) Scan(value interface{}) error {
	*c = CompanyCategory(value.([]byte))
	return nil
}

func (c CompanyCategory) Value() (driver.Value, error) {
	return string(c), nil
}

type Company struct {
	ID          string          `gorm:"primaryKey;type:varchar(255)"`
	DisplayName string          `gorm:"type:varchar(255);not null"`
	Description string          `gorm:"type:text"`
	CompanyURL  string          `gorm:"type:text"`
	IconURL     string          `gorm:"type:text"`
	Category    CompanyCategory `gorm:"type:company_category"`
}
