package handler

import (
	"fmt"
	"net/http"
	"pet"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createSueta(c *gin.Context) {
	var input pet.Sueta

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "ошибка чтения body"})
	}

	id, status, err := h.services.Sueta.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) findAllSueta(c *gin.Context) {
	suetas, err := h.services.Sueta.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, suetas)
}

func (h *Handler) deleteSueta(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot read id"})
	}
	fmt.Println(id)
}
