package postgre

import (
    "database/sql"
    "backenduas_sistemprestasi/database"
    model "backenduas_sistemprestasi/app/models/postgre"
)

type PermissionRepository struct {
    DB *sql.DB
}

func NewPermissionRepository() *PermissionRepository {
    return &PermissionRepository{
        DB: database.PostgresDB,
    }
}

func (r *PermissionRepository) GetAll() ([]model.Permission, error) {
    query := `SELECT id, name, resource, action, description FROM permissions`

    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []model.Permission
    for rows.Next() {
        var p model.Permission
        err := rows.Scan(&p.ID, &p.Name, &p.Resource, &p.Action, &p.Description)
        if err != nil {
            return nil, err
        }
        list = append(list, p)
    }
    return list, nil
}