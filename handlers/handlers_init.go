package handlers

import (
	"awesomeProject/midleware"
)

type Initializer struct {
	profileHandler *ProfileHandler
	orderHandler   *OrderHandler
	authHandler    *midleware.AuthHandler
}

func NewInitializer(profileHandler *ProfileHandler, orderHandler *OrderHandler, authHandler *midleware.AuthHandler) *Initializer {
	return &Initializer{
		profileHandler: profileHandler,
		orderHandler:   orderHandler,
		authHandler:    authHandler,
	}
}

//func (ah handlersInitializer) NewServeMux (handler *ProfileHandler) *http.ServeMux {
//
//	mux := http.ServeMux{}
//
//	mux := http.NewServeMux()
//
//
//
//	return mux
//}
//
//func (ah handlersInitializer) InitRoutes (mux *http.ServeMux) {
//
//}
