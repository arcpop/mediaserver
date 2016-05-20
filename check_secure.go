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

type secureHandlerDummy struct {
    inner http.Handler
    server *Server
}

func (srv *Server) secureHandler(handler http.Handler) http.Handler {
    return &secureHandlerDummy{ inner:handler, server: srv }
}

func (shd *secureHandlerDummy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s, _ := shd.server.checkSession(r)
    if s == nil {
        http.Redirect(w, r, "/static/login.html", http.StatusFound)
        return
    }
    shd.inner.ServeHTTP(w, r)
}