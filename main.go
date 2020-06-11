package main

import (
    "log"
    "net/http"
    "encoding/json"
)


type Intro struct {
    Name        string  `json:"name"`
    Kg          string  `json:"kg"`
    LoginName   string  `json:"login_name"`
}

func Root(w http.ResponseWriter, r *http.Request) {
    data := Intro{
        Name: "Seiki Makino",
        Kg: "ONE",
        LoginName: "kino-ma",
    }

    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/", Root)

    log.Fatal(http.ListenAndServe(":8000", nil))
}
