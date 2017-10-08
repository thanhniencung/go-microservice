package controller

import (
	"backend/employee/common"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	JoinDate  time.Time `json:"joinDate"`
	TeamId    string    `json:"teamId"`
	Avatar    string    `json:"avatar"`
}

func AddNewEmployee(w http.ResponseWriter, r *http.Request) {
	log.Println("[addNewEmployee]")

	firstName := r.PostFormValue("firstName")
	lastName := r.PostFormValue("lastName")
	teamId := r.PostFormValue("teamId")
	avatar := r.PostFormValue("avatar")

	employee := Employee {
		FirstName: firstName,
		LastName:  lastName,
		JoinDate:  time.Now(),
		TeamId:    teamId,
		Avatar:    avatar,
	}

	stmt, err := common.GetDatabaseConnection().Prepare("INSERT INTO employee(firstName,lastName,joinDate,teamId,avatar) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(employee.FirstName, employee.LastName, employee.JoinDate, employee.TeamId, employee.Avatar)
	if err != nil {
		log.Fatal(err)
		defer stmt.Close()
	} else {
		log.Println(res.LastInsertId)
		employeeJson, _ := json.Marshal(employee)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(employeeJson)
	}
	common.GetDatabaseConnection().Close()
}
