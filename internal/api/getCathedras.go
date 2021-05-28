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

func GetCathedras(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var stds = []database.Cathedra{}
	result, err := db.Query("select * from cathedras")
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		var std database.Cathedra
		err := result.Scan(&std.Id, &std.Title, &std.Number)
		if err != nil {
			continue
		}
		stds = append(stds, std)
	}
	sort.Slice(stds, func(i, j int) bool {
		return stds[i].Number < stds[j].Number
	})
	fmt.Println("getCathedras")
	jsonResponse, err := json.Marshal(stds)
	w.Write(jsonResponse)
}
