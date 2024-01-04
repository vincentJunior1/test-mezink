package postgre

import (
	"context"
	"embed"
	"fmt"
	"skeleton-svc/databases/postgre/models"
	"skeleton-svc/helpers"
	uModels "skeleton-svc/usecases/v1/models"
	"time"

	logs "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:generate mockery --name PostgreDatabase --output ../../mocks

//go:embed migrations/*.sql
var embedMigrations embed.FS

// PostgreDatabase ...
type (
	postgreDatabase struct {
		Db   *gorm.DB
		Logs *logs.Logger
	}

	PostgreDatabase interface {
		GetRecords(ctx context.Context, params uModels.GetRecords) ([]models.Records, error)
	}
)

// InitializePostgreDatabase ..
func InitializePostgreDatabase(log *logs.Logger) PostgreDatabase {
	return &postgreDatabase{
		Db:   ConnectPostgre(log),
		Logs: log,
	}
}

// logMode ...
var logMode = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

// ConnectPostgre ...
func ConnectPostgre(log *logs.Logger) *gorm.DB {
	username := helpers.GetEnv("USER_POSTGRE")
	password := helpers.GetEnv("PASS_POSTGRE")
	host := helpers.GetEnv("HOST_POSTGRE")
	port := helpers.GetEnv("PORT_POSTGRE")
	dbName := helpers.GetEnv("DB_POSTGRE")
	ssl := helpers.GetEnv("SSL_POSTGRE")
	debug := helpers.GetEnv("DEBUG_POSTGRE")
	mode := helpers.GetEnv("LOG_MODE_POSTGRE")

	slowLogger := logger.New(
		helpers.NewMyWriter(log),
		logger.Config{
			SlowThreshold:             time.Millisecond,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      true,
			LogLevel:                  logMode[mode],
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta", host, username, password, dbName, port, ssl)
	// gormLogger := helpers.GormLoggerNew()
	// gormLogger.LogMode(logMode[mode])
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      slowLogger,
		QueryFields: true,
	})
	if err != nil {
		log.WithFields(logs.Fields{"Message": err}).Error(helpers.GetCaller())
		panic("Error open postgres connection")
	}

	log.Info("Postgres connected successfully")

	//  Migration ===============
	// goose.SetBaseFS(embedMigrations)

	// if err := goose.SetDialect("postgres"); err != nil {
	// 	logs.WithFields(logs.Fields{"Message": err, "Goose": "postgres"}).Error(helpers.GetCaller())
	// 	return nil
	// }

	// sqlDB, _ := db.DB()

	// if err := goose.Up(sqlDB, "migrations"); err != nil {
	// 	logs.WithFields(logs.Fields{"Message": err, "Connection": dsn}).Error(helpers.GetCaller())
	// return nil
	// }
	// End Migration =============

	if debug == "true" {
		return db.Debug()
	}

	return db
}
