package database

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
    "log"
    "fmt"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DbConfig struct {
		Server string `yaml:"Server"`
		MysqlUser string `yaml:"MysqlUser"`
		MysqlPwd string `yaml:"MysqlPwd"`
		MysqlDatabase string `yaml:"MysqlDatabase"`
}

func Init() bool {
	dbConfig, error := getDbConfig()
	if (error != nil) {
		log.Println(error)	
		return false
	}
	connectDatabase(dbConfig)
	return true
}

func connectDatabase(dbConfig *DbConfig) *sql.DB {
	stringConection := fmt.Sprintf("%s:%s@/hr-management", dbConfig.MysqlUser, dbConfig.MysqlPwd);
	log.Println(stringConection)	
	db, err := sql.Open("mysql", stringConection)
	if err != nil {
		log.Fatalf("[connectDb]: %s\n", err)
		return nil
	}
	return db
}

func getDbConfig() (*DbConfig, error) {
	var db DbConfig
    yamlFile, err := ioutil.ReadFile("/Users/kiennguyen/golang/hr-management/src/backend/employee/conf.yaml")
    if err != nil {
       return nil, err
    }

    err = yaml.Unmarshal(yamlFile, &db)
    if err != nil {
    	return nil, err
    }

    return &db, nil
}