package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HelloWorldHandler)
	e.GET("/health", s.HealthHandler)

	return e
}

// HelloWorldHandler handles requests to the root "/" endpoint
func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{"message": "Hello World"}
	return c.JSON(http.StatusOK, resp)
}

// HealthHandler handles requests to the "/health" endpoint
func (s *Server) HealthHandler(c echo.Context) error {
	// Check database health and return the status
	return c.JSON(http.StatusOK, s.db.Health())
}
