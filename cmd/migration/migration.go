package main

import (
	"awesomeProject/configs"
	"awesomeProject/dbconstructor"
	"flag"
	"fmt"
	"github.com/DTB4/logger/v2"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("starting an migration")
	configs.NewConfig(false)
	cfg := configs.InitConfig()

	myLogger := logger.NewLogger(cfg.LogsPath)

	db := dbconstructor.NewDB(&cfg.DBConfig, myLogger)

	migrateUp := flag.Bool("U", false, "migrate Up")
	migrateDown := flag.Bool("D", false, "migrate down")

	flag.Parse()

	if *migrateDown {

		fmt.Println("migrate Down")
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
		for _, command := range sqlCommands {
			if command != "" {
				_, err := tx.Exec(command)
				if err != nil {
					myLogger.FatalLog("Error in transaction", err)
				}
			}
		}
		err = tx.Commit()
		if err != nil {
			myLogger.FatalLog("Error during Commit", err)
		}
		myLogger.InfoLog("Database successfully migrate Down")

	} else if *migrateUp {

		fmt.Println("migrate UP")

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
		for _, command := range sqlCommands {
			if command != "" {
				_, err := tx.Exec(command)
				if err != nil {
					myLogger.ErrorLog("Error in transaction", err)
				}
			}
		}
		err = tx.Commit()
		if err != nil {
			myLogger.FatalLog("Error during Commit", err)
		}
		myLogger.InfoLog("Database successfully migrate up")

	} else {
		fmt.Println("do nothing... provide a flag or -help to get available flags")
	}

}
