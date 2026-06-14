package main

import (
  "github.com/labstack/echo/v5"
  "github.com/labstack/echo/v5/middleware"
  "log/slog"
  "net/http"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.RequestLogger()) // use the RequestLogger middleware with slog logger
  e.Use(middleware.Recover())       // recover panics as errors for proper error handling

  // Routes
  e.GET("/students", getStudents)

  // Start server
  if err := e.Start(":8081"); err != nil {
    slog.Error("failed to start server", "error", err)
  }
}

// Handler
func getStudents(c *echo.Context) error {
  return c.String(http.StatusOK, "List of all students")
}