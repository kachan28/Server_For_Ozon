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

func GetGroups(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var stds = []database.Group{}
	result, err := db.Query("select groups.id, groups.number, cathedras.number from groups, cathedras where groups.cath_id = cathedras.id")
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		var std database.Group
		err := result.Scan(&std.Id, &std.Number, &std.CathId)
		if err != nil {
			continue
		}
		stds = append(stds, std)
	}
	sort.Slice(stds, func(i, j int) bool {
		return stds[i].Number < stds[j].Number
	})
	fmt.Println("getGroups")
	jsonResponse, err := json.Marshal(stds)
	w.Write(jsonResponse)
}
