package endpoints

import (
	"github.com/gorilla/mux"
	handler "github.com/reemishirsath/leave-module/services/user/pkg/v1/handlers"
	"github.com/reemishirsath/leave-module/services/user/pkg/v1/middleware"
)

//NewUserRoute All Application Routes Are defiend Here
func NewUserRoute(router *mux.Router, handler *handler.UserHandlersImpl) {
	router.HandleFunc("/login", handler.Login).Methods("POST")

	userSubrouter := router.PathPrefix("/user").Subrouter()
	userSubrouter.Use(middleware.AuthMiddleware)
	userSubrouter.HandleFunc("/leaves", handler.ApplyLeave).Methods("POST")
	userSubrouter.HandleFunc("/leaves", handler.GetLeaves).Methods("GET")
	userSubrouter.HandleFunc("/leaves/{status}", handler.GetLeavesByStatus).Methods("GET")
}
