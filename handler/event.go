package handler

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/pamungkaski/camar/datamodel"
	"net/http"
	"context"
	"strconv"
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
	fmt.Println("Endpoint Hit: Get All Event")
	ctx.Header("Content-Type", "application/json")
	response := struct {
		Status int `json:"status"`
		MaxResults int `json:"max_results"`
		Data interface{} `json:"data"`
	}{}
	var err error
	limit := 20
	page := 1

	limS := ctx.Query("limit")
	pageS := ctx.Query("page")
	eventType := ctx.Query("event_type")

	if limS != "" {
		limit, err = strconv.Atoi(limS)
		if err != nil {
			fmt.Println(err)
			response.Data = err
			response.Status = http.StatusBadRequest
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
	}

	if pageS != "" {
		page, err = strconv.Atoi(pageS)
		if err != nil {
			fmt.Println(err)
			response.Data = err
			response.Status = http.StatusBadRequest
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
	}

	data, count, err := h.camar.GetAllEvent(context.Background(), limit, page, eventType)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response = struct {
		Status int `json:"status"`
		MaxResults int `json:"max_results"`
		Data interface{} `json:"data"`
	}{
		Status: http.StatusOK,
		MaxResults: count,
		Data: data,
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CheckBMKG(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Check BMKG")
	response := struct {
		Status int `json:"status"`
		MaxResults int `json:"max_results"`
		Data interface{} `json:"data"`
	}{}
	ctx.Header("Content-Type", "application/json")

	err := h.camar.ListenTheEarth()
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	data, count, err := h.camar.GetAllEvent(context.Background(), 20, 1, "earthquake")
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response = struct {
		Status int `json:"status"`
		MaxResults int `json:"max_results"`
		Data interface{} `json:"data"`
	}{
		Status: http.StatusOK,
		MaxResults: count,
		Data: data,
	}

	ctx.JSON(http.StatusOK, response)
}