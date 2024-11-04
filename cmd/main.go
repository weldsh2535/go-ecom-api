package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/weldsh2535/go-rest-api/cmd/api"
	"github.com/weldsh2535/go-rest-api/config"
	"github.com/weldsh2535/go-rest-api/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:              config.Envs.DBUser,
		Passwd:            config.Envs.DBPassword,
		Addr:              config.Envs.DBAddress,
		DBName:            config.Envs.DBName,
		Net:               "tcp",
		AllowOldPasswords: true,
		ParseTime:         true,
	})

	initStorage(db)

	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Success")
}
