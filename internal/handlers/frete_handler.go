package handlers

import (
	"net/http"

	"github.com/Gierdiaz/internal/domain/entities"
	"github.com/Gierdiaz/internal/services"
	"github.com/gin-gonic/gin"
)

type FreteHandler struct {
	CorreiosService services.CorreiosServiceImpl
}

func NewFreteHandler(correiosService services.CorreiosServiceImpl) *FreteHandler {
	return &FreteHandler{CorreiosService: correiosService}
}

func (h *FreteHandler) CalcularFrete(c *gin.Context) {
	var frete entities.Frete

	if err := c.ShouldBindJSON(&frete); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	calculatedFrete, err := h.CorreiosService.CalcularFrete(&frete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, calculatedFrete)
}
