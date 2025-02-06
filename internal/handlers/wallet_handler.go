package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment_system/internal/usecase"
)

type WalletHandler struct {
	walletUseCase *usecase.WalletUseCase
}

func NewWalletHandler(walletUseCase *usecase.WalletUseCase) *WalletHandler {
	return &WalletHandler{walletUseCase: walletUseCase}
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	address := c.Param("address")
	balance, err := h.walletUseCase.GetBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
