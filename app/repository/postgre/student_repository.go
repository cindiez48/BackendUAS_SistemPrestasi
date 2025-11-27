package postgre

import (
    "database/sql"
    "backenduas_sistemprestasi/database"
    model "backenduas_sistemprestasi/app/models/postgre"
)

type StudentRepository struct {
    DB *sql.DB
}

func NewStudentRepository() *StudentRepository {
    return &StudentRepository{
        DB: database.PostgresDB,
    }
}

func (r *StudentRepository) FindByID(id string) (*model.Student, error) {
    query := `
        SELECT id, user_id, student_id, program_study,
               academic_year, advisor_id, created_at
        FROM students WHERE id=$1`

    row := r.DB.QueryRow(query, id)
    var s model.Student

    err := row.Scan(
        &s.ID, &s.UserID, &s.StudentID, &s.ProgramStudy,
        &s.AcademicYear, &s.AdvisorID, &s.CreatedAt,
    )

    return &s, err
}