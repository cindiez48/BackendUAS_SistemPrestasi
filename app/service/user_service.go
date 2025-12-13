package service

import (
	modelPostgre "backenduas_sistemprestasi/app/models/postgre"
	repoPg "backenduas_sistemprestasi/app/repository/postgre"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func GetAllUsers(c *fiber.Ctx) error {
	users, err := repoPg.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Gagal mengambil data user"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": users})
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := repoPg.UserFindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User tidak ditemukan"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": user})
}

func CreateUser(c *fiber.Ctx) error {
	var req modelPostgre.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	newUser := modelPostgre.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashed),
		FullName:     req.FullName,
		RoleID:       req.RoleID,
	}

	if err := repoPg.Create(newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Gagal membuat user (Duplicate username/email?)"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "User berhasil dibuat"})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var req modelPostgre.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	existingUser, err := repoPg.UserFindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User tidak ditemukan"})
	}

	if req.Username != "" { existingUser.Username = req.Username }
	if req.Email != "" { existingUser.Email = req.Email }
	if req.FullName != "" { existingUser.FullName = req.FullName }
	if req.IsActive != nil { existingUser.IsActive = *req.IsActive }

	if err := repoPg.Update(id, *existingUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Gagal update user"})
	}

	if req.Password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		repoPg.UpdatePassword(id, string(hashed))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User berhasil diupdate"})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repoPg.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Gagal menghapus user"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User berhasil dihapus"})
}

func AssignRole(c *fiber.Ctx) error {
	id := c.Params("id")
	var req modelPostgre.AssignRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	if err := repoPg.UpdateRole(id, req.RoleID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Gagal mengubah role"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Role berhasil diubah"})
}