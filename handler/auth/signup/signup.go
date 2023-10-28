package signup

import "github.com/labstack/echo/v4"

type Signup interface {
	Post(c echo.Context) error
}
