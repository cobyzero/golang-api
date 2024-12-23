package main

import (
	"github.com/golang-api/src"
	database "github.com/golang-api/src/Database"
)

func main() {
	database.DbConnection()
	src.InitServer()
}
