package common

import (
	"context"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type Context struct {
	CorrelationID *string
	Data          *map[string]interface{}
	Logger        *zap.SugaredLogger
	Ctx           context.Context
}

// CreateLoggableContextFromRequest creates a new context with logger and co-relation ID for tracking
func CreateLoggableContextFromRequest(request *http.Request, parentLogger *zap.SugaredLogger) *Context {

	correlationID := GetCorrelationID(request)
	return &Context{
		CorrelationID: &correlationID,
		Logger:        parentLogger.With("co-relation-id", correlationID),
		Ctx:           context.Background(),
	}
}

// GetCorrelationID returns co-relation ID from request, creates one if not present
func GetCorrelationID(request *http.Request) string {

	correlationID := request.Header.Get("CORRELATION-ID")
	if correlationID == "" {
		correlationID = getUniqID()
	}
	return correlationID
}

func getUniqID() string {
	randID := strings.Replace(uuid.Must(uuid.NewV4(), nil).String(), "-", "", -1)
	timeID := strings.Replace(uuid.Must(uuid.NewV1(), nil).String(), "-", "", -1)
	return timeID[:len(timeID)/2] + randID[len(randID)/2:]
}
