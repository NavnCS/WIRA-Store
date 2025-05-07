package handlers

import (
	"wira-backend/database"
	"wira-backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetNFTs(c *fiber.Ctx) error {
	var nfts []models.NFT
	err := database.DB.Find(&nfts).Error
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(nfts)
}

type FavoriteInput struct {
	Wallet string `json:"wallet"`
	NFTID  int    `json:"nft_id"`
	Remove bool   `json:"remove"`
}

func ToggleFavorite(c *fiber.Ctx) error {
	var input FavoriteInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	var user models.User
	err := database.DB.FirstOrCreate(&user, models.User{WalletAddress: input.Wallet}).Error
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if input.Remove {
		// remove favorite
		newFavorites := []int{}
		for _, id := range user.Favorites {
			if id != input.NFTID {
				newFavorites = append(newFavorites, id)
			}
		}
		user.Favorites = newFavorites
	} else {
		user.Favorites = append(user.Favorites, input.NFTID)
	}

	err = database.DB.Save(&user).Error
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(200)
}

type PurchaseInput struct {
	NFTID int    `json:"nft_id"`
	Buyer string `json:"buyer_wallet"`
}

func PurchaseNFT(c *fiber.Ctx) error {
	var input PurchaseInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var nft models.NFT
		if err := tx.First(&nft, input.NFTID).Error; err != nil {
			return err
		}
		if nft.Owner != "" {
			return fiber.NewError(fiber.StatusBadRequest, "NFT already sold")
		}

		nft.Owner = input.Buyer
		if err := tx.Save(&nft).Error; err != nil {
			return err
		}

		transaction := models.Transaction{
			NFTID:       nft.ID,
			BuyerWallet: input.Buyer,
			Status:      "completed",
		}
		return tx.Create(&transaction).Error
	})

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}
