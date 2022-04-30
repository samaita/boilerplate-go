package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	DBName            string
	Host              string
	MaxConnection     int
	MaxIdleConnection int
	Password          string
	Port              string
	SSLMode           string
	User              string
}

// Connect create a new SQLX connection
func (pg *Postgres) Connect() (db *sqlx.DB, err error) {
	return sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", pg.Host, pg.Port, pg.User, pg.Password, pg.DBName, pg.SSLMode))
}
