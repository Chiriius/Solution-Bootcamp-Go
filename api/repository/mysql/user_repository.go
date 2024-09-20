package mysql

import (
	"bootcamp_api/api/entities"
	errorss "bootcamp_api/api/utils/errors"
	"errors"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	GetUser(id string) (entities.User, error)
	AddUser(user entities.User) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
}

type MySqlUserRepository struct {
	db     *sqlx.DB
	logger logrus.FieldLogger
}

func NewMySQLUserRepository(db *sqlx.DB, logger logrus.FieldLogger) *MySqlUserRepository {
	return &MySqlUserRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *MySqlUserRepository) GetUser(id string) (entities.User, error) {
	var user entities.User
	err := repo.db.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {
		repo.logger.Errorln(`Layer:user_repository Method:GetUser,`, err)
		if err == sql.ErrNoRows {
			return user, errorss.ErrorUserNotFound
		}
		return user, err
	}
	repo.logger.Infoln(`Layer:user_repository, Method:GetUser, User:`, user)
	return user, nil
}

func (repo *MySqlUserRepository) AddUser(user entities.User) (entities.User, error) {
	query := "INSERT INTO users (id, password, age, information, parents, email , name ) VALUES( ?, ?, ?, ?, ?, ?, ?)"
	_, err := repo.db.Exec(query, user.ID, user.Password, user.Age, user.Information, user.Parents, user.Email, user.Name)
	if err != nil {
		repo.logger.Errorln(`Layer:user_repository Method:AddUser,`, err)
		if err == sql.ErrNoRows {
			return user, errors.New("Codigo:505 Message:Server error")
		}
		return user, err
	}
	repo.logger.Infoln(`Layer:user_repository, Method:AddUser, User:`, user)
	return user, err
}

func (repo *MySqlUserRepository) UpdateUser(user entities.User) (entities.User, error) {
	query := (`UPDATE users SET password = ?, age = ?, information = ?, parents = ?, email = ?, name = ? WHERE id = ?`)
	_, err := repo.db.Exec(query, user.Password, user.Age, user.Information, user.Parents, user.Email, user.Name, user.ID)

	if err != nil {
		repo.logger.Errorln(`Layer:user_repository Method:UpdateUser,`, err)
		if err == sql.ErrNoRows {
			return user, errors.New("Codigo:505 Message:Server error")
		}
		return user, err
	}
	repo.logger.Infoln(`Layer:user_repository, Method:UpdateUser, User:`, user)
	return user, err
}
