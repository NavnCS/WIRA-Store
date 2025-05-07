package models

type NFT struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Rarity string `json:"rarity"`
	Image  string `json:"image"`
	Price  int64  `json:"price"`
	Owner  string `json:"owner"`
}
