package handler

import (
	"net/http"
	"strconv"

	"github.com/agusheryanto182/go-wangkas/transaction"
	"github.com/gin-gonic/gin"
)

type transactionsHandler struct {
	transactionsService transaction.Service
}

func NewTransactionsHandler(transactionsService transaction.Service) *transactionsHandler {
	return &transactionsHandler{transactionsService}
}

func (h *transactionsHandler) Index(c *gin.Context) {
	transactions, err := h.transactionsService.GetAllDataTransactions()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"transactions": transactions})
}

func (h *transactionsHandler) Create(c *gin.Context) {
	var input transaction.Transaction

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "transaction_new.html", input)
		return
	}

	registerInput := transaction.Transaction{}
	registerInput.Nama = input.Nama
	registerInput.TanggalTransaksi = input.TanggalTransaksi
	registerInput.Keterangan = input.Keterangan
	registerInput.MingguKe = input.MingguKe
	registerInput.JumlahMasuk = input.JumlahMasuk

	_, err = h.transactionsService.CreateData(registerInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/transactions")
}

func (h *transactionsHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	registeredUser, err := h.transactionsService.GetByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := transaction.Transaction{}
	input.ID = registeredUser.ID
	input.Nama = registeredUser.Nama
	input.TanggalTransaksi = registeredUser.TanggalTransaksi
	input.Keterangan = registeredUser.Keterangan
	input.MingguKe = registeredUser.MingguKe
	input.JumlahMasuk = registeredUser.JumlahMasuk

	c.HTML(http.StatusOK, "transaction_edit.html", input)
}

func (h *transactionsHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input transaction.Transaction

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "transaction_edit.html", input)
		return
	}
	input.ID = id

	_, err = h.transactionsService.UpdateData(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/transactions")
}
