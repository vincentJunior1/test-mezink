package controllers

import (
	v1Usecases "skeleton-svc/usecases/v1"

	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

//go:generate mockery --name Controller --output ../mocks

// type key string

// V1Controller ...
type (
	v1Controller struct {
		Usecase v1Usecases.Usecase
		Logs    *logs.Logger
	}
	V1Controller interface {
		GetRecords(ctx *gin.Context)
	}
)

// InitializeV1Controller ..
func InitializeV1Controller(usecases v1Usecases.Usecase, l *logs.Logger) V1Controller {
	return &v1Controller{
		Usecase: usecases,
		Logs:    l,
	}
}
