package service

import (
	"github.com/jphacks/TK_2310_1/entity"
	"github.com/jphacks/TK_2310_1/lib"
	DBRepository "github.com/jphacks/TK_2310_1/repository/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IFUserService interface {
	GetUser(userID string) entity.User
	CreateEvent(input InputCreateEvent) (*event, error)
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

type InputCreateEvent struct {
	ID                  string
	Title               string
	Description         string
	Address             string
	Latitude            float64
	Longitude           float64
	ParticipantCount    int
	UnitPrice           int
	WillStartAt         string
	WillCompleteAt      string
	ApplicationDeadline string
}

type event struct {
	ID                  string
	Title               string
	HostCompanyID       string
	Description         string
	Address             string
	Latitude            float64
	Longitude           float64
	ParticipantCount    int
	UnitPrice           int
	WillStartAt         string
	WillCompleteAt      string
	ApplicationDeadline string
}

func (u *UserService) CreateEvent(input InputCreateEvent) (*event, error) {
	client := u.db.GetDB()
	user := u.getUserByID(input.ID)
	if user.CompanyID == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest)
	}
	uuid := lib.NewUUID()
	event := event{
		ID:                  uuid,
		Title:               input.Title,
		HostCompanyID:       user.CompanyID,
		Description:         input.Description,
		Address:             input.Address,
		Latitude:            input.Latitude,
		Longitude:           input.Longitude,
		ParticipantCount:    input.ParticipantCount,
		UnitPrice:           input.UnitPrice,
		WillStartAt:         input.WillStartAt,
		WillCompleteAt:      input.WillCompleteAt,
		ApplicationDeadline: input.ApplicationDeadline,
	}
	client.Create(&event)
	return &event, nil
}

func (u *UserService) getUserByID(userID string) entity.User {
	client := u.db.GetDB()
	var user entity.User
	client.Table("users").Select("*").Where("id = ?", userID).Find(&user)

	return user
}
