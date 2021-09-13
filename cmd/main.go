package main

import (
	"awesomeProject/dbconstructor"
	"awesomeProject/handlers"
	"awesomeProject/midleware"
	"awesomeProject/parser"
	"awesomeProject/repository"
	"awesomeProject/services"
	"fmt"
	"github.com/DTB4/logger/v2"
	"log"
	"net/http"
	"os"
	"strconv"
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
	orderProductsRepository := repository.NewOrderProductsRepository(db)
	orderService := services.NewOrderService(orderRepository, orderProductsRepository)
	userService := services.NewUserService(userRepository)
	tokenService := services.NewTokenService()
	authHandler := midleware.NewAuthHandler(tokenService, myLogger)
	orderHandler := handlers.NewOrderHandler(orderService, myLogger)
	profileHandler := handlers.NewProfileHandler(userService, tokenService, myLogger)
	url := os.Getenv("URL_FOR_API_PARSER")
	format := os.Getenv("FORMAT_STRING_FOR_API_URL")
	parserDelay, _ := strconv.Atoi(os.Getenv("PARSER_DELAY_SECONDS"))
	menuParser := parser.NewMenuParser(url, format, myLogger, suppliersRepository, productRepository, parserDelay)
	go menuParser.TimedParsing()

	http.HandleFunc("/registration", profileHandler.CreateNewUser)
	http.HandleFunc("/profile", authHandler.AccessTokenCheck(profileHandler.ShowUserProfile))
	http.HandleFunc("/editprofile", authHandler.AccessTokenCheck(profileHandler.EditUserProfile))
	http.HandleFunc("/refresh", authHandler.RefreshTokenCheck(profileHandler.Refresh))
	http.HandleFunc("/login", profileHandler.Login)

	http.HandleFunc("/createorder", authHandler.AccessTokenCheck(orderHandler.Create))
	http.HandleFunc("/getorder", authHandler.AccessTokenCheck(orderHandler.GetByID))
	http.HandleFunc("/getmyorders", authHandler.AccessTokenCheck(orderHandler.GetAll))
	http.HandleFunc("/updateorder", authHandler.AccessTokenCheck(orderHandler.Update))

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
