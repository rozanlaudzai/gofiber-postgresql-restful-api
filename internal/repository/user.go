package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/domain"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(dbConn *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", dbConn),
	}
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (result domain.User, err error) {
	dataset := ur.db.From("user").Where(goqu.C("email").Eq(email))
	_, err = dataset.ScanStruct(&result)
	return
}
