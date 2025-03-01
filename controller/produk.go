package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/siwarung/besw/model"
	"github.com/siwarung/besw/repository"
)

// Menambahkan kategori produk baru
func CreateKategoriProduk(c *fiber.Ctx) error {
	var kategori model.KategoriProduk

	// Parsing body request
	if err := c.BodyParser(&kategori); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format input tidak valid",
		})
	}

	// Simpan ke database
	_, err := repository.CreateKategoriProduk(&kategori)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Kategori produk berhasil ditambahkan",
		"kategori": kategori,
	})
}

// Menambahkan satuan produk baru
func CreateSatuanProduk(c *fiber.Ctx) error {
	var satuan model.SatuanProduk

	// Parsing body request
	if err := c.BodyParser(&satuan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format input tidak valid",
		})
	}

	// Simpan ke database
	_, err := repository.CreateSatuanProduk(&satuan)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Satuan produk berhasil ditambahkan",
		"satuan": satuan,
	})
}

// Menambahkan produk baru dengan validasi typo
func CreateProduk(c *fiber.Ctx) error {
	var produk model.Produk

	// Parsing body request
	if err := c.BodyParser(&produk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format input tidak valid",
		})
	}

	// Simpan ke database
	_, err := repository.CreateProduk(&produk)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Produk berhasil ditambahkan",
		"produk":  produk,
	})
}

// Mengambil daftar produk
func GetAllProduk(c *fiber.Ctx) error {
	produkList, err := repository.GetAllProduk()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data produk",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Daftar produk",
		"produk":  produkList,
	})
}

// Mengambil daftar kategori produk
func GetAllKategoriProduk(c *fiber.Ctx) error {
	kategoriList, err := repository.GetAllKategoriProduk()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data kategori produk",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Daftar kategori produk",
		"kategori":  kategoriList,
	})
}

// Mengambil daftar satuan produk
func GetAllSatuanProduk(c *fiber.Ctx) error {
	satuanList, err := repository.GetAllSatuanProduk()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data satuan produk",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Daftar satuan produk",
		"satuan":  satuanList,
	})
}

// Menghapus satuan produk
func DeleteSatuanProduk(c *fiber.Ctx) error {
	id := c.Params("id")

	// Hapus data dari database
	_, err := repository.DeleteSatuanProduk(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data satuan produk",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Satuan produk berhasil dihapus",
	})
}

// Menghapus kategori produk
func DeleteKategoriProduk(c *fiber.Ctx) error {
	id := c.Params("id")

	// Hapus data dari database
	_, err := repository.DeleteKategoriProduk(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data kategori produk",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Kategori produk berhasil dihapus",
	})
}

// Menghapus produk
func DeleteProduk(c *fiber.Ctx) error {
	id := c.Params("id")

	// Hapus data dari database
	_, err := repository.DeleteProduk(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data produk",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Produk berhasil dihapus",
	})
}