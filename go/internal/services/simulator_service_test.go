package services_test

import (
	"mh3g-skill-simulator/internal/models"
	"mh3g-skill-simulator/internal/services"
	"testing"
)

type MockSimulatorRepository struct{}

func (m *MockSimulatorRepository) Search() (models.Simulator, error) {

	equipmentListHead := []models.Equipment{
		{
			Name:           "Head1",
			InitialDefense: 10,
			FinalDefense:   20,
			Skills:         [5]models.Skill{{Name: "Skill1", Points: 1}},
		},
		{
			Name:           "Head2",
			InitialDefense: 15,
			FinalDefense:   25,
			Skills:         [5]models.Skill{{Name: "Skill2", Points: 2}},
		},
	}
	equipmentListBody := []models.Equipment{
		{
			Name:           "Body1",
			InitialDefense: 20,
			FinalDefense:   30,
			Skills:         [5]models.Skill{{Name: "Skill3", Points: 3}},
		},
		{
			Name:           "Body2",
			InitialDefense: 25,
			FinalDefense:   35,
			Skills:         [5]models.Skill{{Name: "Skill4", Points: 4}},
		},
	}
	equipmentListArm := []models.Equipment{
		{
			Name:           "Arm1",
			InitialDefense: 30,
			FinalDefense:   40,
			Skills:         [5]models.Skill{{Name: "Skill5", Points: 5}},
		},
		{
			Name:           "Arm2",
			InitialDefense: 35,
			FinalDefense:   45,
			Skills:         [5]models.Skill{{Name: "Skill6", Points: 6}},
		},
	}
	equipmentListWaist := []models.Equipment{
		{
			Name:           "Waist1",
			InitialDefense: 40,
			FinalDefense:   50,
			Skills:         [5]models.Skill{{Name: "Skill7", Points: 7}},
		},
		{
			Name:           "Waist2",
			InitialDefense: 45,
			FinalDefense:   55,
			Skills:         [5]models.Skill{{Name: "Skill8", Points: 8}},
		},
	}
	equipmentListLeg := []models.Equipment{
		{
			Name:           "Leg1",
			InitialDefense: 50,
			FinalDefense:   60,
			Skills:         [5]models.Skill{{Name: "Skill9", Points: 9}},
		},
		{
			Name:           "Leg2",
			InitialDefense: 55,
			FinalDefense:   65,
			Skills:         [5]models.Skill{{Name: "Skill10", Points: 10}},
		},
	}
	return models.Simulator{
		EquipmentLists: [5][]models.Equipment{
			equipmentListHead,
			equipmentListBody,
			equipmentListArm,
			equipmentListWaist,
			equipmentListLeg,
		},
	}, nil
}

func TestSimulatorService_Execute(t *testing.T) {
	repo := &MockSimulatorRepository{}
	service := services.SimulatorService{Repository: repo}

	query := models.SearchQuery{MaxResults: 20}
	hunters, err := service.Execute(query)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(hunters) == 0 {
		t.Fatal("expected to find at least one hunter, got none")
	}

	for _, hunter := range hunters {
		if hunter.Head == nil || hunter.Body == nil || hunter.Arms == nil || hunter.Waist == nil || hunter.Legs == nil {
			t.Error("expected all equipment to be assigned, got nil")
		}
		if hunter.Head.FinalDefense <= hunter.Head.InitialDefense {
			t.Error("expected Head FinalDefense to be greater than InitialDefense")
		}
		if len(hunter.Head.Skills) != 5 {
			t.Error("expected Head to have 5 skills, got none")
		}
	}
	if len(hunters) != 20 {
		t.Error("expected 20 hunters, got ", len(hunters))
	}
}
