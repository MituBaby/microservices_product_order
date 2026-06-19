package main

import (
	"log"
	"order-service/config"
	"order-service/handlers"
	"order-service/repository"
	"order-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// 1. Hubungkan ke database khusus order
	config.ConnectDB("bip_order_db")

	// 2. Inisialisasi Layer
	repo := repository.NewOrderRepository()
	svc := service.NewOrderService(repo)
	handler := handlers.NewOrderHandler(svc)

	// 3. Setup Fiber
	app := fiber.New()
	app.Use(cors.New())

	// 4. Endpoint sesuai dengan dokumen soal
	app.Post("/orders", handler.CreateOrder)

	// 5. Berjalan di port 8002
	log.Fatal(app.Listen(":8002"))
}