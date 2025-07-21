package services

import "mh3g-skill-simulator/internal/models"

type SimulatorService struct {
	Repository models.SimulatorRepository
}

type SimulatorServiceInterface interface {
	Execute(searchQuery models.SearchQuery) ([]models.Hunter, error)
}

func (s *SimulatorService) Execute(searchQuery models.SearchQuery) ([]models.Hunter, error) {
	simulator, err := s.Repository.Search()
	if err != nil {
		return nil, err
	}
	hunters, err := simulator.Simulate(searchQuery)
	if err != nil {
		return nil, err
	}
	return hunters, nil
}
