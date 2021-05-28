package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/BinaryArchaism/decanath/internal/config"
	"github.com/BinaryArchaism/decanath/internal/database"
	"net/http"
)

func GetSchedules(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var stds = []database.Schedule{}
	result, err := db.Query("select schedules.id, schedules.cabinet, schedules.lecturer_id, schedules.subject_id, schedules.group_id, schedules.date, groups.cath_id from schedules, groups where schedules.group_id = groups.id")
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		var std database.Schedule
		err := result.Scan(&std.Id, &std.Cabinet, &std.LecturerId, &std.SubjectId, &std.GroupId, &std.Date, &std.CathId)
		if err != nil {
			continue
		}
		stds = append(stds, std)
	}
	fmt.Println("getSchedules")
	jsonResponse, err := json.Marshal(stds)
	w.Write(jsonResponse)
}
