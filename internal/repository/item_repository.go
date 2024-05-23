package repository

import (
    "database/sql"
    "log"

    "github.com/kentoSugawara/outcome/internal/model"
)

type ItemRepository struct {
    DB *sql.DB
}

func (repo *ItemRepository) GetItems() ([]model.Item, error) {
    rows, err := repo.DB.Query("SELECT title, url, tagbody, body FROM items")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []model.Item
    for rows.Next() {
        var item model.Item
        err := rows.Scan(&item.Title, &item.URL, &item.TagBody, &item.Body)
        if err != nil {
            log.Fatal(err)
        }
        items = append(items, item)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return items, nil
}