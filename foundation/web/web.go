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
}

func NewApp(shutdown chan os.Signal) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
	}
}

func (a *App) HandleFunc(method string, path string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(r.Context(), w, r)
	}

	a.ServeMux.HandleFunc(fmt.Sprintf("%s %s", method, path), h)
}
