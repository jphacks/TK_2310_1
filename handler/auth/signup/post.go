package signup

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jphacks/TK_2310_1/entity"
	FirebaseInfrastructure "github.com/jphacks/TK_2310_1/infrastructure/firebase"
	"github.com/jphacks/TK_2310_1/lib"
)

func (s *signupImpl) Post(c echo.Context) error {
	ctx := context.Background()
	firebaseApp := FirebaseInfrastructure.GetFirebaseApp()
	authClient, err := firebaseApp.Auth(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	barerToken, err := lib.GetAuthorizationBarerTokenFromHeader(c.Request().Header)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := authClient.VerifyIDToken(ctx, barerToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	log.Printf("idToken の検証に成功しました。uid -> %s", token.UID)

	var payload entity.User
	if err = c.Bind(&payload); err != nil {
		log.Printf("User 構造体へのリクエストボディのバインドに失敗しました: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	payload.ID = token.UID

	err = s.db.Insert(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload)
}
