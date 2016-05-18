package mediaserver


import (
    "os"
    "path/filepath"
    "strings"
    "net/http"
)


func (srv* Server) userHandler(w http.ResponseWriter, r *http.Request, s *session) {
    p := r.URL.Path
    if p[0] != '/' {
        p = "/" + p
    }
    requestPath := strings.TrimPrefix(p, "/user/")
    if !strings.HasPrefix(requestPath, s.username) {
        http.NotFound(w, r)
        return
    }
    p = strings.TrimPrefix(requestPath, s.username)
    
    if p == "/" || len(p) == 0 {
        http.ServeFile(w, r, srv.datapath + "view.html")
        return
    }
    
    f, err := os.Open(srv.datapath + p)
    if err != nil {
        http.NotFound(w, r)
        return
    }
    
    if info.IsDir() {
    }
    
    dirs := `"dirs": [ `
    files := `"files": [ `
    
    dirs += " ]"
    files += " ]"
}