package router

import (
	"OctavianoRyan25/GOSwagger/controller"

	_ "OctavianoRyan25/GOSwagger/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Car API
// @Version 1.0
// @Description This is a sampple service for managing cars
// @TermOfService http://swagger.io/terms
// @contact.name API SUPPORT
// @host localhost:8080
// @BasePath /

func StartServer() *gin.Engine  {
	router := gin.Default()

	router.POST("/cars", controller.CreateCar)
	router.GET("/cars", controller.GetAllCar)
	router.GET("/car/:id", controller.GetCarByID)
	router.PATCH("/cars/:id", controller.UpdateCarByID)
	router.DELETE("/cars/:id", controller.DeleteCarByID)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}