package dataclient
import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DbConnection struct {
	Username string
	Password string
	Host string
	Port string
	Database string
}

type ISqlClient interface {
	Client() *sql.DB
}

type IDataClient interface {
	PostgresClient() ISqlClient
}

type DataClient struct {
	Postgres ISqlClient
}

func (DataClient DataClient) PostgresClient() ISqlClient {
	return DataClient.Postgres
}

