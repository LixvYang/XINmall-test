// Package models provides 
package models


type Banner struct {
	Id        int
	Title     string
	BannerType int
	BannerImg  string
	Link      string
	Sort      int
	Status    int
	AddTime   int
}

func (b *Banner) TableName() string {
	return "banner"
}
