// controllers.go

package main

import (
  "database/sql"
  "encoding/json"
  "net/http"
  "strconv"
  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
)

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
  count, _ := strconv.Atoi(r.FormValue("count"))
  start, _ := strconv.Atoi(r.FormValue("start"))

  if count > 10 || count < 1 {
    count = 10
  }
  if start < 0 {
    start = 0
  }

  products, err := getUsers(a.DB, start, count)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, products)
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
  var u User
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&u); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }
  defer r.Body.Close()

  if err := u.createUser(a.DB); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusCreated, u)
}

func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid user ID")
    return
  }

  u := User{ID: id}
  if err := u.getUser(a.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      respondWithError(w, http.StatusNotFound, "User not found")
    default:
      respondWithError(w, http.StatusInternalServerError, err.Error())
    }
    return
  }

  respondWithJSON(w, http.StatusOK, u)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid user ID")
    return
  }

  var u User
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&u); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
    return
  }
  defer r.Body.Close()
  u.ID = id

  if err := u.updateUser(a.DB); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, u)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid User ID")
    return
  }

  u := User{ID: id}
  if err := u.deleteUser(a.DB); err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
