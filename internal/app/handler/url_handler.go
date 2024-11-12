package handler

import (
	"net/http"
	"proyectogo/internal/app"
	"proyectogo/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type URLHandler struct {
	service app.URLService
}

func NewURLHandler(service app.URLService) *URLHandler {
	return &URLHandler{service}
}

func (h *URLHandler) CreateURL(c *gin.Context) {
	var url domain.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateURL(&url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, url)
}

func (h *URLHandler) GetURLByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	url, err := h.service.GetURLByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL no encontrada"})
		return
	}
	c.JSON(http.StatusOK, url)
}

func (h *URLHandler) GetAllURLs(c *gin.Context) {
	urls, err := h.service.GetAllURLs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, urls)
}

func (h *URLHandler) UpdateURL(c *gin.Context) {
	var url domain.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateURL(&url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, url)
}

func (h *URLHandler) DeleteURL(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := h.service.DeleteURL(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
