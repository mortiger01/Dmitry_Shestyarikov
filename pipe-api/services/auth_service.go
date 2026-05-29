package services

import (
	"errors"
	"time"

	"pipe-api/database"
	"pipe-api/dto"
	"pipe-api/models"
	"pipe-api/utils"
)

func Register(req dto.RegisterRequest) (*dto.UserResponse, error) {
	var existing models.Pipe
	if err := database.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
		return nil, errors.New("email already exists")
	}
	salt, err := utils.GenerateSalt()
	if err != nil {
		return nil, err
	}
	hash, err := utils.HashPassword(req.Password, salt)
	if err != nil {
		return nil, err
	}
	user := models.User{
		Email:    req.Email,
		Password: hash,
		Salt:     salt,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func Login(req dto.LoginRequest) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}
	if !utils.CheckPassword(req.Password, user.Salt, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}

func StoreRefreshToken(userID uint, tokenString string, expiresAt time.Time) error {
	hash, err := utils.HashToken(tokenString)
	if err != nil {
		return err
	}
	rt := models.RefreshToken{
		UserID:    userID,
		TokenHash: hash,
		ExpiresAt: expiresAt,
	}
	return database.DB.Create(&rt).Error
}

func RefreshTokens(oldRefreshToken string) (uint, error) {
	claims, err := utils.ValidateRefreshToken(oldRefreshToken)
	if err != nil {
		return 0, errors.New("invalid refresh token")
	}
	hash, err := utils.HashToken(oldRefreshToken)
	if err != nil {
		return 0, err
	}
	var rt models.RefreshToken
	if err := database.DB.Where("token_hash = ? AND user_id = ? AND revoked = false AND expires_at > ?",
		hash, claims.UserID, time.Now()).First(&rt).Error; err != nil {
		return 0, errors.New("token invalid or revoked")
	}
	// revoke old token
	rt.Revoked = true
	database.DB.Save(&rt)
	return claims.UserID, nil
}

func RevokeAllUserTokens(userID uint) error {
	return database.DB.Model(&models.RefreshToken{}).
		Where("user_id = ? AND revoked = false", userID).
		Update("revoked", true).Error
}

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}