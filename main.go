package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samridht23/safepasteServer/pkg/router"
	"github.com/samridht23/safepasteServer/pkg/utils"
)

func main() {
    gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	router.Init(app)
	app.Run(":" + utils.Getenv("PORT"))
}
