package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
)

type HttpLoggerMiddleware struct {
	logger *zap.Logger
}

func NewHttpLoggerMiddleware(logger *zap.Logger) *HttpLoggerMiddleware {
	return &HttpLoggerMiddleware{logger: logger}
}

func (h *HttpLoggerMiddleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json, err := httputil.DumpRequest(r, true)
		if err != nil {
			h.logger.Error("fail to parse http request", zap.Error(err))
			return
		}
		h.logger.Debug("", zap.String("request", string(json)))
		next.ServeHTTP(w, r)
		return
	})
}
