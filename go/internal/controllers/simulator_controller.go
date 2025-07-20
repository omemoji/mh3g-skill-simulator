package controllers

import (
	"mh3g-skill-simulator/internal/models"
	"mh3g-skill-simulator/internal/repositories"
	"mh3g-skill-simulator/internal/services"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SimulatorController struct {
	Service services.SimulatorService
}

type Request struct {
	MaxResults int `json:"max_results"`
}

type HunterResponse struct {
	Head  string `json:"head"`
	Body  string `json:"body"`
	Arms  string `json:"arms"`
	Waist string `json:"waist"`
	Legs  string `json:"legs"`
}

type SimulatorResponse struct {
	Hunters []HunterResponse `json:"hunters"`
}

func NewSimulatorController() *SimulatorController {
	filePaths := repositories.FilePaths{
		FilePathHead:  os.Getenv("EQUIPMENT_HEAD"),
		FilePathBody:  os.Getenv("EQUIPMENT_BODY"),
		FilePathArm:   os.Getenv("EQUIPMENT_ARM"),
		FilePathWaist: os.Getenv("EQUIPMENT_WAIST"),
		FilePathLeg:   os.Getenv("EQUIPMENT_LEG"),
	}

	repo := repositories.NewSimulatorRepository(filePaths)

	service := services.SimulatorService{
		Repository: repo,
	}

	ctrl := SimulatorController{
		Service: service,
	}
	return &ctrl
}

func (c *SimulatorController) GetHunters(ctx *gin.Context) {
	maxResults, err := strconv.Atoi(ctx.Query("max_results"))
	if err != nil {
		maxResults = 200
	}
	hunters, err := c.Service.Execute(models.SearchQuery{MaxResults: maxResults})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []HunterResponse
	for _, hunter := range hunters {
		response = append(response, HunterResponse{
			Head:  hunter.Head.Name,
			Body:  hunter.Body.Name,
			Arms:  hunter.Arms.Name,
			Waist: hunter.Waist.Name,
			Legs:  hunter.Legs.Name,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"hunters": response,
	})
}
