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

func SetMark(w http.ResponseWriter, r *http.Request) {
	var req database.MarkFio
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(err)
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

	var stds = []database.ID{}
	var query = fmt.Sprintf("select students.id\n"+
		"from students\n"+
		"where students.title = '%s'", req.Student)

	fmt.Println(query)

	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		var std database.ID
		err := result.Scan(&std.StudentId)
		if err != nil {
			panic(err)
			continue
		}
		stds = append(stds, std)
	}

	var marks = []database.Mark{}
	query = fmt.Sprintf("select * from marks where subject_id = %s and student_id = %d", req.SubjectId, stds[0].StudentId)
	result, err = db.Query(query)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	for result.Next() {
		var mark database.Mark
		err := result.Scan(&mark.Value, &mark.StudentId, &mark.SubjectId)
		if err != nil {
			continue
		}
		marks = append(marks, mark)
	}
	if len(marks) != 0 {
		fmt.Println("Mark updated")
		var insertion = fmt.Sprintf("update marks set value = %s where student_id = %d and subject_id = %s", req.Value, stds[0].StudentId, req.SubjectId)
		_, err = db.Exec(insertion)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Mark inserted")
		var insertion = fmt.Sprintf("insert into marks values (%s, %d, %s)", req.Value, stds[0].StudentId, req.SubjectId)
		_, err = db.Exec(insertion)
		if err != nil {
			panic(err)
		}
	}
}
