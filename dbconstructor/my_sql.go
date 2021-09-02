package dbconstructor

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func NewDB() *sql.DB {

	dbUserName := os.Getenv("DATABASE_USER")
	dbUserPassword := os.Getenv("DATABASE_PASSWORD")
	dbProtocolIpPort := os.Getenv("DATABASE_PROTOCOL_IP_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	dataSourceName := dbUserName + ":" + dbUserPassword + "@" + dbProtocolIpPort + "/" + dbName + "?parseTime=True"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
