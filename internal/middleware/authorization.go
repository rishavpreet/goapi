package middleware

import (
	"errors"
	"net/http"

	"github.com/rishavpreet/goapi/api"
	"github.com/rishavpreet/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var ErrorUnAuthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		var err error

		if username == "" || token == "" {
			log.Error(ErrorUnAuthorized)
			api.RequestErrorHandler(w, ErrorUnAuthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(ErrorUnAuthorized)
			api.RequestErrorHandler(w, ErrorUnAuthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
