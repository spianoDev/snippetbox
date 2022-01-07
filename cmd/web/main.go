package main

import (
    "flag"
    "log"
    "net/http"
    "os"
)

// Creating app struct for all app-wide dependencies
type application struct {
    errorLog *log.Logger
    infoLog *log.Logger
}

func main() {
    // Adding a command line flag
    addr := flag.String("addr", ":4000", "HTTP network address")
    // Function below parses the command line flag, reads it and assigns the addr variable
    // The returned value is a pointer to the flag value so it needs to be prefixed with '*'
    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    // adding Lshortfile to the flags will include relevant file name and line number. can also use Llongfile
    errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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
    err := srv.ListenAndServe()
    errorLog.Fatal(err)
}

// You can store the config settings in env variables by using the below function
// addr := os.Getenv("SNIPPETBOX_ADDR")
// Then, when starting the application, you can run the following bash commands to set a preferred address port
// export SNIPPETBOX_ADDR=":4444" (replace the string value with preferred port)
// go run ./cmd/web -addr=$SNIPPETBOX_ADDR