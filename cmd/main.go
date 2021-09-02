package main

import (
	"awesomeProject/dbconstructor"
	"awesomeProject/handlers"
	"awesomeProject/repository"
	"awesomeProject/services"
	"fmt"
	"github.com/DTB4/logger/v2"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	err := getConfig()
	if err != nil {
		log.Fatalf("Fail to initialise configs: %s", err.Error())
	}
	myLogger := logger.NewLogger(os.Getenv("LOGS_DIRECTORY_PATH"))

	db := dbconstructor.NewDB()
	suppliersRepository := repository.NewSupplierRepository(db)
	productRepository := repository.NewProductsRepository(db)
	userRepository := repository.NewUserRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	userService := services.NewUserService(orderRepository, userRepository)
	tokenService := services.NewTokenService()
	profileHandler := handlers.NewProfileHandler(userService, tokenService, myLogger)

	menuParser := services.NewMenuParser(myLogger, suppliersRepository, productRepository)
	go menuParser.TimedParsing(10)

	http.HandleFunc("/getall", profileHandler.GetAll)
	http.HandleFunc("/registration", profileHandler.CreateNewUser)
	http.HandleFunc("/profile", profileHandler.TokenCheck(profileHandler.ShowUserProfile))
	http.HandleFunc("/editprofile", profileHandler.TokenCheck(profileHandler.EditUserProfile))
	http.HandleFunc("/login", profileHandler.Login)

	log.Fatal(http.ListenAndServe(os.Getenv("HTTP_PORT"), nil))
}

func getConfig() error {
	b, err := os.ReadFile("configs/config.env")
	if err != nil {
		return err
	}
	s := string(b)
	lines := strings.Split(s, "\n")
	for i := range lines {
		if lines[i] == "" {
			break
		}
		params := strings.Split(lines[i], "=")
		err = os.Setenv(params[0], params[1])
		if err != nil {
			fmt.Println("failed to set env parameter")
		}
		//fmt.Println(params[1])
	}

	return nil
}
