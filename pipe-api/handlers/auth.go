package handlers

import (
	"net/http"
	"time"

	"pipe-api/config"
	"pipe-api/database"
	"pipe-api/dto"
	"pipe-api/models"
	"pipe-api/services"
	"pipe-api/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Регистрация
// @Description Регистрирует нового пользователя с email и паролем
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Данные регистрации"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.Register(req)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// @Summary Вход
// @Description Аутентифицирует пользователя, устанавливает HttpOnly cookie с токенами
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Учетные данные"
// @Success 200 {object} dto.UserResponse
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken, _ := utils.GenerateAccessToken(user.ID, user.Email)
	refreshTokenStr, _ := utils.GenerateRefreshToken(user.ID)
	refreshExp := time.Now().Add(time.Duration(config.AppConfig.JwtRefreshExpDays*24) * time.Hour)
	services.StoreRefreshToken(user.ID, refreshTokenStr, refreshExp)

	c.SetCookie("access_token", accessToken, config.AppConfig.JwtAccessExpMin*60, "/", "", false, true)
	c.SetCookie("refresh_token", refreshTokenStr, int(time.Until(refreshExp).Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	})
}

// @Summary Обновление токенов
// @Description Использует refresh_token из cookie для выдачи новой пары
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/refresh [post]
func Refresh(c *gin.Context) {
	refreshStr, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token not found"})
		return
	}
	userID, err := services.RefreshTokens(refreshStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	user, _ := services.GetUserByID(userID)
	accessToken, _ := utils.GenerateAccessToken(user.ID, user.Email)
	newRefresh, _ := utils.GenerateRefreshToken(user.ID)
	refreshExp := time.Now().Add(time.Duration(config.AppConfig.JwtRefreshExpDays*24) * time.Hour)
	services.StoreRefreshToken(user.ID, newRefresh, refreshExp)

	c.SetCookie("access_token", accessToken, config.AppConfig.JwtAccessExpMin*60, "/", "", false, true)
	c.SetCookie("refresh_token", newRefresh, int(time.Until(refreshExp).Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "tokens refreshed"})
}

// @Summary Профиль текущего пользователя
// @Security CookieAuth
// @Tags Auth
// @Produce json
// @Success 200 {object} dto.UserResponse
// @Failure 401 {object} map[string]string
// @Router /auth/whoami [get]
func Whoami(c *gin.Context) {
	userID := c.GetUint("userID")
	user, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	})
}

// @Summary Выход из сессии
// @Security CookieAuth
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /auth/logout [post]
func Logout(c *gin.Context) {
	refreshStr, _ := c.Cookie("refresh_token")
	if refreshStr != "" {
		hash, _ := utils.HashToken(refreshStr)
		database.DB.Model(&models.RefreshToken{}).Where("token_hash = ?", hash).Update("revoked", true)
	}
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

// @Summary Выход из всех сессий
// @Security CookieAuth
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /auth/logout-all [post]
func LogoutAll(c *gin.Context) {
	userID := c.GetUint("userID")
	services.RevokeAllUserTokens(userID)
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "all sessions terminated"})
}

// OAuth Yandex endpoints
func OAuthYandexLogin(c *gin.Context) {
	authURL, state := utils.GetYandexAuthURL()
	database.DB.Create(&models.OAuthState{ID: state})
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

func OAuthYandexCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	if state == "" || code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing code or state"})
		return
	}
	var st models.OAuthState
	if err := database.DB.Where("id = ?", state).First(&st).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid state"})
		return
	}
	database.DB.Delete(&st)

	tokenResp, err := utils.ExchangeYandexCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oauth code exchange failed"})
		return
	}
	userInfo, err := utils.GetYandexUserInfo(tokenResp.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user info"})
		return
	}

	// Find or create user
	var user models.User
	if err := database.DB.Where("yandex_id = ?", userInfo.ID).Or("email = ?", userInfo.DefaultEmail).First(&user).Error; err != nil {
		salt, _ := utils.GenerateSalt()
		user = models.User{
			Email:    userInfo.DefaultEmail,
			YandexID: &userInfo.ID,
			Salt:     salt,
		}
		database.DB.Create(&user)
	} else if user.YandexID == nil || *user.YandexID == "" {
		user.YandexID = &userInfo.ID
		database.DB.Save(&user)
	}

	accessToken, _ := utils.GenerateAccessToken(user.ID, user.Email)
	refreshTokenStr, _ := utils.GenerateRefreshToken(user.ID)
	refreshExp := time.Now().Add(time.Duration(config.AppConfig.JwtRefreshExpDays*24) * time.Hour)
	services.StoreRefreshToken(user.ID, refreshTokenStr, refreshExp)

	c.SetCookie("access_token", accessToken, config.AppConfig.JwtAccessExpMin*60, "/", "", false, true)
	c.SetCookie("refresh_token", refreshTokenStr, int(time.Until(refreshExp).Seconds()), "/", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}