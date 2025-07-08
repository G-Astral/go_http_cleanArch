package middlewares

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ctxKey string

const requestIDKey ctxKey = "requestID"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()

		ctx := context.WithValue(c.Request.Context(), requestIDKey, reqID)
		c.Request = c.Request.WithContext(ctx)

		c.Writer.Header().Set("X-Reques-ID", reqID)

		c.Next()
	}
}

func LoggerMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var level string
		var message string

		status := c.Writer.Status()
		id := c.Request.Context().Value(requestIDKey)
		method := c.Request.Method
		path := c.Request.URL.Path
		msg, exists := c.Get("logMessage")
		if exists {
			switch v := msg.(type) {
			case string:
				message = v
			case error:
				message = v.Error()
			default:
				message = fmt.Sprintf("%v", v)
			}
		}

		switch status {
		case 200:
			level = "INFO"
		case 400, 404:
			level = "WARN"
		case 500:
			level = "ERROR"
		default:
			level = "INFO"
		}

		log.Printf("|%-3d| [%s]\t |%-36s| %-6s | %-10s | %s", status, level, id, method, path, message)
	}
}
