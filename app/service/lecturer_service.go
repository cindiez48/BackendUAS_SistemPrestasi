package service

import (
	"backenduas_sistemprestasi/helper"
	repoPg "backenduas_sistemprestasi/app/repository/postgre"

	"github.com/gofiber/fiber/v2"
)


func GetMyAdvisor(c *fiber.Ctx) error {

    loggedInUserID := c.Locals("user_id").(string)

	student, err := repoPg.FindStudentByUserID(loggedInUserID)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "message": "Data mahasiswa tidak ditemukan. Apakah anda login sebagai Mahasiswa?",
        })
    }

    if student.AdvisorName == nil {
         return c.Status(200).JSON(fiber.Map{
            "status": "success", 
            "message": "Anda belum memiliki Dosen Wali",
            "data": nil,
        })
    }

    return c.Status(200).JSON(fiber.Map{
        "status": "success", 
        "data": fiber.Map{
            "advisorName": *student.AdvisorName,
        },
    })
}

func GetAllLecturers(c *fiber.Ctx) error {
	lecturers, err := repoPg.FindAllLecturers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data dosen"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "data": lecturers})
}

func GetLecturerAdvisees(c *fiber.Ctx) error {
	lecturerID := c.Params("id")

	lecturer, err := repoPg.FindLecturerByID(lecturerID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Dosen tidak ditemukan"})
	}

	loggedInUserID, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	if !helper.IsAdmin(c) && lecturer.UserID != loggedInUserID {
		return c.Status(403).JSON(fiber.Map{"message": "Forbidden: Anda tidak boleh melihat bimbingan dosen lain"})
	}

	students, err := repoPg.FindLecturerAdvisees(lecturerID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data mahasiswa bimbingan"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "data": students})
}
