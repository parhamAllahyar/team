package http

import (
	"admin/configs"
	"admin/internal/core/usecase"
	// "net/http"

	// "github.com/gin-gonic/gin"

)

type HttpHandlers struct {
	AuthUsecase     usecase.AuthUsecase
	AdminUsecase    usecase.AdminUsecase
	PaymentUsecase  usecase.PaymentUsecase
	CustomerUsecase usecase.CustomerUsecase
}

type Server struct {
	HttpHandlers HttpHandlers
	Config       configs.HTTPServer
}

func NewHttpServer(httpHandlers HttpHandlers, config configs.HTTPServer) error {

	srv := Server{
		HttpHandlers: HttpHandlers{
			AuthUsecase:     httpHandlers.AuthUsecase,
			AdminUsecase:    httpHandlers.AdminUsecase,
			PaymentUsecase:  httpHandlers.PaymentUsecase,
			CustomerUsecase: httpHandlers.CustomerUsecase,
		},
		Config: config,
	}

	mux := srv.registerRoutes()

	addr := ":" + config.Port
	err := mux.Run(addr)
	if err != nil {
		return err
	}
	return nil
}

// func (h *Server) runHttpServer(mux *gin.Engine, config configs.HTTPServer) error {
// 	return mux.Run()
// }
