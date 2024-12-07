package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SupportHandler struct{}

func NewSupportHandler() *SupportHandler {
	return &SupportHandler{}
}

func (h *SupportHandler) GetSupport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "microservice support",
		"status":  "success",
	})
}
