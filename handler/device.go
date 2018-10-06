package handler

import (
	"fmt"
	"net/http"

	"github.com/pamungkaski/camar"
	"github.com/gin-gonic/gin"
	"context"
)

// RegisterDevice is used to control the flow of POST /device endpoint
func (h *Handler) RegisterDevice(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Create New Device")
	var device camar.Device

	ctx.Header("Content-Type", "application/json")
	
	if err := ctx.BindJSON(&device); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	device, err := h.camar.NewDevice(context.Background(), device)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, device)
}

// Healthz is used to control the flow of GET /device endpoint
func (h *Handler) GetDevice(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Get Device")
	ctx.Header("Content-Type", "application/json")
	deviceID := ctx.Param("id")

	device, err := h.camar.GetDevice(context.Background(), deviceID)
	if err != nil {
		if device.ID == "" {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}
	ctx.JSON(http.StatusOK, device)
}

// Metric is used to control the flow of GET /metrics endpoint
func (h *Handler) UpdateDevice(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Update Device")
	var device camar.Device

	ctx.Header("Content-Type", "application/json")
	if err := ctx.BindJSON(&device); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(device)
	device, err := h.camar.UpdateDevice(context.Background(), device)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, device)
}
