package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

const (
	defaultPort       = "8423"
	defaultConfigPath = "../../config.yml"
)

var (
	port, confPath string
	engine         *gin.Engine
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)
	internal.InitConfig(confPath)
	logger.InitLogger()
	engine = internal.InitEngine(viper.GetString("gin.mode"))
}

func main() {
	logger.Info("Loading Swagger API")
	admin := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		viper.GetString("gin.login"): viper.GetString("gin.pass"),
	}))
	admin.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	internal.SetupGracefulShutdown(port, engine)
}
