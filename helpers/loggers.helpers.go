package helpers

import (
	"fmt"
	"os"
	"runtime"
	"skeleton-svc/constants"
	"skeleton-svc/helpers/models"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"

	// elastic "github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"

	// elogrus "gopkg.in/sohlich/elogrus.v7"
	"github.com/FWangZil/gorm2logrus"
)

// InitializeNewLogs ...
func InitializeNewLogs() *log.Logger {
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{
		PrettyPrint: false,
	})
	l.SetReportCaller(true)

	l.Out = os.Stdout
	file, err := os.OpenFile(GetEnv("LOG_NAME"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		l.Out = file
	} else {
		l.Warn("Failed to log to file, using default stderr")
	}

	return l
}

// TraceIDHook ...
type TraceIDHook struct {
	TraceID string
}

// NewTraceIDHook ...
func NewTraceIDHook(traceID string) log.Hook {
	hook := TraceIDHook{
		TraceID: traceID,
	}
	return &hook
}

// Fire ...
func (hook *TraceIDHook) Fire(entry *log.Entry) error {
	entry.Data["TRACE_ID"] = hook.TraceID
	entry.Data["SERVICE_NAME"] = constants.SERVICE_NAME
	return nil
}

// Levels ...
func (hook *TraceIDHook) Levels() []log.Level {
	return log.AllLevels
}

// GetCaller ..
func GetCaller() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// GettingResponseLog ...
func GettingResponseLog(g *gin.Context, req, res interface{}, duration time.Duration) log.Fields {
	logsField := models.LogModels{
		METHOD:   g.Request.Method,
		PATH:     g.Request.URL.Path,
		HEADER:   g.Request.Header,
		CLIENTIP: g.Request.RemoteAddr,
		REQUEST:  req,
		RESPONSE: res,
		DURATION: fmt.Sprintf("%v s", duration.Seconds()),
	}
	return structs.Map(logsField)
}

// GettingDetaultLog ...
func GettingDetaultLog(message interface{}) log.Fields {
	return log.Fields{
		"MESSAGE": message,
	}
}

// MyWriter ...
type MyWriter struct {
	log *log.Logger
}

// Printf ...
func (m *MyWriter) Printf(format string, v ...interface{}) {
	m.log.WithFields(log.Fields{
		"REF_FILE":     v[0],
		"DESCRIPTION":  v[1],
		"DURATION":     fmt.Sprintf("%v ms", v[2]),
		"ROW_AFFECTED": v[3],
		"QUERY":        fmt.Sprintf("%v", v[4]),
	}).Info("Query logs")
}

// NewMyWriter ...
func NewMyWriter(log *log.Logger) *MyWriter {
	return &MyWriter{log: log}
}

// GormLogger ...
var GormLogger *gorm2logrus.GormLogrus

// Logger ...
var Logger *log.Logger

// InitLogrus ...
func InitLogrus() (*log.Logger, *gorm2logrus.GormLogrus) {
	logger := gorm2logrus.NewGormLogger()

	logger.SetReportCaller(true)
	logger.SetFormatter(&log.JSONFormatter{
		PrettyPrint: false,
	})

	GormLogger = logger
	Logger = &GormLogger.Logger
	return Logger, GormLogger
}
