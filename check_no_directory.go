package mediaserver

import (
    "net/http"
)

type dircheck struct {
    handler http.Handler
}

func checkNoDirectory(handler http.Handler) http.Handler {
    return &dircheck{ handler: handler }
}

func isPath(s string) bool {
    last := s[len(s) - 1]
    return last == '/' || last == '\\'
}

func (d *dircheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if isPath(r.URL.Path) || len(r.URL.Path) == 0 {
        http.NotFound(w, r)
        return
    }
    d.handler.ServeHTTP(w, r)
}