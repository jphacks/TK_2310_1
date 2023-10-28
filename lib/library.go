package lib

import (
	"golang.org/x/xerrors"
	"net/http"
	"strings"
)

func GetAuthorizationBarerTokenFromHeader(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", xerrors.Errorf("Authorization ヘッダーが設定されていません。")
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return "", xerrors.Errorf("token の形式が不正です。")
	}

	barerToken := splitToken[1]

	return barerToken, nil
}
