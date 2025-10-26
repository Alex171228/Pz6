package httpapi

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func BuildRouter(d *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	h := NewHandlers(d)

	r.Get("/health", h.Health)

	// Users
	r.Post("/users", h.CreateUser)

	// Notes
	r.Post("/notes", h.CreateNote)      // create note with tags
	r.Get("/notes/{id}", h.GetNoteByID) // get note with author and tags

	return r
}
