package router

import (
	"github.com/gin-gonic/gin"
	"github.com/samridht23/safepasteServer/pkg/controllers"
	"github.com/samridht23/safepasteServer/pkg/middleware"
)

func Init(app *gin.Engine) {
	app.Use(middleware.CorsMiddleware())
    app.GET("/",controller.GetData)
	app.POST("/", controller.UploadData)
    return 
}
