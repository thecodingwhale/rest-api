// routes.go

package main

func (a *App) initializeRoutes() {
  a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
  a.Router.HandleFunc("/user", a.createUser).Methods("POST")
  a.Router.HandleFunc("/user/{id:[0-9]+}", a.getUser).Methods("GET")
  a.Router.HandleFunc("/user/{id:[0-9]+}", a.updateUser).Methods("PUT")
  a.Router.HandleFunc("/user/{id:[0-9]+}", a.deleteUser).Methods("DELETE")
}
