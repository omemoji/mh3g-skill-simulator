package models

type Hunter struct {
	WeaponSlots int
	Head        *Equipment
	Body        *Equipment
	Arms        *Equipment
	Waist       *Equipment
	Legs        *Equipment
	Charm       *Charm
}
