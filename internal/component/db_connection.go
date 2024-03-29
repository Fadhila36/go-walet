package component

import (
	"database/sql"
	"fmt"
	"go-walet/internal/config"
	"log"

	_ "github.com/lib/pq"
)

func GetDatabaseConnection(cnf *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable",
		cnf.Database.Host,
		cnf.Database.Port,
		cnf.Database.User,
		cnf.Database.Password,
		cnf.Database.Name,
	)
	connection, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error when open connection %s", err.Error())
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalf("error when open connection %s", err.Error())
	}
	return connection
}
