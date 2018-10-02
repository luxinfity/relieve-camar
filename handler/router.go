package handler

import (
	"github.com/pamungkaski/camar"
	"github.com/julienschmidt/httprouter"
)

// Handler is a struct that handles http endpoint.
type Handler struct{
	camar camar.DisasterReporter
}

func NewRouter(camar camar.DisasterReporter) *httprouter.Router {
	handler := Handler{
		camar:camar,
	}

	router := httprouter.New()
	router.POST("/device", handler.RegisterDevice)
	router.GET("/device/:id", handler.GetDevice)
	router.PUT("/device/:id", handler.UpdateDevice)

	return router
}