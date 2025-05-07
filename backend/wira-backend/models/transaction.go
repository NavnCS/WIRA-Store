package models

type Transaction struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	NFTID       uint   `json:"nft_id"`
	BuyerWallet string `json:"buyer_wallet"`
	Status      string `json:"status"`
}
