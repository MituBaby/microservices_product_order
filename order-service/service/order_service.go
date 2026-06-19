package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"order-service/models"
	"order-service/repository"
)

type OrderService interface {
	CreateOrder(order models.Order) (models.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

// Struct bantuan untuk membaca data dari Product Service
type ProductResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func (s *orderService) CreateOrder(order models.Order) (models.Order, error) {
	// 1. HIT PRODUCT SERVICE: Cek detail produk ke Port 8001
	resp, err := http.Get(fmt.Sprintf("http://localhost:8001/products/%s", order.ProductID))
	if err != nil || resp.StatusCode == http.StatusNotFound {
		return order, errors.New("product not found or product service is down")
	}
	defer resp.Body.Close()

	var product ProductResponse
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
    	return order, err
	}

	// 2. VALIDASI STOK: Apakah stok di Product Service mencukupi?
	if product.Stock < order.Quantity {
		return order, errors.New("insufficient product stock")
	}

	// 3. HIT PRODUCT SERVICE: Kurangi Stok via PUT /products/:id
	product.Stock -= order.Quantity
	requestBody, _ := json.Marshal(product)
	
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8001/products/%s", order.ProductID), bytes.NewBuffer(requestBody))
	if err != nil {
		return order, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	putResp, err := client.Do(req)
	if err != nil || putResp.StatusCode != http.StatusOK {
		return order, errors.New("failed to update product stock")
	}
	defer putResp.Body.Close()

	// 4. SIMPAN ORDER: Jika semua sukses, simpan data order ke DB Order
	order.Status = "SUCCESS"
	return s.repo.Create(order)
}