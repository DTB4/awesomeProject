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

func Start(cfg *models.Config) *http.Server {
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
	supplierService := services.NewSupplierService(suppliersRepository)
	productService := services.NewProductService(productRepository)

	corsHandler := midleware.NewCORSHandler(myLogger, cfg)
	authHandler := midleware.NewAuthHandler(tokenService, myLogger)
	orderHandler := handlers.NewOrderHandler(orderService, myLogger)
	userHandler := handlers.NewUserHandler(userService, tokenService, myLogger)
	supplierHandler := handlers.NewSupplierHandler(supplierService, myLogger)
	productHandler := handlers.NewProductHandler(productService, myLogger)

	menuParser := parser.NewMenuParser(&cfg.ParserConfig, myLogger, suppliersRepository, productRepository)
	go menuParser.TimedParsing()

	mux := http.NewServeMux()

	mux.HandleFunc("/registration", corsHandler.AddCORSHeaders(userHandler.CreateNewUser))
	mux.HandleFunc("/profile", corsHandler.AddCORSHeaders(authHandler.AccessTokenCheck(userHandler.ShowUserProfile)))
	mux.HandleFunc("/editprofile", authHandler.AccessTokenCheck(userHandler.EditUserProfile))
	mux.HandleFunc("/refresh", corsHandler.AddCORSHeaders(authHandler.RefreshTokenCheck(userHandler.Refresh)))
	mux.HandleFunc("/logout", corsHandler.AddCORSHeaders(authHandler.AccessTokenCheck(userHandler.Logout)))
	mux.HandleFunc("/login", corsHandler.AddCORSHeaders(userHandler.Login))

	mux.HandleFunc("/createorder", authHandler.AccessTokenCheck(orderHandler.Create))
	mux.HandleFunc("/getorder", authHandler.AccessTokenCheck(orderHandler.GetByID))
	mux.HandleFunc("/getmyorders", authHandler.AccessTokenCheck(orderHandler.GetAll))
	mux.HandleFunc("/updateorder", authHandler.AccessTokenCheck(orderHandler.Update))

	mux.HandleFunc("/supplier", supplierHandler.GetSupplierByID)
	mux.HandleFunc("/suppliers", supplierHandler.GetAllSuppliers)
	mux.HandleFunc("/supplierstype", supplierHandler.GetSuppliersByType)
	mux.HandleFunc("/supplierstime", supplierHandler.GetSuppliersByTime)

	mux.HandleFunc("/product", productHandler.GetProductDyID)
	mux.HandleFunc("/products", productHandler.GetAll)
	mux.HandleFunc("/productsbytype", productHandler.GetAllByType)
	mux.HandleFunc("/productsbysupplier", productHandler.GetAllBySupplierID)

	srv := http.Server{
		Addr:    cfg.ServerPort,
		Handler: mux,
	}

	return &srv
}
