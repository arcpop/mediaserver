package mediaserver


import (
    "crypto/sha256"
    "net/http"
	"time"
    "strings"
	"os"
	"io"
)


func (srv *Server) uploadHandler(w http.ResponseWriter, r *http.Request, s *session) {
    err := r.ParseMultipartForm(100 << 30) //20MB
    if err != nil {
        http.Error(w, "500 Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "500 Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    defer file.Close()
    csum := sha256.Sum256([]byte(time.Now().String() + header.Filename))
    parts := strings.Split(header.Filename, ".")
    extension := parts[len(parts) - 1]
    if len(extension) != 3 {
        extension = "mp4"
    }
    filename := string(csum[:]) + "." + extension
    fd, err := os.OpenFile("/tmp/" + filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        http.Error(w, "500 Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    defer fd.Close()
    _, err = io.Copy(fd, file)
    if err != nil {
        http.Error(w, "500 Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    pathname := srv.datapath + "/vod/" + string(csum[:])
    err = os.Mkdir(pathname, 0777)
    if err != nil {
        http.Error(w, "500 Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    currentDir := r.FormValue("dir")
    movieName := r.FormValue("name")
    if currentDir == "" || movieName == "" {
        http.Error(w, "400 Bad request", http.StatusBadRequest)
        return
    }
    
    
    go convertVideoToStream(filename, pathname, srv.datapath + currentDir, movieName + ".json")
    
    w.WriteHeader(http.StatusOK)
    return
}
