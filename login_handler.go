package mediaserver

import (
    "net/http"
    "fmt"
)


func validLogin(username, password string) bool {
    return username == "test" && password == "test"
}


func (srv* Server) loginHandler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    username := q.Get("username")
    password := q.Get("password")
    if len(username) < 3 || len(password) < 3 {
        fmt.Fprintf(w, `{"login"": "failed"}`)
        return
    }
    
    loginOK := validLogin(username, password)
    if loginOK {
        
        fmt.Fprintf(w, `{"login"": "ok"}`)
        return
    }
    fmt.Fprintf(w, `{"login"": "failed"}`)
}