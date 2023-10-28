package apiserver

import (
	"github.com/jphacks/TK_2310_1/handler"
	"github.com/jphacks/TK_2310_1/handler/auth/signup"
	"github.com/jphacks/TK_2310_1/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type apiServerImpl struct {
	e                 *echo.Echo
	authSignupHandler signup.Signup
	eventHandler      handler.IFEventHandler
	userHandler       handler.IFUserHandler
}

func New(eventhandler handler.IFEventHandler, userHandler handler.IFUserHandler) ApiServer {
	return &apiServerImpl{
		e:                 echo.New(),
		authSignupHandler: signup.New(),
		eventHandler:      eventhandler,
		userHandler:       userHandler,
	}
}

func (a *apiServerImpl) Start() {
	// Middleware
	a.e.Use(middleware.Logger())
	a.e.Use(middleware.Recover())

	// Routes
	a.e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})
	auth := a.e.Group("")
	auth.Use(middlewares.FirebaseAuth())

	a.e.POST("/auth/signup", a.authSignupHandler.Post)
	a.e.GET("/event/:id", a.eventHandler.GetEventID)
	a.e.GET("/event/:id/participant", a.eventHandler.GetEventIDParticipant)
	a.e.GET("/event/schedule", a.eventHandler.GetEventSchedule)

	a.e.GET("/event/order-recommendation", a.eventHandler.GetOrderRecommendation)
	a.e.GET("/event/search", a.eventHandler.GetSearch)
	a.e.GET("/event/recommendation", a.eventHandler.GetEventRecommendation)
	a.e.POST("/event/:id/start", a.eventHandler.PostStartID)
	a.e.POST("/event/:id/complete", a.eventHandler.PostCompleteID)
	a.e.POST("/event/:id/report", a.eventHandler.PostReportID)
	a.e.POST("/event/:id/application", a.eventHandler.PostEventIDApplication)

	auth.GET("/user", a.userHandler.GetUserID)
	auth.POST("/user/:id/event", a.userHandler.PostUsrIDEvent)

	// Start server
	a.e.Logger.Fatal(a.e.Start(":8080"))
}
