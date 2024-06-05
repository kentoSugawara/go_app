package main

import (
    "database/sql"
    "log"
    "net/http"

    "github.com/kentoSugawara/outcome/internal/service"
    "github.com/kentoSugawara/outcome/internal/repository"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "gopher:password@tcp(db:3306)/local_go?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    repo := &repository.ItemRepository{DB: db}
    svc := &service.ItemService{Repo: repo}

    http.HandleFunc("/list", svc.ListItems)

    fs := http.FileServer(http.Dir("./app/build"))
    http.Handle("/", fs)

    log.Println("Listening on :8080...")
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}