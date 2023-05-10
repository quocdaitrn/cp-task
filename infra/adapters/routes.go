package adapters

import (
	"github.com/quocdaitrn/golang-kit/auth"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	"github.com/quocdaitrn/cp-task/app/transport/api/handler"
	"github.com/quocdaitrn/cp-task/domain/service"
	"github.com/quocdaitrn/cp-task/infra/config"
)

func ProvideRoutes(taskSvc service.TaskService, logger log.Logger, cfg config.Config, authClient auth.AuthenticateClient) RestAPIHandler {
	r := mux.NewRouter()
	handler.MakeAppHandler(r, logger, cfg)

	v1 := r.PathPrefix("/v1").Subrouter()
	handler.MakeTaskHTTPHandler(v1, taskSvc, logger, authClient)

	return setupCORSMiddleware(r)
}

func setupCORSMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		if r.Method == http.MethodOptions {
			// Note: cache CORS for Chrome
			w.WriteHeader(http.StatusOK)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
