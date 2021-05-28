package database

import (
	"database/sql"
	"fmt"
	"github.com/BinaryArchaism/decanath/internal/config"
	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
