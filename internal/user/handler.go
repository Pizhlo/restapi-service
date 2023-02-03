package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest_api_service/internal/apperror"
	"rest_api_service/internal/handlers"
	"rest_api_service/pkg/logging"
)

const (
	usersUrl = "/users"
	userUrl  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersUrl, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userUrl, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPost, usersUrl, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userUrl, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userUrl, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userUrl, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	return apperror.ErrNotFound
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	return apperror.NewAppError(nil, "test", "test123", "test213")
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("this is API error")
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is update user"))

	return nil
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update user"))

	return nil
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))

	return nil
}
