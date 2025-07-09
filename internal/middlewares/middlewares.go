package middlewares

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ctxKey string

const requestIDKey ctxKey = "requestID"

// Injects X-Request-ID into context and headers
func RequestIDContexMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()

		ctx := context.WithValue(c.Request.Context(), requestIDKey, reqID)
		c.Request = c.Request.WithContext(ctx)

		c.Writer.Header().Set("X-Reques-ID", reqID)

		c.Next()
	}
}

func LoggerMiddlware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

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

		fields := []zap.Field{
			zap.Int("status", status),
			zap.Any("requestID", id),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("message", message),
		}

		switch {
		case status >= 500:
			logger.Error("server error", fields...)
		case status >= 400:
			logger.Warn("client error", fields...)
		default:
			logger.Info("successfullR", fields...)
		}
	}
}
