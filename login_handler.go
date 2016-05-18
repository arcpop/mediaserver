package mediaserver

import (
    "net/http"
    "fmt"
)


func validLogin(username, password string) (bool, string) {
    return (username == "test" && password == "test"), "test"
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
    
    loginOK, username := validLogin(username, password)
    if loginOK {
        sess := srv.createSession(username)
        http.SetCookie(w, &http.Cookie{
            Name: "SID",
            Value: sess.sessionID,
            Expires: sess.validUntil,
        })
        fmt.Fprintf(w, `{"login": "ok", "username": "` + username + `" }`)
        return
    }
    fmt.Fprintf(w, `{"login": "failed", "reason": "Wrong credentials"}`)
}