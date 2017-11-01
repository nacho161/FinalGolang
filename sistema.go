package main

import "net/http"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB
var err error

func homePage(res http.ResponseWriter, req *http.Request) {
    http.ServeFile(res, req, "formulario.html")
}

func main() {
  // Create an sql.DB and check for errors
   db, err = sql.Open("mysql", "localhost:jhohan12345@/tiempos")
   if err != nil {
       panic(err.Error())
   }
   // sql.DB should be long lived "defer" closes it once this function ends
   defer db.Close()

   // Test the connection to the database
   err = db.Ping()
   if err != nil {
       panic(err.Error())
   }
    http.HandleFunc("/", homePage)
    http.ListenAndServe(":8080", nil)
}

func form(res http.ResponseWriter, req *http.Request) {
    // If method is GET serve an html login page
    if req.Method != "POST" {
        http.ServeFile(res, req, "formulario.html")
        return
    }
}
