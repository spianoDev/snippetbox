package main

import (
    "database/sql"
    "flag"
    "log"
    "net/http"
    "os"
    "github.com/joho/godotenv"
// The underscore creates an alias for the package name to a blank identifier so the driver init() function can run
    _ "github.com/go-sql-driver/mysql"
)

// Creating app struct for all app-wide dependencies
type application struct {
    errorLog *log.Logger
    infoLog *log.Logger
}
// function to load/read environment file
func goDotEnvVariable(key string) string {
    err := godotenv.Load(".env")
    if err != nil {
        errorLog.Println("Error loading the .env file...")
    }

    return os.Getenv(key)
}

func main() {
    // retrieve password from .env for sql connection
    sqlPass := gotDotEnvVariable("SQL_PASS")
    infoLog.Println(sqlPass)

    // Adding a command line flag
    addr := flag.String("addr", ":4000", "HTTP network address")
    // Adding flag for mysql dsn
    dsn := flag.String("dsn", "web:" + sqlPass + "@/snippetbox?parseTime=true", "MySQL data source name")
    // Function below parses the command line flag, reads it and assigns the addr variable
    // The returned value is a pointer to the flag value so it needs to be prefixed with '*'
    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    // adding Lshortfile to the flags will include relevant file name and line number. can also use Llongfile
    errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    // open the MySQL database
    db, err := openDB(*dsn)
    if err != nil {
        errorLog.Fatal(err)
    }

    // Close connection pool before the main function exits
    defer db.Close()

    // New instance of application with the dependencies
    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
    }

    srv := &http.Server{
        Addr: *addr,
        ErrorLog: errorLog,
        Handler: app.routes(),
    }

    infoLog.Println("Serving up GO on %s", *addr)
    // now that err variable is declared earlier, we assign (=) instead of declare and assign (:=) the variable
    err = srv.ListenAndServe()
    errorLog.Fatal(err)
}

// new function to return a sql.DB connection pool for given DSN
func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}