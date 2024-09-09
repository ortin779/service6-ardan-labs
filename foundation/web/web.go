package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*http.ServeMux
	shutdown chan os.Signal
	mws      []MidHandler
}

func NewApp(shutdown chan os.Signal, mws ...MidHandler) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
		mws:      mws,
	}
}

func (a *App) HandleFunc(method string, path string, handler Handler, mws ...MidHandler) {
	handler = wrapMiddleware(mws, handler)
	handler = wrapMiddleware(a.mws, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(r.Context(), w, r); err != nil {
			fmt.Println(err)
			return
		}
	}

	a.ServeMux.HandleFunc(fmt.Sprintf("%s %s", method, path), h)
}
