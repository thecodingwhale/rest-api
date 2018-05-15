// main.go

package main

import (
  "os"
  "log"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  a := App{}
  a.Initialize(
    os.Getenv("APP_DB_USERNAME"),
    os.Getenv("APP_DB_PASSWORD"),
    os.Getenv("APP_DB_NAME"))

  a.Run(":8080")
}
