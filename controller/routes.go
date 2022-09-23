package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (s *Server) initializeRoutes() {

	// Home Route
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	api := e.Group("/api/v1", serverHeader)
	api.GET("/", s.Home)
	api.POST("/login", s.Login)
	api.POST("/users", s.CreateUser)
	api.GET("/users", s.GetUsers)          // Returns all resources of this product
	api.GET("/users/:id", s.GetUser)       // Returns the resource of this product with that ID
	api.PUT("/users/:id", s.UpdateUser)    // Updates the resource of this product with that ID
	api.DELETE("/users/:id", s.DeleteUser) // Deletes the resource of this product with that ID

	apiV2 := e.Group("/api/v2", serverHeaderVersion2)
	apiV2.GET("/", s.Home)
	apiV2.POST("/login", s.Login)
	apiV2.GET("/users", s.GetUsers)          // Returns all resources of this product
	apiV2.POST("/users", s.CreateUser)       // Creates a resource of this product and stores the data you posted, then returns the ID
	apiV2.GET("/users/:id", s.GetUser)       // Returns the resource of this product with that ID
	apiV2.PUT("/users/:id", s.UpdateUser)    // Updates the resource of this product with that ID
	apiV2.DELETE("/users/:id", s.DeleteUser) // Deletes the resource of this product with that ID
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
