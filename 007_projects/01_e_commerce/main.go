package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vonnue-anandhan/e-commerce/controllers"
	"github.com/vonnue-anandhan/e-commerce/database"
	"github.com/vonnue-anandhan/e-commerce/middleware"
	"github.com/vonnue-anandhan/e-commerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middleware.Authentication)

	routes.UserRoutes(router)

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
