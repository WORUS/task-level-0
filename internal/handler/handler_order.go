package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrder(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{"text": "All good"})
}
