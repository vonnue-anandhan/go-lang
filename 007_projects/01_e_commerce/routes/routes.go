package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vonnue-anandhan/e-commerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/viewproduct", controllers.SearchProduct())
	incomingRoutes.GET("/users/searchproduct", controllers.SearchProductByQuery())
}
