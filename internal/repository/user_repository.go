package repository

import (
	"database/sql"
	"errors"

	"exammple.com/event-booking-api/internal/domain"
	"exammple.com/event-booking-api/pkg/utils"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Signup(user *domain.User) error {
	query := `INSERT INTO users (
	 name, email, password) VALUES (?, ?, ?)`

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Name, user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId

	return err
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	query := `SELECT * FROM users`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	defer rows.Close()

	return users, nil
}

func (r *UserRepository) Login(user *domain.User) error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := r.db.QueryRow(query, user.Email)

	var retrievedPassword string

	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid.")
	}

	isPasswordValid := utils.ValidatePasswordHash(user.Password, retrievedPassword)

	if !isPasswordValid {
		return errors.New("Credentials invalid.")
	}

	return nil
}
