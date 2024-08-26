package mysql

import (
	"bootcamp_api/api/entities"
	errorss "bootcamp_api/api/utils/errors"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUserById(id string) (entities.User, error)
	AddUser(user entities.User) error
}

type MySqlUserRepository struct {
	db *sqlx.DB
}

func NewMySQLUserRepository(db *sqlx.DB) *MySqlUserRepository {
	return &MySqlUserRepository{db: db}
}

func (repo *MySqlUserRepository) GetUserById(id string) (entities.User, error) {
	var user entities.User
	err := repo.db.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errorss.ErrorUserNotFound //Logica cambiada
		}
		return user, err
	}
	return user, nil
}

func (repo *MySqlUserRepository) AddUser(user entities.User) error {
	query := "INSERT INTO users (id, password, age, information, parents, email , name ) VALUES( ?, ?, ?, ?, ?, ?, ?)"
	_, err := repo.db.Exec(query, user.ID, user.Password, user.Age, user.Information, user.Parents, user.Email, user.Name)

	return err
} //Ya devuelve el identificador en el endpoint
