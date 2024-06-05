package service

import (
    "net/http"
    "encoding/json"

    "github.com/kentoSugawara/outcome/internal/repository"
)

type ItemService struct {
    Repo *repository.ItemRepository
}

func (s *ItemService) ListItems(w http.ResponseWriter, r *http.Request) {
    items, err := s.Repo.GetItems()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(items)
}