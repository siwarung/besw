package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var DB *mongo.Database

func ConnectDB() {
	// Load .env hanya jika berjalan di lokal
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Tidak dapat membaca file .env, menggunakan environment Heroku")
		}
	}

	mongoURI := os.Getenv("MONGO_URL")
	dbName := os.Getenv("MONGO_DB")

	if mongoURI == "" || dbName == "" {
		log.Fatal("MONGO_URL dan MONGO_DB tidak boleh kosong")
	}

	// Atur opsi koneksi
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Koneksi ke MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Gagal tersambung ke Database", err)
	}

	// Cek koneksi ke MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Koneksi ke Database gagal", err)
	}

	log.Println("Berhasil tersambung ke Database")

	MongoClient = client
	DB = client.Database(dbName)
}

// Mengambil koleksi
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
