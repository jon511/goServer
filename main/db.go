package main

import (
	"net/http"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type userData struct {
	ID int `json:"id"`
	FirstName string `json:"first-name"`
	LastName string `json:"last-name"`
	Info string `json:"info"`
}

func handleDbRequest(w http.ResponseWriter, r *http.Request){

	db, err := sql.Open("mysql", "root:jobb7116@tcp(127.0.0.1:3306)/test_schema")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	db.Query("INSERT INTO test_table (id, first_name, last_name, info) VALUES (25, 'jon', 'osborne', '{\"key\":\"value\"}')")

	results1 , err := db.Query("SELECT * FROM test_table")
	if err != nil {
		fmt.Println(err)
	}

	var sendResult []string
	var newResult string
	for results1.Next() {
		var tag userData
		err = results1.Scan(&tag.ID, &tag.FirstName, &tag.LastName, &tag.Info)
		if err != nil {
			panic(err.Error())
		}

		log.Printf("%s - %s - %s", tag.FirstName, tag.LastName, tag.Info)

		sendResult = append(sendResult, fmt.Sprintf("%d - %s %s - %s", tag.ID, tag.FirstName, tag.LastName, tag.Info))

		newResult = tag.Info

	}
	fmt.Println(sendResult)
	fmt.Fprintf(w, "%s", newResult)

}