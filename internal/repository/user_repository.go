package repository

import (
	"database/sql"
	"ecommerce-grpc-api/internal/models"
)

type UserRepository interface {
	InsertOne(data *models.User) error
	GetOne(id string) (*models.User, error)
	UpdateOne(id string, data *models.User) error
	DeleteOne(id string) error
	GetAll() ([]*models.User, error)
}

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}
func (r *UserRepositoryImpl) InsertOne(data *models.User) error {
	query := "Insert INTO users (username,email,password) Values (?,?,?)"
	_, err := r.Db.Exec(query, data.Username, data.Email, data.Password_hash)
	if err != nil {
		return err
	}
	return nil

}

func (r *UserRepositoryImpl) GetOne(id string) (*models.User, error) {
	var user models.User
	query := "SELECT FROM users WHERE id=?"
	err := r.Db.QueryRow(query, id).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepositoryImpl) UpdateOne(id string, data *models.User) error {

	query := "UPDATE users SET email=? ,passord=? WHERE id=?"
	_, err := r.Db.Exec(query, data.Email, data.Password_hash, id)
	if err != nil {
		return err
	}
	return nil

}

func (r *UserRepositoryImpl) DeleteOne(id string) error {
	query := "DELETE FROM users WHERE id=?"
	_, err := r.Db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	query := "SELECT FORM users "

	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, err
	}
	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password_hash, &user.Email, &user.Roles)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)

	}
	return users, nil

}
