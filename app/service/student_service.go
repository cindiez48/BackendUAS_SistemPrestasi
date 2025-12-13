package service

import (
	"backenduas_sistemprestasi/helper"
	modelPostgre "backenduas_sistemprestasi/app/models/postgre"
	repo "backenduas_sistemprestasi/app/repository/postgre"
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

    nama_role := c.Locals("role_name")
    if nama_role == "Mahasiswa" {
        return c.Status(403).JSON(fiber.Map{
            "message": "anda bukan seorang admin maupun dosen",
        })
    }
    
    if nama_role == "Dosen Wali" {
    
    }

    students, err := repo.StudentFindAll()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Gagal mengambil data mahasiswa",
            "error":   err.Error(),
        })
    }
    return c.Status(200).JSON(fiber.Map{"status": "success", "data": students})
}


func StudentFindByID(c *fiber.Ctx) error {
    id := c.Params("id")

    student, err := repo.StudentFindByID(id)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"message": "Mahasiswa tidak ditemukan"})
    }

    loggedInUserID := c.Locals("user_id").(string)

    if !helper.IsAdmin(c) && student.UserID != loggedInUserID {
        return c.Status(403).JSON(fiber.Map{"message": "Forbidden: Anda tidak boleh melihat profil mahasiswa lain"})
    }

    return c.Status(200).JSON(fiber.Map{"status": "success", "data": student})
}

func AssignAdvisor(c *fiber.Ctx) error {
    id := c.Params("id")

    var req modelPostgre.AssignAdvisorRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Invalid input"})
    }

    student, err := repo.StudentFindByID(id)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"message": "Mahasiswa tidak ditemukan"})
    }

    loggedInUserID := c.Locals("user_id").(string)
    if !helper.IsAdmin(c) && student.UserID != loggedInUserID {
        return c.Status(403).JSON(fiber.Map{"message": "Forbidden: Anda hanya boleh memilih dosen wali untuk diri sendiri"})
    }

    if err := repo.UpdateAdvisor(id, req.LecturerID); err != nil {
        return c.Status(500).JSON(fiber.Map{"message": "Gagal set dosen wali"})
    }
    return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Dosen Wali berhasil diupdate"})
}