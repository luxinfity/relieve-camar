package handler

import (
	"github.com/pamungkaski/camar"

	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler is a struct that handles http endpoint.
type Handler struct {
	camar camar.DisasterReporter
}

func NewRouter(camar camar.DisasterReporter) *gin.Engine {
	handler := Handler{
		camar: camar,
	}

	router := gin.Default()
	router.GET("/device", handler.GetAllDevice)
	router.POST("/device", handler.RegisterDevice)
	router.GET("/device/:id", handler.GetDevice)
	router.PUT("/device/:id", handler.UpdateDevice)

	router.GET("/earthquakeList", handler.GetEarthquakeList)
	router.GET("/earthquake/:id", handler.GetEarthquakeList)

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "CAMAR RUNNING")
	})

	return router
}
