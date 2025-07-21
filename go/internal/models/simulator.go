package models

type Simulator struct {
	EquipmentLists [5][]Equipment
}

type SearchQuery struct {
	MaxResults int
}

func (s *Simulator) Simulate(searchQuery SearchQuery) ([]Hunter, error) {

	var hunters []Hunter
	// 全ての装備の組み合わせを生成（200件まで）
	for _, head := range s.EquipmentLists[0] {
		for _, body := range s.EquipmentLists[1] {
			for _, arms := range s.EquipmentLists[2] {
				for _, waist := range s.EquipmentLists[3] {
					for _, legs := range s.EquipmentLists[4] {
						hunter := Hunter{
							Head:  &head,
							Body:  &body,
							Arms:  &arms,
							Waist: &waist,
							Legs:  &legs,
						}
						hunters = append(hunters, hunter)
						if len(hunters) >= searchQuery.MaxResults {
							return hunters, nil
						}
					}
				}
			}
		}
	}
	return hunters, nil

}

type SimulatorRepository interface {
	Search() (Simulator, error)
}
