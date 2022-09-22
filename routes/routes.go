package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/wewebplus/wewebapi/core/handler"
	"github.com/wewebplus/wewebapi/core/models"
)

type Routes struct {
	Router *mux.Router
}

func (routes *Routes) Initialize() {

	// Home Route
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	api := e.Group("/api/v1", serverHeader)
	api.POST("/users", handler.CreateUser)
	api.GET("/users", handler.GetUsers)          // Returns all resources of this product
	api.GET("/users/:id", handler.GetUser)       // Returns the resource of this product with that ID
	api.PUT("/users/:id", handler.UpdateUser)    // Updates the resource of this product with that ID
	api.DELETE("/users/:id", handler.DeleteUser) // Deletes the resource of this product with that ID

	apiV2 := e.Group("/api/v2", serverHeaderVersion2)

	apiV2.GET("/users", handler.GetUsers)          // Returns all resources of this product
	apiV2.POST("/users", handler.CreateUser)       // Creates a resource of this product and stores the data you posted, then returns the ID
	apiV2.GET("/users/:id", handler.GetUser)       // Returns the resource of this product with that ID
	apiV2.PUT("/users/:id", handler.UpdateUser)    // Updates the resource of this product with that ID
	apiV2.DELETE("/users/:id", handler.DeleteUser) // Deletes the resource of this product with that ID
	err := models.Ping()
	if err != nil {
		logrus.Fatal(err)
	}

	// service start at port :9090
	err = e.Start(":9090")
	if err != nil {
		logrus.Fatal(err)
	}
}
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "Test/v1.0")
		return next(c)
	}
}

func serverHeaderVersion2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "Test/v2.0")
		return next(c)
	}
}

func (routes *Routes) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, routes.Router))
}
