package config

import (
	"api/src/core/config/dataclient"
	"database/sql"
)

type DataClientMock struct {
	PostgresMockClient PostgresMockClient
}


type PostgresMockClient struct{}

func (PostgresMockClient PostgresMockClient) Client() *sql.DB {
	db, _ := sql.Open("postgres", "root:nimda@/test")

	return db
}

func (DataClient DataClientMock) PostgresClient() dataclient.ISqlClient {
	return DataClient.PostgresMockClient
}
