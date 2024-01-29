package models

import (
    "database/sql"
    "log"
)

type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// GetAllItems retrieves all items from the database.
func GetAllItems(db *sql.DB) ([]Item, error) {
    items := []Item{}
    rows, err := db.Query("SELECT * FROM items;")
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

// GetItem retrieves a single item by its ID.
func GetItem(db *sql.DB, id int) (*Item, error) {
    var i Item
    row := db.QueryRow("SELECT * FROM items WHERE id = $1", id)
    if err := row.Scan(&i.ID, &i.Name); err != nil {
        return nil, err
    }
    return &i, nil
}

// Insert a new item in the database
func CreateItem(db *sql.DB, name string) (int64, error) {
    var id int64
    log.Printf("Executing SQL: INSERT INTO items (name) VALUES ($1) RETURNING id with name=%s", name) // Log the statement with RETURNING clause
    err := db.QueryRow("INSERT INTO items (name) VALUES ($1) RETURNING id;", name).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}


// UpdateItem updates an existing item's name.
func UpdateItem(db *sql.DB, id int, name string) error {
    _, err := db.Exec("UPDATE items SET name = $1 WHERE id = $2;", name, id)
    return err
}

// DeleteItem removes an item from the database.
func DeleteItem(db *sql.DB, id int) error {
    _, err := db.Exec("DELETE FROM items WHERE id = $1;", id)
    return err
}
