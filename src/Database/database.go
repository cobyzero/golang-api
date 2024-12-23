package database

import (
	"log"

	"github.com/golang-api/src/Database/Dao"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func DbConnection() {
	var dst = "host=aws-0-us-west-1.pooler.supabase.com user=postgres.xqpbqcikgjzzaleftrac password=Seek15wish16@ port=5432 dbname=postgres"

	DB, err := gorm.Open(postgres.Open(dst), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		Dao.Use(DB)
		log.Println("Database connected")

	}
}
