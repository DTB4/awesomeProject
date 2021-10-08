package midleware

import (
	"awesomeProject/models"
	"github.com/DTB4/logger/v2"
	"net/http"
)

func NewCORSHandler(logger *logger.Logger, cfg *models.Config) *CORSHandler {
	return &CORSHandler{
		logger: logger,
		cfg:    cfg,
	}
}

type CORSHandlerI interface {
	AddCORSHeaders(next http.HandlerFunc) http.HandlerFunc
}

type CORSHandler struct {
	logger *logger.Logger
	cfg    *models.Config
}

func (ch CORSHandler) AddCORSHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", ch.cfg.CorsHandlerConfig)
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
		w.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin, Authorization ,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

		//catch preflight request
		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		} else {
			next(w, req)
		}

	}
}
