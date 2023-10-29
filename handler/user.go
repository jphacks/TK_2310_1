package handler

import (
	"github.com/jphacks/TK_2310_1/api/gen"
	DBRepository "github.com/jphacks/TK_2310_1/repository/db"
	"github.com/jphacks/TK_2310_1/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type IFUserHandler interface {
	GetUserID(c echo.Context) error
	PostUsrIDEvent(c echo.Context) error
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
	log.Println(userId)
	user := u.userService.GetUser(userId)

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) PostUsrIDEvent(c echo.Context) error {
	userId := c.Get("userId").(string)
	var req gen.PostUserIdEventRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "requestのBindに失敗しました：", err)
	}
	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "requestが不正です：", err)
	}

	input := service.InputCreateEvent{
		ID:                  userId,
		Title:               req.Title,
		Description:         req.Description,
		Address:             req.Address,
		Latitude:            req.Latitude,
		Longitude:           req.Longitude,
		ParticipantCount:    int(req.ParticipantCount),
		UnitPrice:           int(req.UnitPrice),
		WillStartAt:         req.WillStartAt,
		WillCompleteAt:      req.WillCompleteAt,
		ApplicationDeadline: req.ApplicationDeadline,
	}
	event, err := u.userService.CreateEvent(input)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &event)
}
