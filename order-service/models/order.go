package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ProductID string             `json:"productId" bson:"product_id"`
	Quantity  int                `json:"quantity" bson:"quantity"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"created_at,omitempty"`
}