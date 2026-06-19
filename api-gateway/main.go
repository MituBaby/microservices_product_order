package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	// 1. Mengaktifkan CORS
	// Sangat Krusial! Agar Frontend (Port 3000) bisa menembak API tanpa diblokir oleh browser
	app.Use(cors.New())

	// Definisi alamat microservices internal kita
	productServiceURL := "http://localhost:8001"
	orderServiceURL   := "http://localhost:8002"

	// 2. Routing untuk Product Service
	// Semua HTTP Method (GET, POST, PUT, DELETE) ke /products akan diteruskan ke Port 8001
	app.All("/products*", func(c *fiber.Ctx) error {
		targetURL := productServiceURL + c.OriginalURL()
		return proxy.Do(c, targetURL)
	})

	// 3. Routing untuk Order Service
	// Semua HTTP Method ke /orders akan diteruskan ke Port 8002
	app.All("/orders*", func(c *fiber.Ctx) error {
		targetURL := orderServiceURL + c.OriginalURL()
		return proxy.Do(c, targetURL)
	})

	// 4. Jalankan API Gateway di Port 8000 sesuai instruksi lembar soal
	log.Println("API Gateway is running on port 8000...")
	log.Fatal(app.Listen(":8000"))
}