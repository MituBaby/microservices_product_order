package repository

import (
	"context"
	"product-service/config"
	"product-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	Create(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	GetByID(id string) (models.Product, error)
	Update(id string, product models.Product) error
	Delete(id string) error
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository() ProductRepository {
	// config.GetCollection mengambil collection "products" dari database "bip_product_db"
	return &productRepository{
		collection: config.DB.Collection("products"),
	}
}

func (r *productRepository) Create(product models.Product) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, product)
	return product, err
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var products []models.Product
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var p models.Product
		if err := cursor.Decode(&p); err == nil {
			products = append(products, p)
		}
	}
	return products, nil
}

func (r *productRepository) GetByID(id string) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product models.Product
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, err
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	return product, err
}

func (r *productRepository) Update(id string, product models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":  product.Name,
			"price": product.Price,
			"stock": product.Stock,
		},
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func (r *productRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}