package middlewares

import (
	FirebaseInfrastructure "github.com/jphacks/TK_2310_1/infrastructure/firebase"
	"github.com/jphacks/TK_2310_1/lib"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			firebaseApp := FirebaseInfrastructure.GetFirebaseApp()
			authClient, err := firebaseApp.Auth(ctx)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"message": err.Error(),
				})
			}
			barerToken, err := lib.GetAuthorizationBarerTokenFromHeader(c.Request().Header)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"message": err.Error(),
				})
			}
			token, err := authClient.VerifyIDToken(ctx, barerToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": err.Error(),
				})
			}
			log.Printf("idToken の検証に成功しました。uid -> %s", token.UID)
			// TODO : 後で直す
			c.Set("userId", "1aB2cD3eF4gH5iJ6kL7mN8oP9qR")
			return next(c)
		}
	}
}
