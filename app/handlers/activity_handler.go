package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zekhoi/golang-todo/app/dto"
	"github.com/zekhoi/golang-todo/app/entities"
	"github.com/zekhoi/golang-todo/app/services"
)

type ActivityHandler interface {
	GetAllActivities(c *gin.Context)
	GetActivity(c *gin.Context)
	CreateActivity(c *gin.Context)
	UpdateActivity(c *gin.Context)
	DeleteActivity(c *gin.Context)
}

type activityHandler struct {
	service services.ActivityService
}

func NewActivityHandler(service services.ActivityService) ActivityHandler {
	return &activityHandler{service: service}
}

func (h *activityHandler) GetAllActivities(c *gin.Context) {
	activities := h.service.GetAllActivities()

	res := entities.Response{
		// Status:  http.StatusOK,
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	}

	c.JSON(http.StatusOK, res)
}

func (h *activityHandler) GetActivity(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity := h.service.GetActivity(idInt)

	res := entities.Response{
		// Status:  http.StatusOK,
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	}

	c.JSON(http.StatusOK, res)
}

func (h *activityHandler) CreateActivity(c *gin.Context) {
	var data dto.ActivityCreateRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		res := entities.Response{
			// Status:  http.StatusBadRequest,
			Status:  "Error",
			Message: "Error",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	activity := h.service.CreateActivity(data)

	res := entities.Response{
		// Status:  http.StatusOK,
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	}

	c.JSON(http.StatusCreated, res)
}

func (h *activityHandler) UpdateActivity(c *gin.Context) {
	var data dto.ActivityUpdateRequest

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity := h.service.UpdateActivity(idInt, data)

	res := entities.Response{
		// Status:  http.StatusOK,
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	}

	c.JSON(http.StatusOK, res)
}

func (h *activityHandler) DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.DeleteActivity(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := entities.Response{
		// Status:  http.StatusOK,
		Status:  "Success",
		Message: "Success",
	}

	c.JSON(http.StatusNoContent, res)
}
