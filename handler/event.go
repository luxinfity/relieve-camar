package handler

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/pamungkaski/camar/datamodel"
	"net/http"
	"context"
)

func (h *Handler) NewEvent(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Create New Event")
	var response datamodel.Response
	var event datamodel.Event

	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON(&event); err != nil {
		response.Data = err
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	event, err := h.camar.NewEvent(context.Background(), event)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response.Data = event
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetEvent(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Get Event")
	ctx.Header("Content-Type", "application/json")
	eventID := ctx.Param("id")
	var response datamodel.Response

	event, err := h.camar.GetEvent(context.Background(), eventID)
	if err != nil {
		if event.ID == "" {
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

	response.Data = event
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateEvent(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Update Device")
	var event datamodel.Event
	var response datamodel.Response

	ctx.Header("Content-Type", "application/json")
	if err := ctx.BindJSON(&event); err != nil {
		response.Data = err
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	event, err := h.camar.UpdateEvent(context.Background(), event)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}
	response.Data = event
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteEvent(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Delete Device")
	var response datamodel.Response
	ctx.Header("Content-Type", "application/json")
	eventID := ctx.Param("id")

	err := h.camar.DeleteEvent(context.Background(), eventID)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetAllEvent(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Get All Device")
	ctx.Header("Content-Type", "application/json")
	var response datamodel.Response

	event, err := h.camar.GetAllEvent(context.Background())
	if err != nil {
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response.Data = event
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}