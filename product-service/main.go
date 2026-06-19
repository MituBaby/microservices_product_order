package main

import (
	"log"
	"product-service/config"
	"product-service/handlers"
	"product-service/repository"
	"product-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// 1. Hubungkan ke Database khusus produk
	config.ConnectDB("bip_product_db")

	// 2. Inisialisasi Layer (Dependency Injection)
	repo := repository.NewProductRepository()
	svc := service.NewProductService(repo)
	handler := handlers.NewProductHandler(svc)

	// 3. Setup Fiber App
	app := fiber.New()
	app.Use(cors.New())

	// 4. Daftarkan Urutan Routes sesuai Dokumen Soal
	app.Post("/products", handler.CreateProduct)
	app.Get("/products", handler.GetAllProducts)
	app.Get("/products/:id", handler.GetProductByID)
	app.Put("/products/:id", handler.UpdateProduct)
	app.Delete("/products/:id", handler.DeleteProduct)

	// 5. Jalankan di port 8001
	log.Fatal(app.Listen(":8001"))
}