package endpoints

import (
	"github.com/gorilla/mux"
	handler "github.com/reemishirsath/leave-module/services/leave/pkg/v1/handlers"
	"github.com/reemishirsath/leave-module/services/leave/pkg/v1/middleware"
)

//NewUserRoute All Application Routes Are defiend Here
func NewUserRoute(router *mux.Router, handler *handler.UserHandlersImpl) {
	router.HandleFunc("/login", handler.Login).Methods("POST")

	leaveSubrouter := router.PathPrefix("/leaves").Subrouter()
	leaveSubrouter.Use(middleware.AuthMiddleware)
	leaveSubrouter.HandleFunc("/{id}", handler.ProcessLeave).Methods("POST")
	leaveSubrouter.HandleFunc("/user/{id}", handler.GetLeavesByUser).Methods("GET")
	leaveSubrouter.HandleFunc("/user/{id}/{status}", handler.GetLeavesByStatus).Methods("GET")
}
