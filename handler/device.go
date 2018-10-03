package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/pamungkaski/camar"
)

// RegisterDevice is used to control the flow of POST /device endpoint
func (h *Handler) RegisterDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var device camar.Device

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer r.Body.Close()

	device, err := h.camar.NewDevice(context.Background(), device)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(device)
}

// Healthz is used to control the flow of GET /device endpoint
func (h *Handler) GetDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	deviceID := ps.ByName("id")

	device, err := h.camar.GetDevice(context.Background(), deviceID)
	if err != nil {
		if device.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(device)

}

// Metric is used to control the flow of GET /metrics endpoint
func (h *Handler) UpdateDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var device camar.Device

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer r.Body.Close()
	fmt.Println(device)
	device, err := h.camar.UpdateDevice(context.Background(), device)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(device)
}
