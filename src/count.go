package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
  "net/http"
  "os"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
  fmt.Println("got /health request")
  fmt.Fprintln(w, "{\"status\":\"ok\"}")
}

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
  dbhost := os.Getenv("dbhost")
  dbname := os.Getenv("dbname")
  dbuser := os.Getenv("dbuser")
  dbpassword := os.Getenv("dbpassword")
  db_dsn := "" + dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":3306)/" + dbname + ""
  fmt.Println("Connection to db on DSN://" + db_dsn )
  db, err := sql.Open("mysql", db_dsn)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
  db.Ping()

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

  http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {handleHealth(w, r)})
  http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {handleIncrement(w, r, stmtUpdate)})
  http.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request) {handleClear(w, r, stmtClear)})
  err = http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
