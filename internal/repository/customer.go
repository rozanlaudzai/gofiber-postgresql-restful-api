package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/domain"
)

type customerRepository struct {
	db *goqu.Database
}

func NewCustomer(dbConn *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: goqu.New("default", dbConn),
	}
}

func (cr *customerRepository) FindAll(ctx context.Context) (result []domain.Customer, err error) {
	dataset := cr.db.From("customer").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (cr *customerRepository) FindById(ctx context.Context, id string) (result domain.Customer, err error) {
	dataset := cr.db.From("customer").
		Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (cr *customerRepository) Save(ctx context.Context, c *domain.Customer) error {
	executor := cr.db.Insert("customer").Rows(c).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr *customerRepository) Update(ctx context.Context, c *domain.Customer) error {
	executor := cr.db.Update("customer").
		Where(goqu.C("id").Eq(c.ID)).
		Set(c).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr *customerRepository) Delete(ctx context.Context, id string) error {
	executor := cr.db.Update("customer").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
