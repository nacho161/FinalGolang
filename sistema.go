package main

import "net/http"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB
var err error

func main() {
  // Create an sql.DB and check for errors
   db, err = sql.Open("mysql", "bc6d4b66ee9885:4b3997db@tcp(ca-cdbr-azure-central-a.cloudapp.net:3306)/tiempos")
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
    http.HandleFunc("/", form)
    http.ListenAndServe(":8080", nil)
}

func form(res http.ResponseWriter, req *http.Request) {
    // If method is GET serve an html login page
    if req.Method != "POST" {
        http.ServeFile(res, req, "formulario.html")
        return
    }

    fecha := req.FormValue("fecha")
    actividad := req.FormValue("actividad")
    horainicio := req.FormValue("horainicio")
    horafinal := req.FormValue("horafinal")
    empleado := req.FormValue("empleado")

    var databasefecha  string
    var databaseactividad  string
    var databasehorainicio  string
    var databasehorafinal  string
    var databaseempleado  string

    err := db.QueryRow("INSERT INTO Registro(IDActividad, HoraInicio, HoraFinal, IDEmpleado, Fecha) VALUES (?,?,?,?,?)", IDRegistro).Scan(&databaseactividad, &databasehorainicio, &databasehorafinal,&databaseempleado,&databasefecha)
        // If not then redirect to the login page
        if err != nil {
            http.Redirect(res, req, "/login", 301)
            return
        }

        switch {
    // Username is available
    case err == sql.ErrNoRows:
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            http.Error(res, "Server error, unable to create your account.", 500)
            return
        }

        res.Write([]byte("User created!"))
        return
    case err != nil:
        http.Error(res, "Server error, unable to create your account.", 500)
        return
    default:
        http.Redirect(res, req, "/", 301)


}
