package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pamungkaski/camar/datamodel"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetEarthquakeList(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Get earthquake List")
	ctx.Header("Content-Type", "application/json")
	var response datamodel.Response

	limS := ctx.Query("limit")
	pageS := ctx.Query("page")

	limit, err := strconv.Atoi(limS)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	page, err := strconv.Atoi(pageS)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	data, count, err := h.camar.GetEarthquakeList(context.Background(), limit, page)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response.Data = struct {
		Limit int `json:"limit"`
		Page int `json:"page"`
		MaxResults int `json:"max_results"`
		TimeStamp int64 `json:"time_stamp"`
		Data interface{} `json:"data"`
	}{
		Limit: limit,
		Page: page,
		MaxResults: count,
		TimeStamp:  time.Now().Unix(),
		Data: data,
	}
	response.Status = http.StatusOK
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetEarthquake(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Get earthquake")
	ctx.Header("Content-Type", "application/json")
	var response datamodel.Response

	id := ctx.Param("id")

	data, err := h.camar.GetEarthquake(context.Background(), id)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusNotFound
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response.Data = data
	response.Status = http.StatusOK
	ctx.JSON(http.StatusOK, response)
}
