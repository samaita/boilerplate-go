package init

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"github.com/samaita/boilerplate-go/config"
	"github.com/samaita/boilerplate-go/pkg/pg"
)

type DBConnection struct {
	MainDB  *sqlx.DB
	Timeout time.Duration
}

// ConnectDB creating new DB connection(s) from config
func ConnectDB(config config.Config) (conn DBConnection) {
	var err error

	pgConf := config.Datastore.Database.Postgres
	pgConn := pg.Postgres{
		DBName:            pgConf.DBName,
		Host:              pgConf.Host,
		MaxConnection:     pgConf.MaxConnection,
		MaxIdleConnection: pgConf.MaxIdleConnection,
		Password:          pgConf.Password,
		Port:              pgConf.Port,
		SSLMode:           pgConf.SSLMode,
		User:              pgConf.User,
	}

	conn.MainDB, err = pgConn.Connect()
	if err != nil {
		log.Fatalln("Err:", err)
	}

	conn.MainDB.SetMaxOpenConns(pgConn.MaxConnection)
	conn.MainDB.SetMaxIdleConns(pgConf.MaxIdleConnection)

	conn.Timeout = pgConf.Timeout
	if err = conn.testConnection(); err != nil {
		log.Fatalln("Err:", err)
	}

	log.Println("Success Connect:", fmt.Sprintf("%s:%s/%s", pgConf.Host, pgConf.Port, pgConf.DBName))
	log.Println(fmt.Sprintf("MaxConn: %d, MaxIdle: %d, Timeout: %v", pgConf.MaxConnection, pgConf.MaxIdleConnection, pgConf.Timeout))
	return
}

// testConnection do ping within predefined timeout, better to be called once on connect
func (conn *DBConnection) testConnection() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), conn.Timeout)
	defer cancel()

	if err = conn.MainDB.PingContext(ctx); err != nil {
		return
	}

	return
}
