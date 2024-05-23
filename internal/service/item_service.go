package service

import (
    "net/http"
    "text/template"

    "github.com/kentoSugawara/outcome/internal/model"
    "github.com/kentoSugawara/outcome/internal/repository"
)

type ItemService struct {
    Repo *repository.ItemRepository
}

func (svc *ItemService) ListItems(w http.ResponseWriter, r *http.Request) {
    items, err := svc.Repo.GetItems()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles("./views/list.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, items)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}