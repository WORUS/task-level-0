package handler

import (
	"encoding/json"
	"net/http"
	"task-level-0/internal/domain/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.service.GetOrderById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var ord model.Order

	if err := json.Unmarshal(order, &ord); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erorr": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ord)
}
