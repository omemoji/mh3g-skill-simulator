package repositories

// csv からデータを読み込むリポジトリの実装
import (
	"encoding/csv"
	"mh3g-skill-simulator/internal/models"
	"os"
	"strconv"
)

type FilePaths struct {
	FilePathHead  string
	FilePathBody  string
	FilePathArm   string
	FilePathWaist string
	FilePathLeg   string
}

type SimulatorRepository struct {
	FilePaths FilePaths
	response  models.Simulator
}

func NewSimulatorRepository(filePaths FilePaths) *SimulatorRepository {
	return &SimulatorRepository{
		FilePaths: filePaths,
	}
}

func MustAtoi(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func DecideAvailable(value string) bool {
	return value != "99"
}

func (r *SimulatorRepository) Search() (models.Simulator, error) {
	for i, filePath := range []string{
		r.FilePaths.FilePathHead,
		r.FilePaths.FilePathBody,
		r.FilePaths.FilePathArm,
		r.FilePaths.FilePathWaist,
		r.FilePaths.FilePathLeg,
	} {
		equipmentList, err := processFile(filePath)
		if err != nil {
			return models.Simulator{}, err
		}
		r.response.EquipmentLists[i] = equipmentList
	}
	return r.response, nil
}

func processFile(filePath string) ([]models.Equipment, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var equipmentList []models.Equipment
	for _, record := range records {
		if record[0][0] == '#' {
			continue
		}
		equipment := models.Equipment{
			Name:             record[0],
			Gender:           MustAtoi(record[1]),
			Type:             MustAtoi(record[2]),
			Rarity:           MustAtoi(record[3]),
			Slots:            MustAtoi(record[4]),
			PortAvailable:    DecideAvailable(record[5]),
			VillageAvailable: DecideAvailable(record[6]),
			InitialDefense:   MustAtoi(record[7]),
			FinalDefense:     MustAtoi(record[8]),
			Resistance:       [5]int{MustAtoi(record[9]), MustAtoi(record[10]), MustAtoi(record[11]), MustAtoi(record[12]), MustAtoi(record[13])},
			Skills: [5]models.Skill{
				{Name: record[14], Points: MustAtoi(record[15])},
				{Name: record[16], Points: MustAtoi(record[17])},
				{Name: record[18], Points: MustAtoi(record[19])},
				{Name: record[20], Points: MustAtoi(record[21])},
				{Name: record[22], Points: MustAtoi(record[23])},
			},
		}
		equipmentList = append(equipmentList, equipment)
	}
	return equipmentList, nil
}
