package databases

import (
	"skeleton-svc/databases/postgre"

	logs "github.com/sirupsen/logrus"
)

//go:generate mockery --name Database --output ../mocks

// Database ..
type (
	database struct {
		Postgre postgre.PostgreDatabase
	}
	Database interface {
		GetPostgre() postgre.PostgreDatabase
		// Insert other database here ...
	}
)

// InitializeDatabase ..
func InitializeDatabase(
	psqlCon postgre.PostgreDatabase,
	l *logs.Logger,
) Database {
	return &database{
		Postgre: psqlCon,
	}
}

// GetPostgre for get connection postgre ...
func (d *database) GetPostgre() postgre.PostgreDatabase {
	return d.Postgre
}
