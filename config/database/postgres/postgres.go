package postgres

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/ogabrielrodrigues/crud-huncoding/config"
	"github.com/ogabrielrodrigues/crud-huncoding/config/database"
	"github.com/ogabrielrodrigues/crud-huncoding/config/logger"
)

type postgres struct{}

func GetDB() database.DB {
	return &postgres{}
}

func (*postgres) Connect() (*sql.DB, error) {
	conn, err := sql.Open("postgres", config.GetDBConfig().ConnString)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	err = conn.Ping()

	return conn, err
}
