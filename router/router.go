package router

import (
	"github.com/ade-iqbal/fga-assignment-3/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()
	
	router.LoadHTMLGlob("views/*.html")
	router.GET("/index", controllers.GetTemplate)

	return router
}