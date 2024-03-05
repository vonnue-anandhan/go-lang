package controllers

import "github.com/gin-gonic/gin"

func VerifyPassword(userPassword, givenPassword string) (bool, string) {}

func HashPassword(password string) string {
}

func SignUp() gin.HandlerFunc {}

func Login() gin.HandlerFunc {}

func ProductViewerAdmin() gin.HandlerFunc {}

func SearchProduct() gin.HandlerFunc {}

func SearchProductByQuery() gin.HandlerFunc {}
