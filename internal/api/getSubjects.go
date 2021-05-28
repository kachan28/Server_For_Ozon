package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/BinaryArchaism/decanath/internal/config"
	"github.com/BinaryArchaism/decanath/internal/database"
	"net/http"
)

func GetSubjects(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var stds = []database.Subject{}
	result, err := db.Query("select * from subjects")
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		var std database.Subject
		err := result.Scan(&std.Id, &std.Title, &std.TimeAmount)
		if err != nil {
			continue
		}
		stds = append(stds, std)
	}

	fmt.Println("getSubjects")
	jsonResponse, err := json.Marshal(stds)
	w.Write(jsonResponse)
}
