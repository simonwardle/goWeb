package handlers

import (
	"net/http"

	"github.com/simonwardle/goWeb/pkg/config"
	"github.com/simonwardle/goWeb/pkg/models"
	"github.com/simonwardle/goWeb/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (rec *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.tmpl.html", &models.TemplateData{})
}

// About is the about page handler
func (rec *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again!"
	render.RenderTemplate(w, "about.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
