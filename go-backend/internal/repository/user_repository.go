package repository

import (
	"database/sql"
	"go-backend/internal/database"
	"go-backend/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	CountByUsername(username string) (int, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: database.DB,
	}
}

func (r *userRepository) Create(user *models.User) error {
	query := `INSERT INTO users (username, password, name) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, user.Username, user.Password, user.Name)
	return err
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	query := `SELECT username, password, name, token FROM users WHERE username = ?`
	row := r.db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.Username, &user.Password, &user.Name, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(user *models.User) error {
	query := `UPDATE users SET password = ?, name = ?, token = ? WHERE username = ?`
	_, err := r.db.Exec(query, user.Password, user.Name, user.Token, user.Username)
	return err
}

func (r *userRepository) CountByUsername(username string) (int, error) {
	query := `SELECT COUNT(*) FROM users WHERE username = ?`
	row := r.db.QueryRow(query, username)

	var count int
	err := row.Scan(&count)
	return count, err
}