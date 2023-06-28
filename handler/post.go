package handler

import (
	"go-article/model"
	"go-article/response"
	"go-article/service"
	"net/http"
	"strconv"

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

	var postsResponse []response.PostResponse

	for _, p := range post {
		postResponse := GenerateResponse(p)
		postsResponse = append(postsResponse, postResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   postsResponse,
	})
}

func (h *handler) GetSinglePostHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	postResponse := GenerateResponse(post)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   postResponse,
	})
}

func GenerateResponse(p model.Post) response.PostResponse {
	postResponse := response.PostResponse{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Status:      p.Status,
		AdminID:     p.AdminID,
	}

	admin := make(map[string]interface{})
	admin["id"] = p.Admin.ID
	admin["name"] = p.Admin.Name
	admin["email"] = p.Admin.Email
	postResponse.Admin = admin

	var comments []map[string]interface{}
	for _, c := range p.Comments {
		comment := make(map[string]interface{})
		comment["id"] = c.ID
		comment["name"] = c.Admin.Name
		comment["comment"] = c.Comment
		comments = append(comments, comment)
	}

	postResponse.Comments = comments
	return postResponse
}
