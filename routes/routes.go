package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/number-classifier-api/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/classify-number", handlers.ClassifyNumber)
}
