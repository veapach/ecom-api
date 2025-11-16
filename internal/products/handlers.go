package products

import (
	"log"
	"net/http"

	"github.com/veapach/ecom-api/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Printf("failed to list products: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}

	products := struct {
		Product []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)
}
