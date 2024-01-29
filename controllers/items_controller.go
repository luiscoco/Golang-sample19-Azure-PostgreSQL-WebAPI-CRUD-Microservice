package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "go_application/models"
    "go_application/util"
    "github.com/gorilla/mux"
)

// GetItems handles the request to get all items.
func GetItems(w http.ResponseWriter, r *http.Request) {
    db := util.GetDB()
    items, err := models.GetAllItems(db)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(items)
}

// GetItem handles the request to get a single item by its ID.
func GetItem(w http.ResponseWriter, r *http.Request) {
    db := util.GetDB()
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    item, err := models.GetItem(db, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(item)
}

// CreateItem handles the request to create a new item.
func CreateItem(w http.ResponseWriter, r *http.Request) {
    db := util.GetDB()
    var item models.Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := models.CreateItem(db, item.Name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    item.ID = int(id)
    json.NewEncoder(w).Encode(item)
}

// UpdateItem handles the request to update an existing item.
func UpdateItem(w http.ResponseWriter, r *http.Request) {
    db := util.GetDB()
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var item models.Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = models.UpdateItem(db, id, item.Name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(item)
}

// DeleteItem handles the request to delete an item.
func DeleteItem(w http.ResponseWriter, r *http.Request) {
    db := util.GetDB()
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = models.DeleteItem(db, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

