package handlers

import (
	"net/http"
	"strconv"
	"time"
	"wigata-ai/views"
	"wigata-ai/views/components"

	"github.com/a-h/templ"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.FS(views.StaticFiles))))

	mux.Handle("/api/", http.StripPrefix("/api", apiHandler()))
	mux.Handle("/", viewHandler())

	return mux
}

func apiHandler() http.Handler {
	mux := http.NewServeMux()

	// TODO: Add HealthCheck & Logger Middleware

	return mux
}

func viewHandler() http.Handler {
	mux := http.NewServeMux()

	yearNow := strconv.Itoa(time.Now().Year())
	indexComponents := components.Index(yearNow)

	mux.Handle("/", templ.Handler(indexComponents))

	return mux
}
