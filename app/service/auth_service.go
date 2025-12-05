package service

import (
	"time"

	"backenduas_sistemprestasi/helper"

	modelPostgre "backenduas_sistemprestasi/app/models/postgre"
	repository "backenduas_sistemprestasi/app/repository"
	memory "backenduas_sistemprestasi/memory"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var req modelPostgre.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	user, err := repository.FindByUsername(req.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Username atau password salah"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Username atau password salah"})
	}

	permissions, _ := repository.GetPermissionsByRoleID(user.RoleID)

	accessToken, _ := helper.GenerateJWT(user.ID, user.RoleID, permissions, time.Hour*1)
	refreshToken, _ := helper.GenerateJWT(user.ID, user.RoleID, permissions, time.Hour*24*7)

	response := modelPostgre.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		User: modelPostgre.UserDetail{
			ID:          user.ID,
			Username:    user.Username,
			FullName:    user.FullName,
			RoleID:      user.RoleID,
			Role:        user.RoleName,
			Permissions: permissions,
		},
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login berhasil",
		"data":    response,
	})
}

func Refresh(c *fiber.Ctx) error {
	var req modelPostgre.RefreshRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Format JSON salah. Gunakan key 'refreshToken'"})
	}

	if req.RefreshToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "RefreshToken wajib diisi"})
	}

	claims, err := helper.ValidateJWT(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Refresh token tidak valid atau expired"})
	}

	userID := claims["user_id"].(string)
	roleID := claims["role_id"].(string)

	var permissions []string
	if permInter, ok := claims["permissions"].([]interface{}); ok {
		for _, p := range permInter {
			permissions = append(permissions, p.(string))
		}
	}

	newAccessToken, _ := helper.GenerateJWT(userID, roleID, permissions, time.Hour*1)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"token": newAccessToken,
		},
	})
}

func Logout(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if len(authHeader) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Token invalid"})
	}

	tokenString := authHeader[7:]
	memory.AddToBlacklist(tokenString)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil logout",
	})
}

func Profile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	user, err := repository.UserFindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User tidak ditemukan"})
	}

	permissions, _ := repository.GetPermissionsByRoleID(user.RoleID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"id":          user.ID,
			"username":    user.Username,
			"fullName":    user.FullName,
			"email":       user.Email,
			"role":        user.RoleName,
			"permissions": permissions,
			"isActive":    user.IsActive,
			"joinedAt":    user.CreatedAt,
		},
	})
}
