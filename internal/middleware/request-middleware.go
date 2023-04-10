package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RequestLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		ctx.Next()

		var logEvent *zerolog.Event
		if ctx.Writer.Status() >= 400 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}
		logEvent.Str("client_id", ctx.ClientIP()).
			Str("method", ctx.Request.Method).
			Int("status_code", ctx.Writer.Status()).
			Int("body_size", ctx.Writer.Size()).
			Str("path", path).
			Str("latency", time.Since(start).String()).
			Msg(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}
