package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/BinaryArchaism/decanath/internal/config"
	"github.com/BinaryArchaism/decanath/internal/database"
	"net/http"
	"sort"
)

func GetLecturers(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var stds = []database.Lecturer{}
	result, err := db.Query("select * from lecturer")
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		var std database.Lecturer
		err := result.Scan(&std.Id, &std.Fio)
		if err != nil {
			continue
		}
		stds = append(stds, std)
	}
	sort.Slice(stds, func(i, j int) bool {
		return stds[i].Fio < stds[j].Fio
	})
	fmt.Println("getLecturers")
	jsonResponse, err := json.Marshal(stds)
	w.Write(jsonResponse)
}
