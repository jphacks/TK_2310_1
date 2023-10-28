package service

import (
	DBRepository "github.com/jphacks/TK_2310_1/repository/db"
)

type IFUserService interface {
}

type User struct {
	db DBRepository.DB
}

func NewUserService(db DBRepository.DB) IFUserService {
	return &Event{
		db: db,
	}
}
