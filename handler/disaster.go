package handler

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"context"
	"strconv"
	"github.com/pamungkaski/camar/datamodel"
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

	data, err := h.camar.GetEarthquakeList(context.Background(), limit, page)
	if err != nil {
		fmt.Println(err)
		response.Data = err
		response.Status = http.StatusServiceUnavailable
		ctx.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response.Data = data
	response.Status = http.StatusOK
	ctx.JSON(http.StatusOK, response)
}