package controllers

import (
    "encoding/json"
    "net/http"
    "go_application/models"
    "go_application/util"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
    db := util.GetDB()
    items, err := models.GetAllItems(db)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(items)
}
