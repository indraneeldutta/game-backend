package config

import (
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

	_ = router.Group("/v1")

	router.Run(viper.GetString("SERVER_PORT"))
}
