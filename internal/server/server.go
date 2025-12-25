package server

import (
	"fmt"
	"net/http"

	"github.com/nturbo1/reverse-proxy/internal/configs"
	"github.com/nturbo1/reverse-proxy/internal/middleware"
	"github.com/nturbo1/reverse-proxy/internal/routing"
	"github.com/nturbo1/reverse-proxy/internal/log"
)

func NewServer(
	appConfigs *configs.AppConfigs, env *configs.Environment,
) *http.Server {

	mux := http.NewServeMux()
	serverHandler := NewServerHandler(mux)
	log.Debug("Setting up the routes...")
	routing.SetUpRouteHandlers(appConfigs, env, mux)

	return &http.Server{
		Addr:           fmt.Sprintf(":%d", appConfigs.Server.Port),
		Handler:        serverHandler,
		ReadTimeout:    appConfigs.Server.Timeout,
		WriteTimeout:   appConfigs.Server.Timeout,
		MaxHeaderBytes: 1 << 20,
	}
}

func NewServerHandler(mux *http.ServeMux) http.Handler {

	return middleware.PrependMiddlewareChain(
		mux,
		middleware.RateLimitMiddleware,
		middleware.LogMiddleware,
		middleware.AuthMiddleware,
	)
}
