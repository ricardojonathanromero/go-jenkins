package controllers

import (
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/ricardojonathanromero/go-jenkins/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var customers = make([]*models.Customer, 0)

type Controller struct{}

func findCustomer(ID primitive.ObjectID) (int, *models.Customer, error) {
	for pos, customer := range customers {
		if customer.ID == ID {
			return pos, customer, nil
		}
	}

	return 0, nil, errors.New("customer not found")
}

func editCustomer(ID primitive.ObjectID, customer *models.Customer) error {
	pos, _, err := findCustomer(ID)
	if err != nil {
		log.Errorf("Error (%v)", err)
		return errors.New("customer not updated")
	}

	customers[pos] = customer
	return nil
}
