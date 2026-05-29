package main

import (
    "pipe-api/config"
    "pipe-api/database"
    "pipe-api/handlers"
    "pipe-api/middleware"
    
    _ "pipe-api/docs"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Plumbing Shop API
// @version 1.0
// @description API для управления магазином сантехнических труб (лабораторные работы 2-4).
// @description ## Аутентификация
// @description Используются JWT токены в HttpOnly cookie (access_token, refresh_token).
// @description Для тестирования защищённых эндпоинтов выполните `/auth/login` или `/auth/register`.
// @description ## OAuth2
// @description Поддерживается вход через Яндекс ID.
// @host localhost:4200
// @BasePath /
// @securityDefinitions.apikey CookieAuth
// @in cookie
// @name access_token

func main() {
	config.LoadConfig()

	// Подключаемся к БД
	database.ConnectDB()

	// Запускаем миграции
	database.RunMigrations()

	// Создаём роутер
	r := gin.Default()

	// Swagger UI только в development окружении
	if config.AppConfig.AppEnv != "production" {
		r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Публичные маршруты
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/refresh", handlers.Refresh)
		auth.GET("/oauth/yandex", handlers.OAuthYandexLogin)
		auth.GET("/oauth/yandex/callback", handlers.OAuthYandexCallback)
	}

	// Защищённые маршруты
	protected := r.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/auth/whoami", handlers.Whoami)
		protected.POST("/auth/logout", handlers.Logout)
		protected.POST("/auth/logout-all", handlers.LogoutAll)

		protected.GET("/pipes", handlers.GetPipes)
		protected.GET("/pipes/:id", handlers.GetPipe)
		protected.POST("/pipes", handlers.CreatePipe)
		protected.PUT("/pipes/:id", handlers.UpdatePipePut)
		protected.PATCH("/pipes/:id", handlers.UpdatePipePatch)
		protected.DELETE("/pipes/:id", handlers.DeletePipe)
	}

	r.Run(":" + config.AppConfig.AppPort)
}