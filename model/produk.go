package model

import "time"

type Produk struct {
	ProdukID           string    `json:"produk_id,omitempty" bson:"_id,omitempty"`
	NamaProduk         string    `json:"nama_produk" bson:"nama_produk"`
	KategoriProdukID   string    `json:"kategori_produk_id" bson:"kategori_produk_id"`
	NamaKategoriProduk string    `json:"nama_kategori_produk" bson:"nama_kategori_produk"`
	SatuanProdukID     string    `json:"satuan_produk_id" bson:"satuan_produk_id"`
	NamaSatuanProduk   string    `json:"nama_satuan_produk" bson:"nama_satuan_produk"`
	Harga              string    `json:"harga" bson:"harga"`
	CreatedAt          time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type SatuanProduk struct {
	SatuanProdukID   string `json:"satuan_produk_id,omitempty" bson:"_id,omitempty"`
	NamaSatuanProduk string `json:"nama_satuan_produk" bson:"nama_satuan_produk"`
}

type KategoriProduk struct {
	KategoriProdukID   string `json:"kategori_produk_id,omitempty" bson:"_id,omitempty"`
	NamaKategoriProduk string `json:"nama_kategori_produk" bson:"nama_kategori_produk"`
}
