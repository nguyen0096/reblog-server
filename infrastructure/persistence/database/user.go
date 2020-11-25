package database

import (
	"fmt"
	"reblog-auth/domain/entity"
	"reblog-auth/domain/repository"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserSQL struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func (s *userRepository) GetUserByID(id string) (*entity.User, error) {
	data := []UserSQL{}

	fmt.Println("Test")
	queryString := fmt.Sprintf("SELECT id, first_name, last_name FROM rb_core.user WHERE id = 1")
	err := s.db.Select(&data, queryString)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		ID:        data[0].ID,
		FirstName: data[0].FirstName,
		LastName:  data[0].LastName,
	}
	return user, nil
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}
