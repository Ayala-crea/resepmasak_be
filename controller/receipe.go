package controller

import (
	"Ayala-Crea/ResepBe/model"
	repo "Ayala-Crea/ResepBe/repository"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateReceipe(c *fiber.Ctx) error {
	// cek token header auth
	tokenStr := c.Get("login")
	if tokenStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header tidak ada")
	}

	// parse untuk mendapatkan id
	token, err := jwt.ParseWithClaims(tokenStr, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil // Ganti "secret_key" dengan kunci rahasia Anda
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	idUser := claims.IdUser

	var receipe model.Receipt

	if err := c.BodyParser(&receipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	receipe.IdUser = int(idUser)

	db := c.Locals("db").(*gorm.DB)

	if err := repo.InsertReceipt(db, receipe); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"code": http.StatusCreated, "success": true, "status": "success", "message": "Task berhasil disimpan", "data": receipe})
}

func GetAllReceipe(c *fiber.Ctx) error {
	// cek token header auth
	tokenStr := c.Get("login")
	if tokenStr == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header tidak ada")
	}

	// parse untuk mendapatkan id
	token, err := jwt.ParseWithClaims(tokenStr, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil // Ganti "secret_key" dengan kunci rahasia Anda
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	idUser := claims.IdUser
	var receipes model.Receipt
	db := c.Locals("db").(*gorm.DB)

	receipe, err := repo.GetAllReceipe(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data role, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	receipes.IdUser = int(idUser)

	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    receipe,
	}

	return c.Status(http.StatusOK).JSON(response)
}
func GetReceipeById(c *fiber.Ctx) error {
	// Cek token header autentikasi
	token := c.Get("login")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header tidak ada")
	}

	// mencari id dari data
	id := c.Query("recipe_id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID task tidak boleh kosong"})
	}

	db := c.Locals("db").(*gorm.DB)

	receipe, err := repo.GetReceipetById(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Data Tidak Ditemukan"})
	}

	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "data": receipe})
}
