package controllers

import (
	"github.com/bostaurus/jsend"
	"github.com/labstack/echo"
	"net/http"
)

// HandleReadCustomers returns all customers
func (c *Controller) HandleReadCustomers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, jsend.Success(customers))
}
