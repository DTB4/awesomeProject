package dbconstructor

import (
	"awesomeProject/models"
	"database/sql"
	"github.com/DTB4/logger/v2"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(cfg *models.DBConfig, logger *logger.Logger) *sql.DB {

	dbUserName := cfg.DatabaseUser
	dbUserPassword := cfg.DatabasePassword
	dbProtocolIpPort := cfg.DatabaseProtocolIPAndPort
	dbName := cfg.DatabaseName

	dataSourceName := dbUserName + ":" + dbUserPassword + "@" + dbProtocolIpPort + "/" + dbName + "?parseTime=True"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.FatalLog("fail to establish a db connection", err)
	}
	return db
}
