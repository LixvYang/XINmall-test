// Package models provides 
package models

type OrderItem struct {
	Id           int
	OrderId      int
	Uid          int
	ProductTitle string
	ProductId    int
	ProductImg   string
	ProductPrice float64
	ProductNum   int
	ProductVersion string
	ProductColor   string
	AddTime      int
}

func (o *OrderItem) TableName() string {
	return "order_item"
}
