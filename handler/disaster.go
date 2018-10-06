package handler

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"context"
	"strconv"
)

func (h *Handler) GetEarthquakeList(ctx *gin.Context) {
	fmt.Println("Endpoint Hit: Get earthquake List")
	ctx.Header("Content-Type", "application/json")

	limS := ctx.Query("limit")
	pageS := ctx.Query("page")

	limit, err := strconv.Atoi(limS)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	page, err := strconv.Atoi(pageS)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	data, err := h.camar.GetEarthquakeList(context.Background(), limit, page)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}