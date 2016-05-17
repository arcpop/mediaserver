package mediaserver

import (
    "sync"
    "net/http"
)

type Server struct {
    server *http.Server
    mux *http.ServeMux
    path string
    sessions map[string]*session
    sessionsByUser map[string]*session
    sessionsLock sync.RWMutex
}

func New(server *http.Server, path string) (*Server, error) {
    return createMediaServer(server, path)
}

func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error {
    return srv.server.ListenAndServeTLS(certFile, keyFile)
}
