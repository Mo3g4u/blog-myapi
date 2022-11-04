package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mo3g4u/blog-myapi/api"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("faul to connect DB")
		return
	}
	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	// ListenAndServeの第二引数はサーバーの中で使うルータを指定する
	log.Fatal(http.ListenAndServe(":8080", r))
}
