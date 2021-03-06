// Package models provides 
package models

type Cart struct {
	Id           int
	Title        string
	Price        float64
	ProductVersion string
	Num          int
	ProductGift    string
	ProductFitting string
	ProductColor   string
	ProductImg     string
	ProductAttr    string
	Checked	bool `gorm:"_"`	//neglect
}

func (c *Cart) TableName() string {
	return "cart"
}

//If cart has data
func CartHasData(cartList []Cart, currentData Cart) bool {
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == currentData.Id &&
			cartList[i].ProductColor == currentData.ProductColor &&
			cartList[i].ProductAttr == currentData.ProductAttr {
			return true
		}
	}
	return false
}

