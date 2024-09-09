package mid

import (
	"context"
	"net/http"

	"github.com/ortin779/service/foundation/logger"
	"github.com/ortin779/service/foundation/web"
)

func Logger(log *logger.Logger) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// Logging Here
			log.Info(ctx, "request started", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			err := handler(ctx, w, r)

			// Logging Here
			log.Info(ctx, "request ended", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
			return err
		}
		return h
	}
	return m
}
