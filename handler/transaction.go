package handler

import (
	"net/http"
	"strconv"

	"github.com/agusheryanto182/go-wangkas/helper"
	"github.com/agusheryanto182/go-wangkas/transaction"
	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransaksiHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) GetAllData(c *gin.Context) {
	result, err := h.transactionService.GetAllDataTransactions()
	if err != nil {
		response := helper.APIResponse("Failed to get all data transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success get all data transactions", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetDataByWeekID(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	result, err := h.transactionService.GetByWeekID(ID)
	if err != nil {
		response := helper.APIResponse("Failed to get user detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success get data by week id", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}
