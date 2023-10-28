package lib

import (
	"github.com/google/uuid"
)

/*
NewUUID 新たなUUID(v4)を作成

	36文字のstring形式で返します
*/
func NewUUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}
