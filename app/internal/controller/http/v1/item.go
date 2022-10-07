package v1

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vladjong/ThinkEat/internal/entities"
	"github.com/vladjong/ThinkEat/internal/interfaces"
	"github.com/vladjong/ThinkEat/pkg/logging"
)

const (
	itemsUrl = "/api/items"
	itemUrl  = "/api/item/:uuid"
)

type itemHandler struct {
	itemUseCase interfaces.Item
	logger      *logging.Logger
}

func NewItemHandler(itemUseCase interfaces.Item, logger *logging.Logger) *itemHandler {
	return &itemHandler{
		itemUseCase: itemUseCase,
		logger:      logger,
	}
}

func (h *itemHandler) Register(router *httprouter.Router) {
	// router.HandlerFunc(http.MethodGet, itemsUrl, apperror.MiddleWare(h.GetItems))
	router.GET(itemsUrl, h.GetItems)
	router.POST(itemsUrl, h.CreateItems)
	router.PUT(itemsUrl, h.UpdateItem)
	router.GET(itemUrl, h.GetItemByUUID)
	// router.PATCH(itemUrl, h.PartiallyUpdateItem)
	router.DELETE(itemUrl, h.DeleteUser)
}

func (h *itemHandler) GetItems(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	items, err := h.itemUseCase.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(400)
		panic(err)
	}
	itemsBytes, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(itemsBytes)
}

func (h *itemHandler) GetItemByUUID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("uuid")
	item, err := h.itemUseCase.GetID(r.Context(), id)
	if err != nil {
		panic(err)
	}
	itemBytes, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(itemBytes)
}

func (h *itemHandler) CreateItems(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var item entities.Item
	if err := decoder.Decode(&item); err != nil {
		panic(err)
	}
	h.itemUseCase.Create(r.Context(), &item)
	w.WriteHeader(http.StatusCreated)
}

func (h *itemHandler) UpdateItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var item entities.Item
	if err := decoder.Decode(&item); err != nil {
		panic(err)
	}
	if err := h.itemUseCase.Update(r.Context(), &item); err != nil {
		panic(err)
	}
	w.Write([]byte("this update item"))
	w.WriteHeader(204)
}

// func (h *itemHandler) PartiallyUpdateItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	w.Write([]byte("this partional item"))
// 	w.WriteHeader(204)
// }

func (h *itemHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("uuid")
	if err := h.itemUseCase.Delete(r.Context(), id); err != nil {
		panic(err)
	}
	w.Write([]byte("this delete item"))
	w.WriteHeader(204)
}
