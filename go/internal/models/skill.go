package models

type Skill struct {
	Name     string
	Category string
	Points   int
	Type     int // 0=両方, 1=剣士, 2=ガンナー
}
