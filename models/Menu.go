// Package models provides 
package models


type Menu struct {
	Id        int
	Title     string
	Link      string
	Position  int
	IsOpennew int
	Relation  string
	Sort      int
	Status    int
	AddTime   int
	ProductItem []Product `gorm:"-"`
}

func (m *Menu) TableName() string {
	return "menu"
}
