package services

import (
	"errors"
	"math"

	"pipe-api/database"
	"pipe-api/dto"
	"pipe-api/models"

	"gorm.io/gorm"
)

func CreatePipe(userID uint, req dto.CreatePipeRequest) (*models.Pipe, error) {
	pipe := models.Pipe{
		UserID:      userID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	err := database.DB.Create(&pipe).Error
	return &pipe, err
}

func GetPipeByID(pipeID, userID uint) (*models.Pipe, error) {
	var pipe models.Pipe
	if err := database.DB.Where("id = ? AND user_id = ?", pipeID, userID).First(&pipe).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("pipe not found")
		}
		return nil, err
	}
	return &pipe, nil
}

func GetAllPipes(userID uint, page, limit int) (*dto.PaginatedPipes, error) {
	var pipes []models.Pipe
	var total int64
	offset := (page - 1) * limit

	query := database.DB.Model(&models.Pipe{}).Where("user_id = ?", userID)
	query.Count(&total)
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&pipes).Error; err != nil {
		return nil, err
	}

	data := make([]dto.PipeResponse, len(pipes))
	for i, p := range pipes {
		data[i] = dto.PipeResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			CreatedAt:   p.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   p.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return &dto.PaginatedPipes{
		Data: data,
		Meta: dto.PaginationMeta{
			Total:      total,
			Page:       page,
			Limit:      limit,
			TotalPages: totalPages,
		},
	}, nil
}

func UpdatePipe(pipeID, userID uint, req dto.UpdatePipeRequest) (*models.Pipe, error) {
	pipe, err := GetPipeByID(pipeID, userID)
	if err != nil {
		return nil, err
	}
	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Price > 0 {
		updates["price"] = req.Price
	}
	if req.Stock >= 0 {
		updates["stock"] = req.Stock
	}
	if len(updates) > 0 {
		database.DB.Model(&pipe).Updates(updates)
	}
	return GetPipeByID(pipeID, userID)
}

func DeletePipe(pipeID, userID uint) error {
	pipe, err := GetPipeByID(pipeID, userID)
	if err != nil {
		return err
	}
	return database.DB.Delete(&pipe).Error
}