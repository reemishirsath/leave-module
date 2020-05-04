package handlers

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/reemishirsath/leave-module/logs"

	"net/http"

	"github.com/reemishirsath/leave-module/models"
	service "github.com/reemishirsath/leave-module/services/user/pkg/v1/services"
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

// ApplyLeave handler Function
func (userHandlersImpl UserHandlersImpl) ApplyLeave(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	leave := models.Leave{}
	log.Logger(ctx).Info("in request")
	err := ReadInput(req.Body, &leave)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	resp, err := userHandlersImpl.userSvc.ApplyLeave(ctx, leave)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	WriteOKResponse(w, resp)
}

// GetLeaves handler Function
func (userHandlersImpl UserHandlersImpl) GetLeaves(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, erri := parseToken(req)
	if erri != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}
	log.Logger(ctx).Info("in request")
	resp, err := userHandlersImpl.userSvc.GetLeaves(ctx, id)
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
	status := vars["status"]
	id, erri := parseToken(req)
	if erri != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}
	log.Logger(ctx).Info("in request")
	resp, err := userHandlersImpl.userSvc.GetLeavesByStatus(ctx, id, status)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	WriteOKResponse(w, resp)
}

//parseToken Helper Function
func parseToken(req *http.Request) (userID int, err error) {
	authorizationHeader := req.Header.Get("authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("95a31f74-a1cd-4321-8a9d-bdb0735e445a"), nil
			})
			if error != nil {
				return 0, error
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				role := claims["roleID"].(float64)
				roleID := int(role)
				return roleID, nil
			}
		}
	}
	return
}
