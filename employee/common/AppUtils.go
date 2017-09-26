package common

import (
	"net/http"
)

type (
	appError struct {
		Error   string `json:"error"`
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	configuration struct {
		Server, MysqlUser, MysqlPwd, MysqlDatabase string
	}
)

func displayAppError(w http.ResponseWriter, handlerError error, message string, code int) {

}
