package api

import (
	"github.com/go-chi/chi"
	"github.com/ronyelkahfi/todos/internal/service"
)

// API contains all functions for api endpoints.
type API struct {
	service service.Service
	key     string
	env     string
}

// New to create new api endpoints.
func New(service service.Service, key, env string) *API {
	return &API{
		service: service,
		key:     key,
		env:     env,
	}
}
// Register to register api routes.
func (api *API) Register(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r.Post("/todo/create", api.HandleCreateTodo)
		r.Get("/todo/",api.HandleGetTodo)
		r.Get("/todo/complete/{id}",api.HandleCompleteTodo)
		
	})
}