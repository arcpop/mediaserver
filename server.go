package mediaserver

import (
    "sync"
    "net/http"
)

type Server struct {
    server *http.Server
    mux *http.ServeMux
    htmlpath string
    datapath string
    sessions map[string]*session
    sessionsByUser map[string]*session
    sessionsLock sync.RWMutex
}

func New(server *http.Server, htmlpath, datapath string) (*Server, error) {
    return createMediaServer(server, htmlpath, datapath)
}

func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error {
    return srv.server.ListenAndServeTLS(certFile, keyFile)
}
