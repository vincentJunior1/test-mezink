package usecases

import (
	// "context"

	"context"
	"skeleton-svc/databases"
	hModels "skeleton-svc/helpers/models"
	"skeleton-svc/usecases/v1/models"

	logs "github.com/sirupsen/logrus"
	// pModels "skeleton-svc/databases/postgre/models"
	// hModels "skeleton-svc/helpers/models"
)

//go:generate mockery --name Usecase --output ../mocks

// Usecase ...
type (
	usecase struct {
		DB   databases.Database
		Logs *logs.Logger
	}

	Usecase interface {
		GetRecords(ctx context.Context, params models.GetRecords) hModels.Meta
	}
)

// InitializeV1Usecase ...
func InitializeV1Usecase(
	db databases.Database,
	l *logs.Logger,
) Usecase {
	return &usecase{
		DB:   db,
		Logs: l,
	}
}
