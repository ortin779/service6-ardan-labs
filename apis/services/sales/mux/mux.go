package mux

import (
	"net/http"

	"github.com/ortin779/service/apis/services/sales/route/sys/checkapi"
)

func WebAPI() http.Handler {
	mux := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
