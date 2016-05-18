package mediaserver


import (
    "fmt"
    "net/http"
)


func (srv* Server) userHandler(w http.ResponseWriter, r *http.Request, s *session) {
    fmt.Fprintln(w, "Hello world!")
}