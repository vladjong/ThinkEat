package handlershtpp

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vladjong/ThinkEat/internal/apperror"
	"github.com/vladjong/ThinkEat/internal/handlers"
)

const (
	itemsUrl = "/api/item"
	itemUrl  = "/api/item/:uuid"
)

type handler struct {
	//логер
	//сервис
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, itemsUrl, apperror.MiddleWare(h.GetItems))
	router.HandlerFunc(http.MethodPost, itemsUrl, apperror.MiddleWare(h.CreateItems))
	router.HandlerFunc(http.MethodGet, itemUrl, apperror.MiddleWare(h.GetItemByUUID))
	router.HandlerFunc(http.MethodPut, itemUrl, apperror.MiddleWare(h.UpdateItem))
	router.HandlerFunc(http.MethodPatch, itemUrl, apperror.MiddleWare(h.PartiallyUpdateItem))
	router.HandlerFunc(http.MethodDelete, itemUrl, apperror.MiddleWare(h.DeleteUser))
}

func (h *handler) GetItems(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this items"))
	w.WriteHeader(200)
	return apperror.ErrNotFound
}

func (h *handler) GetItemByUUID(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this item"))
	w.WriteHeader(200)
	return nil
	// return apperror.ErrNotFound
}

func (h *handler) CreateItems(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this create item"))
	w.WriteHeader(204)
	return nil
}

func (h *handler) UpdateItem(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this update item"))
	w.WriteHeader(204)
	return nil
}

func (h *handler) PartiallyUpdateItem(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this partional item"))
	w.WriteHeader(204)
	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("this delete item"))
	w.WriteHeader(204)
	return nil
}
