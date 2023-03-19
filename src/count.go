package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
  "net/http"
)

func handleIncrement(w http.ResponseWriter, r *http.Request, stmtUpdate *sql.Stmt) {
  fmt.Println("got /increment request")
  _, err := stmtUpdate.Exec() // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
}

func handleClear(w http.ResponseWriter, r *http.Request, stmtClear *sql.Stmt) {
  fmt.Println("got /clear request")
  _, err := stmtClear.Exec()
  if err != nil {
    panic(err.Error())
  }
}

func main() {
	db, err := sql.Open("mysql", "root:wQUMsqxELx@/my_database")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for increment count
	stmtUpdate, err := db.Prepare("UPDATE test set tal = tal + 1 where id = 1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtUpdate.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for clear count
	stmtClear, err := db.Prepare("UPDATE test set tal = 1 where id = 1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtClear.Close() // Close the statement when we leave main() / the program terminates

  http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {handleIncrement(w, r, stmtUpdate)})
  http.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request) {handleClear(w, r, stmtClear)})
  err = http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
