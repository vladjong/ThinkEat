package item

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vladjong/ThinkEat/internal/handlers"
)

const (
	itemsUrl = "/api/item"
	itemUrl  = "/api/item:uuid"
)

type handler struct {
	//логер
	//сервис
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(itemsUrl, h.GetItems)
	router.POST(itemsUrl, h.CreateItems)
	router.GET(itemUrl, h.GetItemByUUID)
	router.PUT(itemUrl, h.UpdateItem)
	router.PATCH(itemUrl, h.PartiallyUpdateItem)
	router.DELETE(itemUrl, h.DeleteUser)
}

func (h *handler) GetItems(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this items"))
	w.WriteHeader(200)
}

func (h *handler) GetItemByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this item"))
	w.WriteHeader(200)
}

func (h *handler) CreateItems(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this create item"))
	w.WriteHeader(204)
}

func (h *handler) UpdateItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this update item"))
	w.WriteHeader(204)
}

func (h *handler) PartiallyUpdateItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this partional item"))
	w.WriteHeader(204)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this delete item"))
	w.WriteHeader(204)
}
