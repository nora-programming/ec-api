package infrastructure

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nora-programming/ec-api/interfaces/controllers"
	"github.com/nora-programming/ec-api/middlewares"
	"gorm.io/gorm"
)

type Routing struct {
	Port string
	db   *gorm.DB
}

func NewRouting(db *gorm.DB) *Routing {
	r := &Routing{
		Port: "8080",
		db:   db,
	}
	return r
}

func (r *Routing) Run() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	api := e.Group("")
	api.Use(middlewares.SetUserID)

	userController := controllers.NewUserController(r.db)

	api.GET("/me", userController.Me)
	api.DELETE("/signout", userController.Signout)
	api.PUT("/users/:id", userController.Update)
	e.POST("/signin", userController.Signin)
	e.POST("/signup", userController.Signup)

	e.Logger.Fatal(e.Start(":8080"))
}
