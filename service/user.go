package service

import (
	"github.com/jphacks/TK_2310_1/entity"
	DBRepository "github.com/jphacks/TK_2310_1/repository/db"
)

type IFUserService interface {
	GetUser(userID string) entity.User
}

type UserService struct {
	db DBRepository.DB
}

func NewUserService(db DBRepository.DB) IFUserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) GetUser(userID string) entity.User {
	user := u.getUserByID(userID)
	return user
}

func (u *UserService) getUserByID(userID string) entity.User {
	client := u.db.GetDB()
	var user entity.User
	client.Table("users").Select("*").Where("id = ?", userID).Find(&user)

	return user
}
