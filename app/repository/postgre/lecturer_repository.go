package postgre

import (
    "database/sql"
    "backenduas_sistemprestasi/database"
    model "backenduas_sistemprestasi/app/models/postgre"
)

type LecturerRepository struct {
    DB *sql.DB
}

func NewLecturerRepository() *LecturerRepository {
    return &LecturerRepository{
        DB: database.PostgresDB,
    }
}

func (r *LecturerRepository) GetAll() ([]model.Lecturer, error) {
    query := `SELECT id, user_id, lecturer_id, department, created_at FROM lecturers`
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []model.Lecturer

    for rows.Next() {
        var l model.Lecturer
        err := rows.Scan(&l.ID, &l.UserID, &l.LecturerID, &l.Department, &l.CreatedAt)
        if err != nil {
            return nil, err
        }
        list = append(list, l)
    }

    return list, nil
}