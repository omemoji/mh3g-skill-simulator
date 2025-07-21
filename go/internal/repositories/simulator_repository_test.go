package repositories

import (
	"testing"
)

func TestSimulatorRepository_Search(t *testing.T) {
	filePaths := FilePaths{
		FilePathHead:  "../test_data/test_head.csv",
		FilePathBody:  "../test_data/test_body.csv",
		FilePathArm:   "../test_data/test_arm.csv",
		FilePathWaist: "../test_data/test_wst.csv",
		FilePathLeg:   "../test_data/test_leg.csv",
	}
	repo := NewSimulatorRepository(filePaths)

	simulator, err := repo.Search()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(simulator.EquipmentLists) == 0 {
		t.Fatal("expected to find at least one equipment, got none")
	}
	for i, equipmentList := range simulator.EquipmentLists {
		if len(equipmentList) == 0 {
			t.Errorf("expected equipment list %d to have items, got none", i)
		}
		for _, equipment := range equipmentList {
			if equipment.Name == "" {
				t.Error("expected equipment to have a name, got empty")
			}
			if equipment.FinalDefense <= equipment.InitialDefense {
				t.Error("expected FinalDefense to be greater than InitialDefense")
			}
			if len(equipment.Skills) != 5 {
				t.Error("expected equipment to have 5 skills, got none")
			}
		}
	}
}
