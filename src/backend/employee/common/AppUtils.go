package common

import (
	
	"encoding/json"
	"log"
	"net/http"
	"os"
"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}

	errorResource struct {
		Data appError `json:"data"`
	}

	configuration struct {
		Server, MysqlUser, MysqlPwd, MysqlDatabase string
	}
)

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

func InitConfig() {
	file, err := os.Open("/Users/kiennguyen/Desktop/hr-management/src/backend/employee/config/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[logAppConfig]: %s\n", err)
	} else {
		log.Println("AppConfig loaded." + AppConfig.MysqlUser)
	}
}

func GetDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:likeyou1234@/hr-management")
	if err != nil {
		log.Fatalf("[connectDb]: %s\n", err)
	}
	return db
}
