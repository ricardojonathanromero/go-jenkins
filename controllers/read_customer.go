package controllers

import (
	"github.com/bostaurus/jsend"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/ricardojonathanromero/go-jenkins/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// HandleReadCustomer returns a specific customer info
func (c *Controller) HandleReadCustomer(ctx echo.Context) error {
	customerID := ctx.Param("customerID")

	id, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		log.Errorf("Error (%v)", err)

		return ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.Error{
			Code:    "invalid_request_params",
			Message: "The customerID must be a valid ID",
			Source:  "api.validations.path",
		}))
	}

	_, customer, err := findCustomer(id)
	if err != nil {
		log.Errorf("Error finding customer with id %v Error (%v)", id, err)

		return ctx.JSON(http.StatusNotFound, jsend.Fail(&models.Error{
			Code:    "not_found",
			Message: err.Error(),
			Source:  "api.customers",
		}))
	}

	return ctx.JSON(http.StatusOK, jsend.Success(customer))
}
