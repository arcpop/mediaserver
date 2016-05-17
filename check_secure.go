package mediaserver

import (
    "net/http"
)
func (srv *Server) checkSecure(handlerFunc func (http.ResponseWriter, *http.Request, *session)) func (http.ResponseWriter, *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        s, _ := srv.checkSession(r)
        if s == nil {
            http.Redirect(w, r, "/static/login.html", http.StatusFound)
            return
        }
        handlerFunc(w, r, s)
    }
}