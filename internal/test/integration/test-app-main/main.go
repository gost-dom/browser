package main

import (
	"fmt"
	"net/http"

	app "github.com/gost-dom/browser/internal/test/integration/test-app"
)

func main() {
	server := app.CreateServer()
	fmt.Println("Starting server on port 4000")
	http.ListenAndServe(":4000", server)
}
