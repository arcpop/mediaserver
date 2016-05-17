package mediaserver

import (
    "net/http"
)


func (srv* Server) indexHandler(w http.ResponseWriter, r *http.Request) {
    s, _ := srv.checkSession(r)
    if s == nil {
        http.Redirect(w, r, "/static/login.html", http.StatusFound)
        return
    }
    http.Redirect(w, r, "/user/", http.StatusFound)
}