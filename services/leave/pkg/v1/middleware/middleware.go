package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/reemishirsath/leave-module/services/leave/pkg/v1/handlers"

	"github.com/dgrijalva/jwt-go"

	log "github.com/reemishirsath/leave-module/logs"
)

//LoggingMiddleware log request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req = req.WithContext(log.WithRqID(req.Context()))
		w.Header().Add("Content-Type", "application/json")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, req)
	})
}

//AuthMiddleware **
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
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
					handlers.WriteHTTPError(w, http.StatusInternalServerError)
					return
				}
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.Header().Set("Access-Control-Allow-Credentials", "true")
					w.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS, PUT, DELETE")
					w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,x-requested-with, XMLHttpRequest, Access-Control-Allow-Methods")
					role := claims["roleID"].(float64)
					roleID := int(role)
					log.Logger(ctx).Info("\nRoleID:", roleID)
					next.ServeHTTP(w, req)
				} else {
					handlers.WriteCustomHTTPError(w, http.StatusBadRequest, "Invalid Token")
					return
				}
			}
		} else {
			handlers.WriteCustomHTTPError(w, http.StatusUnauthorized, "Authorization Header Required")
			return
		}
	})
}
