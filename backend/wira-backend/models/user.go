package models

type User struct {
	WalletAddress string `json:"wallet_address" gorm:"primaryKey"`
	Favorites     []int  `gorm:"type:integer[]"`
}
