package main

import (
	"awesomeProject/dbconstructor"
	"awesomeProject/handlers"
	"awesomeProject/repository"
	"awesomeProject/services"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Fail to initialise configs: %s", err.Error())
	}

	db := dbconstructor.NewDB()
	userRepository := repository.NewUserRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	userService := services.NewUserService(orderRepository, userRepository)
	tokenService := services.NewTokenService()
	profileHandler := handlers.NewProfileHandler(userService, tokenService)

	http.HandleFunc("/getall", profileHandler.GetAll)
	http.HandleFunc("/registration", profileHandler.CreateNewUser)
	http.HandleFunc("/profile", profileHandler.TokenCheck(profileHandler.ShowUserProfile))
	http.HandleFunc("/editprofile", profileHandler.TokenCheck(profileHandler.EditUserProfile))
	http.HandleFunc("/login", profileHandler.Login)

	log.Fatal(http.ListenAndServe(viper.GetString("HTTP_PORT"), nil))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
