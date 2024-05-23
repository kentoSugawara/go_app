package service

import (
    "net/http"
    "path/filepath"
)

func ServeLoginFile(w http.ResponseWriter, r *http.Request) {
    path, _ := filepath.Abs("./views/login.html")
    http.ServeFile(w, r, path)
}