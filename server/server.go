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
	tokenRepository := repository.NewTokenRepository(db)

	orderService := services.NewOrderService(orderRepository, orderProductsRepository)
	userService := services.NewUserService(userRepository)
	tokenService := services.NewTokenService(&cfg.AuthConfig, tokenRepository)

	authHandler := midleware.NewAuthHandler(tokenService, myLogger)
	orderHandler := handlers.NewOrderHandler(orderService, myLogger)
	userHandler := handlers.NewUserHandler(userService, tokenService, myLogger)

	menuParser := parser.NewMenuParser(&cfg.ParserConfig, myLogger, suppliersRepository, productRepository)
	go menuParser.TimedParsing()

	mux := http.NewServeMux()

	mux.HandleFunc("/registration", userHandler.CreateNewUser)
	mux.HandleFunc("/profile", authHandler.AccessTokenCheck(userHandler.ShowUserProfile))
	mux.HandleFunc("/editprofile", authHandler.AccessTokenCheck(userHandler.EditUserProfile))
	mux.HandleFunc("/refresh", authHandler.RefreshTokenCheck(userHandler.Refresh))
	mux.HandleFunc("/login", userHandler.Login)

	mux.HandleFunc("/createorder", authHandler.AccessTokenCheck(orderHandler.Create))
	mux.HandleFunc("/getorder", authHandler.AccessTokenCheck(orderHandler.GetByID))
	mux.HandleFunc("/getmyorders", authHandler.AccessTokenCheck(orderHandler.GetAll))
	mux.HandleFunc("/updateorder", authHandler.AccessTokenCheck(orderHandler.Update))

	myLogger.ErrorLog("Fail to start server", http.ListenAndServe(cfg.ServerPort, mux))

}
