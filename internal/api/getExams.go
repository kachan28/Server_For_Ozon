package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/BinaryArchaism/decanath/internal/config"
	"github.com/BinaryArchaism/decanath/internal/database"
	"log"
	"net/http"
)

func GetExams(w http.ResponseWriter, r *http.Request) {
	var req database.Group
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Fatal(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var exams = []database.Exam{}
	var query = fmt.Sprintf("select subject_id, s.title\n"+
		"from schedules\n"+
		"join subjects s on schedules.subject_id = s.id\n"+
		"where group_id = %d", req.Id)

	exs, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer exs.Close()

	for exs.Next() {
		var std database.Exam
		err := exs.Scan(&std.SubjectId, &std.SubjectTitle)
		if err != nil {
			continue
		}
		exams = append(exams, std)
	}
	fmt.Println("getExams")
	jsonResponse, err := json.Marshal(exams)
	w.Write(jsonResponse)
}
