package postgre

import (
    "database/sql"
    "backenduas_sistemprestasi/database"
    model "backenduas_sistemprestasi/app/models/postgre"
)

type RoleRepository struct {
    DB *sql.DB
}

func NewRoleRepository() *RoleRepository {
    return &RoleRepository{
        DB: database.PostgresDB,
    }
}

func (r *RoleRepository) GetAll() ([]model.Role, error) {
    query := `SELECT id, name, description, created_at FROM roles`

    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []model.Role
    for rows.Next() {
        var role model.Role
        err := rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt)
        if err != nil {
            return nil, err
        }
        list = append(list, role)
    }
    return list, nil
}