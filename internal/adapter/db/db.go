package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"

	_ "github.com/lib/pq"
)

const (
	retry int = 10
	sleep     = 5 * time.Second
)

func Init(driverName string, dsn string) *sql.DB {
	conn := connect(driverName, dsn)
	if conn == nil {
		log.Panicln("unable to connect to db")
	}
	return conn
}

func MakeDsn(conf *config.Db) string {
	return fmt.Sprintf(
		"dbname='%s' "+
			"host='%s' "+
			"port='%d' "+
			"user='%s' "+
			"password='%s' "+
			"sslmode=disable",
		conf.Name, conf.Host, conf.Port, conf.User, conf.Password,
	)
}

func connect(driverName string, dsn string) *sql.DB {
	var (
		db  *sql.DB
		err error
	)

	for i := 0; i < retry; i++ {
		db, err = open(driverName, dsn)

		if err == nil {
			return db
		}

		log.Printf("db is not ready trying again in %d seconds...\n", sleep/time.Second)

		time.Sleep(sleep)
	}

	return nil
}

func open(driverName string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
