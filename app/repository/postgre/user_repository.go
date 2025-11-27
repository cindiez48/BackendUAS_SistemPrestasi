package postgre

import (
    "database/sql"
    "backenduas_sistemprestasi/database"
    model "backenduas_sistemprestasi/app/models/postgre"
)

type UserRepository struct {
    DB *sql.DB
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        DB: database.PostgresDB,
    }
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
    query := `
        SELECT id, username, email, password_hash, full_name, role_id,
               is_active, created_at, updated_at
        FROM users WHERE email=$1`

    row := r.DB.QueryRow(query, email)

    var user model.User
    err := row.Scan(
        &user.ID, &user.Username, &user.Email, &user.PasswordHash,
        &user.FullName, &user.RoleID, &user.IsActive,
        &user.CreatedAt, &user.UpdatedAt,
    )

    return &user, err
}
