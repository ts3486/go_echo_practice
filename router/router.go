package router

import (
	"myapp/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
    e := echo.New()


	// Middleware
 e.Use(middleware.Logger()) // Logger 
 e.Use(middleware.Recover()) // Recover

// CORS
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"*"},
  AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH,echo.POST, echo.DELETE},
 }))

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
	e.GET("/user", user.Create)
    e.GET("/users/:id", user.GetUser)
    e.PUT("/users/:id", user.UpdateUser)
    e.DELETE("/users/:id", user.DeleteUser)
    e.GET("/users", user.GetAllUsers)
    
    return e
}