package main

import (
	"net/http"

	app "github.com/gost-dom/browser/internal/test/integration/test-app"
)

func main() {
	server := app.CreateServer()
	http.ListenAndServe(":4000", server)
}
