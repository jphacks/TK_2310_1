package service

import (
	"context"
	"github.com/giraffe-org/backend/algo"
	"github.com/giraffe-org/backend/entity"
	DBRepository "github.com/giraffe-org/backend/repository/db"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

type IFEventService interface {
	OrderRecommendation(ctx context.Context, input InputOrderRecommendation) OutOrderRecommendation
	Search(ctx context.Context, input InputSearch) (*OutSearch, error)
	GetUserEventSchedule(uid string) ([]entity.Event, error)
}

type Event struct {
	db DBRepository.DB
}

func NewEventService(db DBRepository.DB) IFEventService {
	return &Event{
		db: db,
	}
}

type InputOrderRecommendation struct {
	Address string
	StartAt string
	EndAt   string
}

type OutOrderRecommendation struct {
	Events   []entity.Event
	SumPrice int
}

func (e *Event) OrderRecommendation(ctx context.Context, input InputOrderRecommendation) OutOrderRecommendation {
	client := e.db.GetDB()
	var event []entity.Event
	address := "%" + input.Address + "%"
	client.Table("events").Select("*").Where("address LIKE ?", address).Where("will_start_at >= ? AND will_complete_at <= ?", input.StartAt, input.EndAt).Find(&event)
	maxPrice, events := algo.Optimalplan(event)
	result := OutOrderRecommendation{
		Events:   events,
		SumPrice: maxPrice,
	}
	return result
}

type InputSearch struct {
	Keyword  string
	MinPrice int
	StartAt  string
	Limit    int
	Offset   int
}

type OutSearch struct {
	Events []entity.Event
}

func (e *Event) Search(ctx context.Context, input InputSearch) (*OutSearch, error) {
	client := e.db.GetDB()
	var events []entity.Event
	query := client.Table("events").Select("events.*, companies.display_name").
		Joins("left join companies on events.host_company_id = companies.id")
	if input.Keyword != "" {
		query = query.Where("events.title LIKE ? OR events.address LIKE ? OR companies.display_name LIKE ?", "%"+input.Keyword+"%", "%"+input.Keyword+"%", "%"+input.Keyword+"%")
	}

	query.Where("events.unit_price >= ?", input.MinPrice)

	if input.StartAt != "" {
		query.Where("events.will_start_at >= ?", input.StartAt)
	}

	err := query.Find(&events).Error
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	result := &OutSearch{
		Events: events,
	}
	return result, nil

}

func (e *Event) GetUserEventSchedule(uid string) ([]entity.Event, error) {
	var upcomingEvents []entity.Event

	// 現在の日時を取得
	currentTime := time.Now()

	// ユーザが参加予定のイベントを直近順に取得
	err := e.db.GetDB().Joins("JOIN applications on applications.event_id = events.id").
		Where("applications.user_id = ? AND applications.status = ? AND events.will_start_at > ?", uid, entity.ParticipantUser, currentTime).
		Order("events.will_start_at ASC").
		Find(&upcomingEvents).Error

	if err != nil {
		return nil, xerrors.Errorf("直近のイベントの取得に失敗しました: %w", err)
	}

	return upcomingEvents, nil
}
