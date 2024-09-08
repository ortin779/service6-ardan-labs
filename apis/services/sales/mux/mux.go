package mux

import (
	"net/http"
	"os"

	"github.com/ortin779/service/apis/services/sales/route/sys/checkapi"
	"github.com/ortin779/service/foundation/web"
)

func WebAPI(shutdown chan os.Signal) http.Handler {
	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
