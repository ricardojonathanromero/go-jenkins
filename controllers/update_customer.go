package controllers

import (
	"github.com/bostaurus/jsend"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/ricardojonathanromero/go-jenkins/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// HandleUpdateCustomer update the customer info
func (c *Controller) HandleUpdateCustomer(ctx echo.Context) error {
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

	// set up validator function
	ctx.Echo().Validator = &models.Validator{Validator: validator.New()}

	// define struct
	payload := new(models.Customer)

	// validating request has a body
	if err := ctx.Bind(payload); err != nil {
		log.Errorf("Error (%v)", err)
		return ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.Error{
			Code:    "bad_request",
			Message: err.Error(),
			Source:  "api.validations.body",
		}))
	}

	// validate body request has required fields
	if err := ctx.Validate(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.Error{
			Code:    "invalid_param_values",
			Message: err.Error(),
			Source:  "api.validations.payload",
		}))
	}

	err = editCustomer(id, payload)
	if err != nil {
		return ctx.JSON(http.StatusConflict, jsend.Fail(&models.Error{
			Code:    "customer_not_updated",
			Message: err.Error(),
			Source:  "api.customers",
		}))
	}

	return ctx.NoContent(http.StatusAccepted)
}
