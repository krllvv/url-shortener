package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/internal/entity"
	"url-shortener/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes(router *gin.Engine) {
	router.POST("/", h.ShortenURL)
	router.GET("/:alias", h.ExpandURL)
}

func (h *Handler) ShortenURL(c *gin.Context) {
	var input entity.URL
	if err := c.ShouldBindJSON(&input.Url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alias, err := h.service.SaveURL(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	aliasLink := fmt.Sprintf("%v://%v%v%v", scheme, c.Request.Host, c.Request.URL.Path, alias)

	c.String(http.StatusOK, aliasLink)
}

func (h *Handler) ExpandURL(c *gin.Context) {
	alias := c.Param("alias")

	url, err := h.service.GetURL(entity.URL{Alias: alias})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, url)
}
