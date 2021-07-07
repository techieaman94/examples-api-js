package main

import (
    "fmt"
    "net/http"
)

func greetings(w http.ResponseWriter, req *http.Request) {

    fmt.Println("GET params were:", req.URL.Query())

    message := req.URL.Query()["message"]
    // message := req.URL.Query().Get("message")
    if len(message) > 0 {
        fmt.Fprintf(w, "<h1>%s</h1>", message[0])
    } else {
        fmt.Fprintf(w, "<h1>Hello World !!!</h1>")
    }

}

func main() {
    http.HandleFunc("/", greetings)
    http.ListenAndServe(":5000", nil)
}
