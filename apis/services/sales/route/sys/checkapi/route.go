package checkapi

import (
	"github.com/ortin779/service/foundation/web"
)

func Routes(app *web.App) {
	app.HandleFunc("GET", "/liveness", liveness)
	app.HandleFunc("GET", "/readiness", readiness)

}
