package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/williamluisan/golang_pdfgenerator/controllers"
)

func Routes(router *gin.Engine) {
	router.POST("/generate", controllers.Request)
}