package controllers

import (
	"net/http"

	database "github.com/golang-api/src/Database"
	"github.com/golang-api/src/Database/Dao"
)

func CreateTask(response *http.ResponseWriter, request *http.Request) {

	Dao.Query.Task.Create(database.DB, Dao.Task)

}
