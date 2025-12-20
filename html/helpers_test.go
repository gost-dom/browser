package html_test

import (
	"net/http"
)

// A simple http Handler that just calls ParseForm, making the form data
// available on the http request object. When combined with
// [HTTPRequestRecorder], test code can easily inspect the submitted form
// values.
var ParseFormHandler http.Handler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) { r.ParseForm() },
)
