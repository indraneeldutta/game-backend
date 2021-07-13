package config

import (
	"github.com/game-backend/apis"
	gameservice "github.com/game-backend/services/game_service"
	mongoservice "github.com/game-backend/services/mongo_service"
	userservice "github.com/game-backend/services/user_service"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// InitializeApplicationConfig initialises the application with required instances and runs the server
func InitializeApplicationConfig() {
	if viper.GetString("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	securityConfig := secure.DefaultConfig()
	securityConfig.SSLRedirect = false
	securityConfig.ReferrerPolicy = "strict-origin-when-cross-origin"
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(secure.New(securityConfig))

	router.Use(gin.Recovery())

	mongoClient := mongoservice.NewMongoConnection()

	userService := userservice.NewUserService(mongoClient)
	gameservice := gameservice.NewGameService(mongoClient)

	v1 := router.Group("/v1")

	apis.NewUserController(v1, &userService)
	apis.NewGameStateController(v1, &gameservice)

	router.Run(viper.GetString("SERVER_PORT"))
}
