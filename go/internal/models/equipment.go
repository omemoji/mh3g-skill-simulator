package models

type Material struct {
	Name   string
	Amount int
}

type Equipment struct {
	Parts            int // 0=頭, 1=胴, 2=腕, 3=腰, 4=脚
	Name             string
	Gender           int // 0=両, 1=男, 2=女
	Type             int // 0=両方, 1=剣士, 2=ガンナー
	Rarity           int
	Slots            int
	PortAvailable    bool // true=港入手可能, false=港入手不可
	VillageAvailable bool // true=村入手可能, false=村入手不可
	InitialDefense   int
	FinalDefense     int
	Resistance       [5]int // Fire, Water, Ice, Thunder, Dragon
	Skills           [5]Skill
	Materials        [4]Material
}
