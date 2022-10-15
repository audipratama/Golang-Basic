package usecase

import (
	"github.com/gin-gonic/gin"
	"golang-basic/model"
	"net/http"
	"strconv"
)

type OrderRequest struct {
	GoodsName string `json:"goods_name"`
}

type OrderResponse struct {
	Id        int64  `json:"id"`
	GoodsName string `json:"goods_name"`
}

func (u usecase) GetOrder(context *gin.Context) {
	id := context.Param("id")
	orderId , err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})

		context.Abort()
		return
	}

	orders, err := u.orderRepo.GetOrdersByIDs(context, []int64{orderId})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error Get to DB",
		})

		context.Abort()
		return
	}
	if len(orders) < 1{
		context.JSON(http.StatusNotFound, gin.H{
			"message": "No Data Found",
		})

		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": OrderResponse{
			Id:        orders[0].Id,
			GoodsName: orders[0].GoodsName},
		"message": "Get Order Success",
	})
}

func (u usecase) CreateOrder(context *gin.Context) {
	var request OrderRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})

		context.Abort()
		return
	}

	order := model.Orders{
		GoodsName: request.GoodsName,
	}

	result, err := u.orderRepo.Insert(context, order)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error insert to DB",
		})

		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"data": OrderResponse{
			Id:        result,
			GoodsName: request.GoodsName},
		"message": "Create Order Success",
	})
}
