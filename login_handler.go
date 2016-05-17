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
    if len(username) < 1 {
        fmt.Fprintf(w, `{"login": "failed", "reason": "No username specified"}`)
        return
    }
    if len(password) < 1 {
        fmt.Fprintf(w, `{"login": "failed", "reason": "No password specified"}`)
        return
    }
    
    loginOK := validLogin(username, password)
    if loginOK {
        
        fmt.Fprintf(w, `{"login": "ok"}`)
        return
    }
    fmt.Fprintf(w, `{"login": "failed", "reason": "Wrong credentials"}`)
}