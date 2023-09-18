package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/ronyelkahfi/todos/internal/errors"
	"github.com/ronyelkahfi/todos/internal/service"
	"github.com/ronyelkahfi/todos/internal/utils"
)

func (api *API) HandleCreateTodo(w http.ResponseWriter, r *http.Request){
	var request service.CreateTodoRequest
	
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.ResponseWithJSON(w, http.StatusBadRequest, nil, errors.Wrap(r.Context(), errors.ErrInvalidRequestFormat, err))
		return
	}

	todo, code, err := api.service.Create(r.Context(), request)
	utils.ResponseWithJSON(w, code, todo, errors.Wrap(r.Context(), err))

}

func (api *API) HandleGetTodo(w http.ResponseWriter, r *http.Request){
	todos, code, err := api.service.GetAll(r.Context())
	utils.ResponseWithJSON(w, code, todos, errors.Wrap(r.Context(), err))
}
func (api *API) HandleCompleteTodo(w http.ResponseWriter, r *http.Request){
	
	// strconv.ParseInt(chi.URLParam(r, "id"),10,64)
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	todos,code, err := api.service.UpdateStatusCompleteByID(r.Context(), id)
	utils.ResponseWithJSON(w, code, todos, errors.Wrap(r.Context(), err))
}
