package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rozanlaudzai/gofiber-postgresql-restful-api/internal/config"
)

func GetDatabase(config *config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v timezone=%v dbname=%v sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.TimeZone,
		config.Name,
	)
	return sql.Open("postgres", dsn)
}
