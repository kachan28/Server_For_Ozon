package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/BinaryArchaism/decanath/internal/config"
	"github.com/BinaryArchaism/decanath/internal/database"
	"net/http"
	"strconv"
)

func GetStatements(w http.ResponseWriter, r *http.Request) {
	groupId, err := strconv.Atoi(r.URL.Query()["group"][0])
	subjectId, err := strconv.Atoi(r.URL.Query()["subject"][0])

	fmt.Println(groupId)
	fmt.Println(subjectId)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var stdsWithoutMarks = []database.Statement{}
	var query = fmt.Sprintf("select c.number, l.fio, subjects.title, s.date, s2.title, 0\n"+
		"FROM subjects\nJOIN schedules s on subjects.id = s.subject_id\n"+
		"JOIN lecturer l on s.lecturer_id = l.id\n"+
		"JOIN groups g on s.group_id = g.id\n"+
		"JOIN students s2 on g.id = s2.group_id\n"+
		"JOIN cathedras c on g.cath_id = c.id\n"+
		"WHERE subject_id = %d and s.group_id = %d", subjectId, groupId)

	studentsWithoutMarks, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer studentsWithoutMarks.Close()

	for studentsWithoutMarks.Next() {
		var std database.Statement
		err := studentsWithoutMarks.Scan(&std.Cath, &std.Fio, &std.SubjectName, &std.Date, &std.StudentsList, &std.MarksList)
		if err != nil {
			continue
		}
		stdsWithoutMarks = append(stdsWithoutMarks, std)
	}

	var stdsWithMarks = []database.Statement{}
	query = fmt.Sprintf("select 0, 0, 0, 0, students.title, marks.value "+
		"from marks, students where (students.group_id = %d and student_id = students.id) and subject_id = %d", groupId, subjectId)

	studentsWithMarks, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer studentsWithMarks.Close()

	for studentsWithMarks.Next() {
		var std database.Statement
		err := studentsWithMarks.Scan(&std.Cath, &std.Fio, &std.SubjectName, &std.Date, &std.StudentsList, &std.MarksList)
		if err != nil {
			continue
		}
		stdsWithMarks = append(stdsWithMarks, std)
	}

	for k, i := range stdsWithoutMarks {
		for p, j := range stdsWithMarks {
			if i.StudentsList == j.StudentsList {
				stdsWithoutMarks[k].MarksList = stdsWithMarks[p].MarksList
			}
		}
	}

	fmt.Println(stdsWithoutMarks)
	fmt.Println("getStatements")
	jsonResponse, err := json.Marshal(stdsWithoutMarks)
	w.Write(jsonResponse)
}
