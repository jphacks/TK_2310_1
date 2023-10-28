package handler

import (
	DBRepository "github.com/jphacks/TK_2310_1/repository/db"
	"github.com/jphacks/TK_2310_1/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IFUserHandler interface {
	GetUserID(c echo.Context) error
}

func NewUserHandler(db DBRepository.DB, service service.IFUserService) IFUserHandler {
	return &UserHandler{
		db:          db,
		userService: service,
	}
}

type UserHandler struct {
	db          DBRepository.DB
	userService service.IFUserService
}

func (u *UserHandler) GetUserID(c echo.Context) error {
	userId := c.Get("userId").(string)
	user := u.userService.GetUser(userId)

	return c.JSON(http.StatusOK, user)
}
