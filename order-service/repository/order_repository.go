package repository

import (
	"context"
	"order-service/config"
	"order-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	Create(order models.Order) (models.Order, error)
}

type orderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{
		collection: config.DB.Collection("orders"),
	}
}

func (r *orderRepository) Create(order models.Order) (models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order.ID = primitive.NewObjectID()
	order.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	if order.Status == "" {
		order.Status = "PENDING"
	}

	_, err := r.collection.InsertOne(ctx, order)
	return order, err
}