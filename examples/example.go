package main

import (
        "errors"
        "net/http"

        "github.com/gorilla/mux"
        elmahio "github.com/jimiit92/elmah.io-go"
)

// Application entry point
func main() {
        err := elmahio.Setup("Your-API-Key", "LogId")
        elmahio.SetVersion(1.0)
        elmahio.SetSource("TestApp")
        if err != nil {
                panic(err.Error())
        }
        rtr := mux.NewRouter()
        rtr.Handle("/error", elmahio.ElmahHandler(handler)).Methods("GET")
        http.ListenAndServe(":8080", rtr)
}

// Handler for the /error route
// Will throw an error that will be logged on Elmah.io
func handler(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
        err := errors.New("Hello from Go")
        response := http.Response{
                StatusCode: 400,
        }
        return &response, err
}