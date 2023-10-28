package handler

import (
	DBRepository "github.com/jphacks/TK_2310_1/repository/db"
	"github.com/jphacks/TK_2310_1/service"
)

type IFUserHandler interface {
}

func NewUserHandler(db DBRepository.DB, service service.IFUserService) IFUserHandler {
	return &UserHandler{
		db:           db,
		eventService: service,
	}
}

type UserHandler struct {
	db           DBRepository.DB
	eventService service.IFUserService
}
