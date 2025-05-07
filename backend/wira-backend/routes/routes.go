package routes

import (
	"wira-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/nfts", handlers.GetNFTs)
	api.Post("/favorites", handlers.ToggleFavorite)
	api.Post("/purchase", handlers.PurchaseNFT)
}
