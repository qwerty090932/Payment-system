package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment_system/internal/usecase"
	"strconv"
)

type TransactionHandler struct {
	transactionUseCase *usecase.TransactionUseCase
}

func NewTransactionHandler(transactionUseCase *usecase.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{transactionUseCase: transactionUseCase}
}

func (h *TransactionHandler) Send(c *gin.Context) {
	var req struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.transactionUseCase.Send(req.From, req.To, req.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *TransactionHandler) GetLastTransactions(c *gin.Context) {
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid count parameter"})
		return
	}

	transactions, err := h.transactionUseCase.GetLastTransactions(count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
