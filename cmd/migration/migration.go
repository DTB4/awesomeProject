package main

import (
	"awesomeProject/configs"
	"awesomeProject/dbconstructor"
	"flag"
	"github.com/DTB4/logger/v2"
	"io/ioutil"
	"strings"
)

func main() {

	var errorsCount int

	migrateUp := flag.Bool("U", false, "migrate Up")
	migrateDown := flag.Bool("D", false, "migrate down")
	forTesting := flag.Bool("T", false, "take config for testing DB (do nothing without direction flag -U or -D)")
	flag.Parse()

	path := "./configs/config.env"
	msgString := "main_database"
	if *forTesting {
		msgString = "test_database"
		path = "./configs/test_config.env"

	}
	configs.NewConfig(path)
	cfg := configs.InitConfig()

	myLogger := logger.NewLogger(cfg.LogsPath)

	db := dbconstructor.NewDB(&cfg.DBConfig, myLogger)

	if *migrateDown {

		myLogger.InfoLog("starting migration Down for ", msgString)
		file, err := ioutil.ReadFile("cmd/migration/001_Down.sql")
		if err != nil {
			myLogger.FatalLog("Error while file opening", err)
		}
		stringSQL := string(file)
		sqlCommands := strings.Split(stringSQL, ";")
		tx, err := db.Begin()
		if err != nil {
			myLogger.FatalLog("Fail to start transaction", err)
		}
		myLogger.InfoLog("Transaction find ", len(sqlCommands), " commands")
		for i, command := range sqlCommands {
			if command != "" {
				_, err := tx.Exec(command)
				if err != nil {
					errorsCount++
					myLogger.ErrorLog("Error in transaction", err)
				}
			}
			myLogger.InfoLog("In process command ", i+1, " from", len(sqlCommands))
		}
		err = tx.Commit()
		if err != nil {
			myLogger.FatalLog("Error during Commit", err)
		}
		if errorsCount != 0 {
			myLogger.InfoLog("SQL commands was executed with ", errorsCount, " errors! Manual DB check recommended")
		} else {
			myLogger.InfoLog("Successfully migrate down ", msgString)
		}

	} else if *migrateUp {

		myLogger.InfoLog("starting migration UP for ", msgString)

		file, err := ioutil.ReadFile("cmd/migration/001_Up.sql")
		if err != nil {
			myLogger.FatalLog("Error while file opening", err)
		}
		stringSQL := string(file)
		sqlCommands := strings.Split(stringSQL, ";")
		tx, err := db.Begin()
		if err != nil {
			myLogger.FatalLog("Fail to start transaction", err)

		}
		myLogger.InfoLog("Transaction find ", len(sqlCommands), " commands")
		for i, command := range sqlCommands {
			if command != "" {
				_, err := tx.Exec(command)
				if err != nil {
					errorsCount++
					myLogger.ErrorLog("Error in transaction", err)
				}
			}
			myLogger.InfoLog("In process command ", i+1, " from", len(sqlCommands))
		}
		err = tx.Commit()
		if err != nil {
			myLogger.FatalLog("Error during Commit", err)
		}
		if errorsCount != 0 {
			myLogger.InfoLog("SQL commands was executed with ", errorsCount, " errors! Manual DB check or Migrate Down recommended")
		} else {
			myLogger.InfoLog("Successfully migrate up ", msgString)
		}

	} else {
		myLogger.InfoLog("do nothing... provide a direction flag or -help to get available flags")
	}

}
