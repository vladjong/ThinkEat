package v1

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vladjong/ThinkEat/internal/entities"
	"github.com/vladjong/ThinkEat/internal/interfaces"
)

const (
	placesUrl = "/api/places"
	placeUrl  = "/api/place/:uuid"
)

type placeHandler struct {
	placeUseCase interfaces.Place
}

func NewPlaceHandler(place interfaces.Place) *placeHandler {
	return &placeHandler{
		placeUseCase: place,
	}
}

func (h *placeHandler) Register(router *httprouter.Router) {
	router.GET(placesUrl, h.GetPlaces)
	router.POST(placesUrl, h.CreatePlace)
	router.PUT(placesUrl, h.UpdatePlace)
	router.GET(placeUrl, h.GetPlaceByUUID)
	// router.PATCH(itemUrl, h.PartiallyUpdateItem)
	router.DELETE(placeUrl, h.DeletePlace)
}

func (h *placeHandler) GetPlaces(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	items, err := h.placeUseCase.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	placesBytes, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(placesBytes)
}

func (h *placeHandler) GetPlaceByUUID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("uuid")
	place, err := h.placeUseCase.GetID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	placeBytes, err := json.Marshal(place)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(placeBytes)
}

func (h *placeHandler) CreatePlace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var place entities.Place
	if err := decoder.Decode(&place); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	_, err := h.placeUseCase.Create(r.Context(), &place)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *placeHandler) UpdatePlace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var place entities.Place
	if err := decoder.Decode(&place); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	if err := h.placeUseCase.Update(r.Context(), &place); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *placeHandler) DeletePlace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("uuid")
	if err := h.placeUseCase.Delete(r.Context(), id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusNoContent)
}
