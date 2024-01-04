package controllers

import (
	v1Controller "skeleton-svc/controllers/v1"

	logs "github.com/sirupsen/logrus"
)

//go:generate mockery --name Controller --output ../mocks

// Controller ...
type (
	controller struct {
		V1Controller v1Controller.V1Controller
		Logs         *logs.Logger
	}
	Controller interface {
		V1() v1Controller.V1Controller
	}
)

// InitializeController ..
func InitializeController(
	v1 v1Controller.V1Controller,
	l *logs.Logger,
) Controller {
	return &controller{
		V1Controller: v1,
		Logs:         l,
	}
}

// V1 controller ...
func (c *controller) V1() v1Controller.V1Controller {
	return c.V1Controller
}
