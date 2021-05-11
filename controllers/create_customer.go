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

// HandleCreateCustomer append a new customer
func (c *Controller) HandleCreateCustomer(ctx echo.Context) error {
	// set up validator function
	ctx.Echo().Validator = &models.Validator{Validator: validator.New()}

	// define struct
	payload := new(models.Customer)
	payload.ID = primitive.NewObjectID()

	// validating request has a body
	if err := ctx.Bind(payload); err != nil {
		log.Errorf("Error (%v)", err)
		return ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.Error{
			Code:    "bad_request",
			Message: "Request body can't be empty",
			Source:  "api.validations.body"}))
	}

	// validate body request has required fields
	if err := ctx.Validate(payload); err != nil {
		log.Errorf("Error (%v)", err)
		return ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.Error{
			Code:    "invalid_param_values",
			Message: err.Error(),
			Source:  "api.validations.payload",
		}))
	}

	customers = append(customers, payload)
	return ctx.JSON(http.StatusCreated, jsend.Success("Customer created successfully"))
}
