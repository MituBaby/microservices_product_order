package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB(dbName string) {
	// Gunakan URI default MongoDB lokal
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping database untuk memastikan koneksi berhasil
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	
	// Inisialisasi database spesifik untuk service ini
	DB = client.Database(dbName)
}

// Fungsi helper untuk mendapatkan collection
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}