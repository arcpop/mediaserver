package mediaserver

import (
    "net/http"
)
func createMediaServer(server *http.Server, htmlpath, datapath string) (*Server, error) {
    if htmlpath[len(htmlpath) - 1] == '/' {
        htmlpath = htmlpath[:len(htmlpath)-1]
    }
    if datapath[len(datapath) - 1] == '/' {
        datapath = datapath[:len(datapath)-1]
    }
    
    s := &Server {
        server: server,
        mux: http.NewServeMux(),
        htmlpath: htmlpath,
        datapath: datapath,
        sessions: make(map[string]*session),
        sessionsByUser: make(map[string]*session),
    }
    s.server.Handler = s.mux
    s.mux.HandleFunc("/", s.indexHandler)
    s.mux.Handle("/static/", checkNoDirectory(http.FileServer(http.Dir(htmlpath))))
    //s.mux.HandleFunc("/user/", s.checkSecure(s.userHandler))
    //s.mux.HandleFunc("/shared/", s.checkSecure(s.sharedHandler))
    s.mux.HandleFunc("/login", s.loginHandler)
    s.mux.Handle("/vod/", s.secureHandler(http.FileServer(http.Dir(s.datapath))))
    s.mux.HandleFunc("/upload", s.checkSecure(s.uploadHandler))
    return s, nil
}