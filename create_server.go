package mediaserver

import (
    "net/http"
)
func createMediaServer(server *http.Server, path string) (*Server, error) {
    if path[len(path) - 1] == '/' {
        path = path[:len(path)-1]
    }
    s := &Server {
        server: server,
        mux: http.NewServeMux(),
        path: path,
        sessions: make(map[string]*session),
        sessionsByUser: make(map[string]*session),
    }
    s.server.Handler = s.mux
    s.mux.HandleFunc("/", s.indexHandler)
    s.mux.Handle("/static/", http.FileServer(http.Dir(path)))
    s.mux.HandleFunc("/user/", s.checkSecure(s.userHandler))
    s.mux.HandleFunc("/shared/", s.checkSecure(s.sharedHandler))
    s.mux.HandleFunc("/login", s.loginHandler)
    return s, nil
}