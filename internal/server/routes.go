package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"server-template/internal/database"
)

var db database.Service

type userData struct {
	Username string `json:"username" validate:"required"`    // Add validation with a tag
	Email    string `json:"email" validate:"required,email"` // Add validation with tags
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HelloWorldHandler)
	e.POST("/user", s.AddUser)
	e.GET("/health", s.HealthHandler)

	return e
}

// HelloWorldHandler handles requests to the root "/" endpoint
func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{"message": "Hello World"}
	return c.JSON(http.StatusOK, resp)
}

// AddUser example POST to add a user "/user"
func (s *Server) AddUser(c echo.Context) error {
	var user userData
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body format"})
	}

	userResponse := db.AddUser(user.Username, user.Email)
	if userResponse["error"] != "" {
		return c.JSON(http.StatusInternalServerError, userResponse)
	}

	return c.JSON(http.StatusCreated, userResponse)
}

// HealthHandler handles requests to the "/health" endpoint
func (s *Server) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
