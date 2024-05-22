package main

import (
    "fmt"
    "net/http"
)

func main() {
    // HTTPリクエストハンドラを設定
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<p>Hello, World!</p>")
    })

    // サーバを8080ポートで起動
    fmt.Println("Server is running on http://localhost:8080/")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
