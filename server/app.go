package server

import (
	"github.com/labstack/echo"
	"github.com/ricardojonathanromero/go-jenkins/controllers"
)

// NewRouter func configures the routes that
// the server should be listening.
// By default, the routes will be configured under
// the package controllers.
func NewRouter() *echo.Echo {
	// here goes all routes will be used by the server
	e := echo.New()
	c := controllers.Controller{}

	// implement middlewares here
	// ...

	// routes goes here
	OpenAPI := e.Group("/api/v1")
	OpenAPI.POST("/customers", c.HandleCreateCustomer)              // This function push a new customer
	OpenAPI.GET("/customers", c.HandleReadCustomers)                // This function returns all customers
	OpenAPI.GET("/customers/:customerID", c.HandleReadCustomer)     // This function returns a specific customer by id
	OpenAPI.PATCH("/customers/:customerID", c.HandleUpdateCustomer) // This function updates the fields of a specific customer

	return e
}
