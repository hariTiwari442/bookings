package handlers

import (
	"net/http"

	modals "github.com/hariTiwari442/bookings/pkg/Modals"
	"github.com/hariTiwari442/bookings/pkg/config"
	"github.com/hariTiwari442/bookings/pkg/render"
	// modals "github.com/hariTiwari442/bookings/pkg/Modals"
	// "github.com/hariTiwari442/go-course1/pkg/render"
)

// TemplateData holds data sent from handlers to templates

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// NewHandlers sets the repository for the  handle

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &modals.TemplateData{})
}

// About is the about pahge handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	render.RenderTemplate(w, "about.page.html", &modals.TemplateData{
		StringMap: stringMap,
	})

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
}

// This function add two values as simple as it is
// func addValues(x, y int) int {
// 	return x + y
// }

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)

// 	if err != nil {
// 		fmt.Fprintf(w, "Can not divide by zero")
// 	}
// 	fmt.Fprintf(w, fmt.Sprintf("%f divid eby %f is %f", 100.0, 10.0, f))
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }
