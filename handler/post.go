package handler

import (
	"go-article/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.Service
}

func NewHandler(s service.Service) *handler {
	return &handler{s}
}

func (h *handler) GetAllPostHandler(c *gin.Context) {
	post, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   post,
	})

}
