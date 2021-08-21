package dbconstructor

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

func NewDB() *sql.DB {

	dbUserName := viper.GetString("DATABASE_USER")
	dbUserPassword := viper.GetString("DATABASE_PASSWORD")
	dbProtocolIpPort := viper.GetString("DATABASE_PROTOCOL_IP_PORT")
	dbName := viper.GetString("DATABASE_NAME")

	dataSourceName := dbUserName + ":" + dbUserPassword + "@" + dbProtocolIpPort + "/" + dbName
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
