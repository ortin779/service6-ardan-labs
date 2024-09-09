package mux

import (
	"net/http"
	"os"

	"github.com/ortin779/service/apis/services/api/mid"
	"github.com/ortin779/service/apis/services/sales/route/sys/checkapi"
	"github.com/ortin779/service/foundation/logger"
	"github.com/ortin779/service/foundation/web"
)

func WebAPI(log *logger.Logger, shutdown chan os.Signal) http.Handler {
	mux := web.NewApp(shutdown, mid.Logger(log))

	checkapi.Routes(mux)

	return mux
}
