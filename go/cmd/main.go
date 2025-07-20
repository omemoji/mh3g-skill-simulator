package main

import (
	"mh3g-skill-simulator/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	// コントローラーの初期化
	r := gin.Default()

	// APIエンドポイントの設定
	api := r.Group("/api/v1")
	api.GET("/", controllers.NewSimulatorController().GetHunters)

	r.Run(":8080")
}
