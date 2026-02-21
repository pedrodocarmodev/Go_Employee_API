package employee

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	emp, err := h.service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emp)
}

func (h *Handler) GetById(c *gin.Context) {
	ctx := c.Request.Context()

	str_id := c.Param("id")

	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id or not integer"})
		return
	}

	emp, err := h.service.GetById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, emp)
}

func (h *Handler) RegisterEmployee(c *gin.Context){
	ctx := c.Request.Context()

	var emp Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	id, err := h.service.RegisterEmployee(ctx, &emp)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, id)
}

func (h *Handler) Fire(c * gin.Context){
	ctx := c.Request.Context()

	str_id := c.Param("id")

	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id or not integer"})
		return
	}

	emp, err := h.service.Fire(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, emp)
}

func (h *Handler) Employ(c * gin.Context){
	ctx := c.Request.Context()

	str_id := c.Param("id")

	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id or not integer"})
		return
	}

	emp, err := h.service.Employ(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, emp)
}