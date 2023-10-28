package handler

import (
	"context"
	"errors"
	"github.com/jphacks/TK_2310_1/api/gen"
	"github.com/jphacks/TK_2310_1/entity"
	FirebaseInfrastructure "github.com/jphacks/TK_2310_1/infrastructure/firebase"
	"github.com/jphacks/TK_2310_1/lib"
	DBRepository "github.com/jphacks/TK_2310_1/repository/db"
	"github.com/jphacks/TK_2310_1/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type IFEventHandler interface {
	GetOrderRecommendation(c echo.Context) error
	GetEventSchedule(c echo.Context) error
	GetEventID(c echo.Context) error
	GetSearch(c echo.Context) error
	PostStartID(c echo.Context) error
	PostCompleteID(c echo.Context) error
	PostReportID(c echo.Context) error
	GetEventIDParticipant(c echo.Context) error
}

func NewEventHandler(db DBRepository.DB, service service.IFEventService) IFEventHandler {
	return &EventHandler{
		db:           db,
		eventService: service,
	}
}

type EventHandler struct {
	db           DBRepository.DB
	eventService service.IFEventService
}

func (e *EventHandler) GetOrderRecommendation(c echo.Context) error {
	ctx := c.Request().Context()

	address := c.QueryParam("address")
	start := c.QueryParam("start_at")
	end := c.QueryParam("complete_at")

	input := service.InputOrderRecommendation{
		Address: address,
		StartAt: start,
		EndAt:   end,
	}
	events := e.eventService.OrderRecommendation(ctx, input)

	return c.JSON(http.StatusOK, events)
}

func (e *EventHandler) GetEventSchedule(c echo.Context) error {
	ctx := context.Background()
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

	events, err := e.eventService.GetUserEventSchedule(token.UID)

	//uid := "5qR8s9T0u1V2w3X4y5Z6a7B8c9D"
	//events, err := e.eventService.GetUserEventSchedule(uid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, events)
}

func (e *EventHandler) GetEventID(c echo.Context) error {
	ctx := context.Background()
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

	eventID := c.Param("id")

	var event entity.Event
	if err := e.db.GetDB().Where("id = ?", eventID).First(&event).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Event not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error fetching event",
		})
	}
	return c.JSON(http.StatusOK, event)
}

func (e *EventHandler) GetSearch(c echo.Context) error {

	ctx := c.Request().Context()

	keyword := c.QueryParam("keyword")
	minUnitPrice := c.QueryParam("min_unit_price")
	minPrice := 0
	if minUnitPrice != "" {
		minPrice, _ = strconv.Atoi(minUnitPrice)
	}
	willStartAt := c.QueryParam("will_start_at")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	input := service.InputSearch{
		keyword,
		minPrice,
		willStartAt,
		limit,
		offset,
	}
	events, err := e.eventService.Search(ctx, input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &events)
}

func (e *EventHandler) PostStartID(c echo.Context) error {
	ctx := context.Background()
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

	eventID := c.Param("id")

	intput := service.InputPostStartID{
		Id:      token.UID,
		EventID: eventID,
	}
	err = e.eventService.PostStartID(intput)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (e *EventHandler) PostCompleteID(c echo.Context) error {
	ctx := context.Background()
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

	eventID := c.Param("id")

	var req gen.PostEventIdCompleteRequest
	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "requestのBindに失敗しました：", err)
	}
	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "requestが不正です：", err)
	}

	input := service.InputPostCompleteID{
		Id:                        token.UID,
		EventID:                   eventID,
		ProofParticipantsImageUrl: req.ProofParticipantsImageUrl,
		ProofGarbageImageUrl:      req.ProofGarbageImageUrl,
		Report:                    req.Report,
	}
	err = e.eventService.PostCompleteID(input)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (e *EventHandler) PostReportID(c echo.Context) error {

	ctx := context.Background()
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

	eventID := c.Param("id")

	intput := service.InputPostStartID{
		Id:      token.UID,
		EventID: eventID,
	}
	err = e.eventService.PostReportID(intput)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (e *EventHandler) GetEventIDParticipant(c echo.Context) error {
	eventID := c.Param("id")
	var participants []entity.Participant

	if err := e.db.GetDB().Where("event_id = ? AND status = 'not_completed'", eventID).Find(&participants).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch participants",
		})
	}

	type participantResponse struct {
		UserID string `json:"user_id"`
		Status string `json:"status"`
	}
	respParticipants := []participantResponse{}
	for _, p := range participants {
		respParticipants = append(respParticipants, participantResponse{
			UserID: p.UserID,
			Status: string(p.Status),
		})
	}

	return c.JSON(http.StatusOK, struct {
		Participants []participantResponse `json:"participants"`
	}{
		Participants: respParticipants,
	})

}
