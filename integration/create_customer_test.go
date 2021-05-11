package integration

import (
	"github.com/labstack/echo"
	"github.com/ricardojonathanromero/go-jenkins/controllers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const path = "/api/v1/customers"

func TestCreateCustomerBadRequestBody(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(``))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	controller := &controllers.Controller{}
	// Assertions
	if assert.NoError(t, controller.HandleCreateCustomer(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `{"status":"fail","data":{"code":"bad_request","message":"Request body can't be empty","source":"api.validations.body"}}`, strings.ReplaceAll(rec.Body.String(), "\n", ""))
	}
}

func TestCreateCustomerBadRequestRequiredField(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(`{"name": "Ernesto"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := &controllers.Controller{}

	// Assertions
	if assert.NoError(t, controller.HandleCreateCustomer(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `{"status":"fail","data":{"code":"invalid_param_values","message":"Key: 'Customer.Lastname' Error:Field validation for 'Lastname' failed on the 'required' tag\nKey: 'Customer.Age' Error:Field validation for 'Age' failed on the 'required' tag\nKey: 'Customer.Email' Error:Field validation for 'Email' failed on the 'email' tag","source":"api.validations.payload"}}`,
			strings.ReplaceAll(rec.Body.String(), "\n", ""))
	}
}

func TestCreateCustomer(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(`{"name": "Ernesto", "lastname": "Mussolini", "age": 25, "email": "benito.mussolini@gmail.com", "gender": "male"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := &controllers.Controller{}

	// Assertions
	if assert.NoError(t, controller.HandleCreateCustomer(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `{"status":"success","data":"Customer created successfully"}`, strings.ReplaceAll(rec.Body.String(), "\n", ""))
	}
}
