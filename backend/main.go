package main

import (
	"projetoapi/model"
	"projetoapi/routes"
	"projetoapi/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var identityKey = "id"

func init() {
	services.OpenDatabase()
	services.Db.AutoMigrate(&model.Access{})
	services.Db.AutoMigrate(&model.Worker{})
	services.Db.AutoMigrate(&model.Zone{})
	services.Db.AutoMigrate(&model.WorkersZone{})

	defer services.Db.Close()
}

func main() {

	services.FormatSwagger()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.New()
	router.Use(cors.New(corsConfig))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// AUTH
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	/* Zonas
	evaluation := router.Group("/api/v1/evaluation")
	evaluation.Use(services.AuthorizationRequired())
	{
		evaluation.POST("/", routes.AddEvaluation)
		evaluation.GET("/", routes.GetAllEvaluation)
		evaluation.GET("/:id", routes.GetEvaluationById)
		evaluation.PUT("/:id", routes.UpdateEvaluation)
		evaluation.DELETE("/:id", routes.DeleteEvaluation)
	} */

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/login", routes.GenerateToken)
		auth.POST("/register", routes.RegisterUser)
		auth.PUT("/refresh_token", services.AuthorizationRequired(), routes.RefreshToken)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8081")
}
