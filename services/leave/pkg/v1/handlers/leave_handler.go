package handlers

import (
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/reemishirsath/leave-module/logs"

	"net/http"

	"github.com/reemishirsath/leave-module/models"
	service "github.com/reemishirsath/leave-module/services/leave/pkg/v1/services"
)

//UserHandlersImpl for handler Functions
type UserHandlersImpl struct {
	userSvc service.UsersService
}

//NewUserHandlerImpl inits dependancies for graphQL and Handlers
func NewUserHandlerImpl(userService service.UsersService) *UserHandlersImpl {
	return &UserHandlersImpl{userSvc: userService}
}

var httpErr models.HTTPErr

// Login Handler
func (userHandlersImpl UserHandlersImpl) Login(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	email, password, _ := req.BasicAuth()
	user := models.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	log.Logger(ctx).Info("in request")

	resp, err := userHandlersImpl.userSvc.Login(ctx, user)
	if err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	WriteOKResponse(w, resp)
}

// ProcessLeave handler Function
func (userHandlersImpl UserHandlersImpl) ProcessLeave(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	vars := mux.Vars(req)
	lid := vars["id"]
	id, _ := strconv.Atoi(lid)
	leave := models.Leave{}
	log.Logger(ctx).Info("in request")
	err := ReadInput(req.Body, &leave)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	resp, err := userHandlersImpl.userSvc.ProcessLeave(ctx, id, leave)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	WriteOKResponse(w, resp)
}

// GetLeavesByUser handler Function
func (userHandlersImpl UserHandlersImpl) GetLeavesByUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	uid := vars["id"]
	id, _ := strconv.Atoi(uid)
	log.Logger(ctx).Info("in request")
	resp, err := userHandlersImpl.userSvc.GetLeavesByUser(ctx, id)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	WriteOKResponse(w, resp)
}

// GetLeavesByStatus handler Function
func (userHandlersImpl UserHandlersImpl) GetLeavesByStatus(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	lid := vars["id"]
	id, _ := strconv.Atoi(lid)
	status := vars["status"]
	log.Logger(ctx).Info("in request")
	resp, err := userHandlersImpl.userSvc.GetLeavesByStatus(ctx, id, status)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	WriteOKResponse(w, resp)
}
