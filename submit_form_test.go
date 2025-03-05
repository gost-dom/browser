package browser

import (
	"net/http"
	"testing"
)

func TestSubmitForm(t *testing.T) {
	r := http.NewServeMux()
	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
		<form action="/submit" method="post">
			<input type="hidden" name="hidden" value="hidden-value">
			<button name="name" value="John">John</button>
			<input type="submit" name="name" value="Jane">Jane</button>
		</form>
		`))
	})
	r.HandleFunc("POST /submit", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			t.Fatal(err)
		}
		if r.FormValue("hidden") != "hidden-value" {
			t.Errorf("Invalid hidden value: %s", r.FormValue("hidden"))
			http.Error(w, "Invalid hidden value", http.StatusBadRequest)
		}
		switch r.FormValue("name") {
		case "John":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`<h1>Form submitted</h1>`))
		case "Jane":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`<h1>Form submitted</h1>`))
		default:
			t.Errorf("Invalid name: %s", r.FormValue("name"))
			http.Error(w, "Invalid name", http.StatusBadRequest)
		}
	})

	b := NewFromHandler(r)
	win, err := b.Open("http://example.com")
	if err != nil {
		t.Fatal(err)
	}
	doc := win.Document()
	t.Run("Works for John button", func(t *testing.T) {
		btn, err := doc.QuerySelector("button[value=John]")
		if err != nil {
			t.Fatal(err)
		}
		if !btn.Click() {
			t.Fatal(err)
		}
	})

	t.Run("Works for Jane input button", func(t *testing.T) {
		btn, err := doc.QuerySelector("input[value=Jane]")
		if err != nil {
			t.Fatal(err)
		}
		if !btn.Click() {
			t.Fatal("Failed to click button")
		}
	})
}
