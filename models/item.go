package models

import (
    "database/sql"
)

type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func GetAllItems(db *sql.DB) ([]Item, error) {
    items := []Item{}
    rows, err := db.Query("SELECT id, name FROM items")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var i Item
        if err := rows.Scan(&i.ID, &i.Name); err != nil {
            return nil, err
        }
        items = append(items, i)
    }
    return items, nil
}
