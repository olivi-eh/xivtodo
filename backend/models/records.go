package models

type IDdEntity struct {
	ID uint32
}

type NamedEntity struct {
	ID   uint32
	Name string

	NameEN string
	NameJA string
	NameDE string
	NameFR string
}

type GenderedEntity struct {
	ID   uint32
	Name string

	NameMasculineEN string
	NameMasculineJA string
	NameMasculineDE string
	NameMasculineFR string
	NameFeminineEN  string
	NameFeminineJA  string
	NameFeminineDE  string
	NameFeminineFR  string
}

type IconedNamedEntity struct {
	NamedEntity
	Icon string
}
