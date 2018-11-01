package handler

import (
	"fmt"
	"net/http"

	"github.com/pamungkaski/camar"
	"github.com/gin-gonic/gin"
	"context"
	"github.com/pamungkaski/camar/datamodel"
)

// RegisterDevice is used to control the flow of POST /device endpoint
func (h *Handler) RegisterDevice(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Create New Device")
	var response datamodel.Response
	var device camar.Device

	ctx.Header("Content-Type", "application/json")
	
	if err := ctx.BindJSON(&device); err != nil {
		response.Data = err
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	device, err := h.camar.NewDevice(context.Background(), device)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response.Data = device
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

// Healthz is used to control the flow of GET /device endpoint
func (h *Handler) GetDevice(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Get Device")
	ctx.Header("Content-Type", "application/json")
	deviceID := ctx.Param("id")
	var response datamodel.Response

	device, err := h.camar.GetDevice(context.Background(), deviceID)
	if err != nil {
		if device.ID == "" {
			response.Data = err
			response.Status = http.StatusNotFound
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response.Data = device
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}


// Healthz is used to control the flow of GET /device endpoint
func (h *Handler) GetAllDevice(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Get All Device")
	ctx.Header("Content-Type", "application/json")
	var response datamodel.Response

	device, err := h.camar.GetAllDevice(context.Background())
	if err != nil {
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response.Data = device
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}


// Metric is used to control the flow of GET /metrics endpoint
func (h *Handler) UpdateDevice(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Update Device")
	var device camar.Device
	var response datamodel.Response

	ctx.Header("Content-Type", "application/json")
	if err := ctx.BindJSON(&device); err != nil {
		response.Data = err
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(device)
	device, err := h.camar.UpdateDevice(context.Background(), device)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}
	response.Data = device
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}
