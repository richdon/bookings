package handlers

import (
	"net/http"

	"github.com/richdon/bookings/pkg/config"
	"github.com/richdon/bookings/pkg/models"
	"github.com/richdon/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository declares the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the repository type
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers sets the repository for the handler
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{}
	remote_ip := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remote_ip)
	render.RenderTemplate(w, "home.page.html", &td)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	sm := map[string]string{"test": "yo, yo, yo"}
	sm["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	td := models.TemplateData{
		StringMap: sm,
	}
	render.RenderTemplate(w, "about.page.html", &td)
}
