package main

import (
	"mh3g-skill-simulator/internal/controllers"
	"mh3g-skill-simulator/internal/repositories"
	"mh3g-skill-simulator/internal/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	filePaths := repositories.FilePaths{
		FilePathHead:  os.Getenv("EQUIPMENT_HEAD"),
		FilePathBody:  os.Getenv("EQUIPMENT_BODY"),
		FilePathArm:   os.Getenv("EQUIPMENT_ARM"),
		FilePathWaist: os.Getenv("EQUIPMENT_WAIST"),
		FilePathLeg:   os.Getenv("EQUIPMENT_LEG"),
	}

	repo := repositories.NewSimulatorRepository(filePaths)
	service := &services.SimulatorService{
		Repository: repo,
	}
	ctrl := controllers.NewSimulatorController(service)

	r := gin.Default()

	api := r.Group("/api/v1")
	api.GET("/", ctrl.GetHunters)

	r.Run(":8080")
}
