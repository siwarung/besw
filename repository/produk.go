package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/siwarung/besw/config"
	"github.com/siwarung/besw/model"
	"github.com/siwarung/besw/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Membuat kategori produk baru dengan validasi duplikasi
func CreateKategoriProduk(produk *model.KategoriProduk) (*mongo.InsertOneResult, error) {
	produkCollection := config.DB.Collection("kategori")

	// Cek apakah kategori sudah ada
	exists, err := CheckKategoriProduk(produk.NamaKategoriProduk)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("kategori sudah ada")
	}

	// Buat ID jika kosong
	if produk.KategoriProdukID == "" {
		produk.KategoriProdukID = primitive.NewObjectID().Hex()
	}

	// Simpan kategori ke database
	insertData := bson.M{
		"_id":                  produk.KategoriProdukID,
		"nama_kategori_produk": produk.NamaKategoriProduk,
	}

	result, err := produkCollection.InsertOne(context.Background(), insertData)
	if err != nil {
		fmt.Println("Error Insert:", err)
		return nil, err
	}
	return result, nil
}

// Periksa apakah kategori produk sudah ada
func CheckKategoriProduk(namaKategoriProduk string) (bool, error) {
	produkCollection := config.DB.Collection("kategori")
	count, err := produkCollection.CountDocuments(context.Background(), bson.M{"nama_kategori_produk": namaKategoriProduk})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Membuat satuan produk baru dengan validasi duplikasi
func CreateSatuanProduk(produk *model.SatuanProduk) (*mongo.InsertOneResult, error) {
	produkCollection := config.DB.Collection("satuan")

	// Cek apakah satuan sudah ada
	exists, err := CheckSatuanProduk(produk.NamaSatuanProduk)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("satuan produk sudah ada")
	}

	// Buat ID jika kosong
	if produk.SatuanProdukID == "" {
		produk.SatuanProdukID = primitive.NewObjectID().Hex()
	}

	// Simpan satuan ke database
	insertData := bson.M{
		"_id":                produk.SatuanProdukID,
		"nama_satuan_produk": produk.NamaSatuanProduk,
	}

	result, err := produkCollection.InsertOne(context.Background(), insertData)
	if err != nil {
		fmt.Println("Error Insert:", err)
		return nil, err
	}
	return result, nil
}

// Periksa apakah satuan produk sudah ada
func CheckSatuanProduk(namaSatuanProduk string) (bool, error) {
	produkCollection := config.DB.Collection("satuan")
	count, err := produkCollection.CountDocuments(context.Background(), bson.M{"nama_satuan_produk": namaSatuanProduk})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Cek apakah produk dengan nama yang mirip sudah ada di database
func IsSimilarProdukExists(namaProduk string) (bool, error) {
	produkCollection := config.DB.Collection("produk")

	cursor, err := produkCollection.Find(context.Background(), bson.M{},
		options.Find().SetProjection(bson.M{"nama_produk": 1}))
	if err != nil {
		return false, err
	}
	defer cursor.Close(context.Background())

	// Loop semua produk untuk dibandingkan dengan nama baru
	var produk model.Produk
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&produk)
		if err != nil {
			continue
		}

		// Cek kemiripan dengan nama produk yang ada
		if utils.IsSimilarString(produk.NamaProduk, namaProduk, 2) { // 2 = Maksimal 2 kesalahan huruf
			return true, nil
		}
	}

	return false, nil
}

// Membuat produk baru dengan pengecekan typo dan validasi duplikasi
func CreateProduk(produk *model.Produk) (*mongo.InsertOneResult, error) {
	produkCollection := config.DB.Collection("produk")

	// Cek apakah nama produk mirip dengan yang sudah ada
	similarExists, err := IsSimilarProdukExists(produk.NamaProduk)
	if err != nil {
		return nil, err
	}
	if similarExists {
		return nil, errors.New("nama produk sudah ada atau terlalu mirip")
	}

	// Set waktu CreatedAt dan UpdatedAt dengan time.Time
	if produk.ProdukID == "" {
		produk.ProdukID = primitive.NewObjectID().Hex()
	}
	produk.CreatedAt = time.Now()
	produk.UpdatedAt = time.Now()

	// Simpan produk ke database
	insertData := bson.M{
		"_id":                  produk.ProdukID,
		"nama_produk":          produk.NamaProduk,
		"kategori_produk_id":   produk.KategoriProdukID,
		"nama_kategori_produk": produk.NamaKategoriProduk,
		"satuan_produk_id":     produk.SatuanProdukID,
		"nama_satuan_produk":   produk.NamaSatuanProduk,
		"harga":                produk.Harga,
		"created_at":           primitive.NewDateTimeFromTime(produk.CreatedAt),
		"updated_at":           primitive.NewDateTimeFromTime(produk.UpdatedAt),
	}

	result, err := produkCollection.InsertOne(context.Background(), insertData)
	if err != nil {
		fmt.Println("Error Insert:", err)
		return nil, err
	}
	return result, nil
}

// Mengambil semua produk dari database
func GetAllProduk() ([]model.Produk, error) {
	produkCollection := config.DB.Collection("produk")

	cursor, err := produkCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var produkList []model.Produk
	for cursor.Next(context.Background()) {
		var produk model.Produk
		if err := cursor.Decode(&produk); err != nil {
			continue
		}
		produkList = append(produkList, produk)
	}

	return produkList, nil
}

// Mengambil semua kategori produk dari database
func GetAllKategoriProduk() ([]model.KategoriProduk, error) {
	kategoriCollection := config.DB.Collection("kategori")

	cursor, err := kategoriCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var kategoriList []model.KategoriProduk
	for cursor.Next(context.Background()) {
		var kategori model.KategoriProduk
		if err := cursor.Decode(&kategori); err != nil {
			continue
		}
		kategoriList = append(kategoriList, kategori)
	}

	return kategoriList, nil
}

// Mengambil semua satuan produk dari database
func GetAllSatuanProduk() ([]model.SatuanProduk, error) {
	satuanCollection := config.DB.Collection("satuan")

	cursor, err := satuanCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var satuanList []model.SatuanProduk
	for cursor.Next(context.Background()) {
		var satuan model.SatuanProduk
		if err := cursor.Decode(&satuan); err != nil {
			continue
		}
		satuanList = append(satuanList, satuan)
	}

	return satuanList, nil
}
