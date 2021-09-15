package server

import (
	"awesomeProject/dbconstructor"
	"awesomeProject/handlers"
	"awesomeProject/midleware"
	"awesomeProject/models"
	"awesomeProject/parser"
	"awesomeProject/repository"
	"awesomeProject/services"
	"github.com/DTB4/logger/v2"
	"net/http"
)

func Start(cfg *models.Config) {
	myLogger := logger.NewLogger(cfg.LogsPath)

	db := dbconstructor.NewDB(&cfg.DBConfig, myLogger)

	suppliersRepository := repository.NewSupplierRepository(db)
	productRepository := repository.NewProductsRepository(db)
	userRepository := repository.NewUserRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	orderProductsRepository := repository.NewOrderProductsRepository(db)

	orderService := services.NewOrderService(orderRepository, orderProductsRepository)
	userService := services.NewUserService(userRepository)
	tokenService := services.NewTokenService(&cfg.AuthConfig)

	authHandler := midleware.NewAuthHandler(&cfg.AuthConfig, tokenService, myLogger)
	orderHandler := handlers.NewOrderHandler(orderService, myLogger)
	profileHandler := handlers.NewProfileHandler(userService, tokenService, myLogger)

	menuParser := parser.NewMenuParser(&cfg.ParserConfig, myLogger, suppliersRepository, productRepository)
	go menuParser.TimedParsing()

	mux := http.NewServeMux()

	mux.HandleFunc("/registration", profileHandler.CreateNewUser)
	mux.HandleFunc("/profile", authHandler.AccessTokenCheck(profileHandler.ShowUserProfile))
	mux.HandleFunc("/editprofile", authHandler.AccessTokenCheck(profileHandler.EditUserProfile))
	mux.HandleFunc("/refresh", authHandler.RefreshTokenCheck(profileHandler.Refresh))
	mux.HandleFunc("/login", profileHandler.Login)

	mux.HandleFunc("/createorder", authHandler.AccessTokenCheck(orderHandler.Create))
	mux.HandleFunc("/getorder", authHandler.AccessTokenCheck(orderHandler.GetByID))
	mux.HandleFunc("/getmyorders", authHandler.AccessTokenCheck(orderHandler.GetAll))
	mux.HandleFunc("/updateorder", authHandler.AccessTokenCheck(orderHandler.Update))

	myLogger.ErrorLog("Fail to start server", http.ListenAndServe(cfg.ServerPort, mux))

}
