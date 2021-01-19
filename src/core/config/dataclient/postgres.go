package dataclient

import (
	"database/sql"
	"fmt"
	"log"
)

type Postgres struct {
	DbConnection DbConnection
}

func (Client Postgres) Client() *sql.DB {

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		Client.DbConnection.Username,
		Client.DbConnection.Password,
		Client.DbConnection.Host,
		Client.DbConnection.Port,
		Client.DbConnection.Database)

	var err error
	var db *sql.DB

	if db, err = sql.Open("postgres", url); err != nil {
		log.Println(fmt.Sprintf("postgres connect error : (%v)", err))
		return nil
	}

	if err = db.Ping(); err != nil {
		log.Println(fmt.Sprintf("postgres ping error : (%v)", err))
		return nil
	}
	return db
}
