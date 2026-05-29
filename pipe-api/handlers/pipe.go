package handlers

import (
	"net/http"
	"strconv"

	"pipe-api/database"
	"pipe-api/dto"
	"pipe-api/models"
	"pipe-api/services"
	"github.com/gin-gonic/gin"
)

// @Summary Список труб
// @Description Возвращает пагинированный список труб текущего пользователя
// @Tags Pipes
// @Security CookieAuth
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Записей на странице" default(10)
// @Success 200 {object} dto.PaginatedPipes
// @Failure 401 {object} map[string]string
// @Router /pipes [get]
func GetPipes(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	userID := c.GetUint("userID")
	result, err := services.GetAllPipes(userID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Труба по ID
// @Tags Pipes
// @Security CookieAuth
// @Produce json
// @Param id path int true "ID трубы"
// @Success 200 {object} dto.PipeResponse
// @Failure 404 {object} map[string]string
// @Router /pipes/{id} [get]
func GetPipe(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userID := c.GetUint("userID")
	pipe, err := services.GetPipeByID(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pipeToResponse(*pipe))
}

// @Summary Создать трубу
// @Tags Pipes
// @Security CookieAuth
// @Accept json
// @Produce json
// @Param request body dto.CreatePipeRequest true "Данные трубы"
// @Success 201 {object} dto.PipeResponse
// @Router /pipes [post]
func CreatePipe(c *gin.Context) {
	var req dto.CreatePipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint("userID")
	pipe, err := services.CreatePipe(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pipeToResponse(*pipe))
}

// @Summary Полное обновление трубы
// @Tags Pipes
// @Security CookieAuth
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body dto.CreatePipeRequest true "Новые данные"
// @Success 200 {object} dto.PipeResponse
// @Router /pipes/{id} [put]
func UpdatePipePut(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var req dto.CreatePipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint("userID")
	pipe, err := services.GetPipeByID(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	pipe.Name = req.Name
	pipe.Description = req.Description
	pipe.Price = req.Price
	pipe.Stock = req.Stock
	database.DB.Save(&pipe)
	c.JSON(http.StatusOK, pipeToResponse(*pipe))
}

// @Summary Частичное обновление трубы
// @Tags Pipes
// @Security CookieAuth
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body dto.UpdatePipeRequest true "Поля для обновления"
// @Success 200 {object} dto.PipeResponse
// @Router /pipes/{id} [patch]
func UpdatePipePatch(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var req dto.UpdatePipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint("userID")
	pipe, err := services.UpdatePipe(uint(id), userID, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pipeToResponse(*pipe))
}

// @Summary Удалить трубу (soft delete)
// @Tags Pipes
// @Security CookieAuth
// @Param id path int true "ID"
// @Success 204
// @Router /pipes/{id} [delete]
func DeletePipe(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userID := c.GetUint("userID")
	if err := services.DeletePipe(uint(id), userID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func pipeToResponse(p models.Pipe) dto.PipeResponse {
	return dto.PipeResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		CreatedAt:   p.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   p.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}