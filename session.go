package mediaserver

import (
    "io"
    "time"
    "net/http"
    "crypto/rand"
	"encoding/base64"
	"log"
)

type session struct {
    username string
    token string
    validUntil time.Time
    sessionID string
}

var (
    //DefaultSessionTimeout is the timeout for any session. If a user sends no requests during this time, 
    //they have to login again.
    DefaultSessionTimeout = time.Hour * 24 * 7
)


func (srv *Server) checkSession(r *http.Request) (s *session, expired bool) {
    cookie, err := r.Cookie("SID")
    if err != nil {
        return
    }
    srv.sessionsLock.RLock()
    s = srv.sessions[cookie.Value]
    srv.sessionsLock.RUnlock()
    if s == nil {
        return 
    }
    
    //Check if session is still valid.
    if s.validUntil.Before(time.Now()) {
        srv.sessionsLock.Lock()
        delete(srv.sessions, cookie.Value)
        delete(srv.sessionsByUser, s.username)
        srv.sessionsLock.Unlock()
        return nil, true
    }
    s.validUntil = time.Now().Add(DefaultSessionTimeout)
    return
}

func (srv *Server) createSession(username string) (*session) {
    srv.sessionsLock.Lock()
    defer srv.sessionsLock.Unlock()
    
    s, ok := srv.sessionsByUser[username]
    if ok {
        log.Print("Found old user session, maybe a bug?")
    }
    data := make([]byte, 128)
    _, err := io.ReadFull(rand.Reader, data)
    if err != nil {
        return nil
    }
    sessionID := base64.RawStdEncoding.EncodeToString(data);
    for ; srv.sessions[sessionID] != nil; sessionID = base64.RawStdEncoding.EncodeToString(data) {
        _, err := io.ReadFull(rand.Reader, data)
        if err != nil {
            return nil
        }
    }
    s = &session{
        username: username,
        token: "",
        validUntil: time.Now().Add(DefaultSessionTimeout),
        sessionID: sessionID,
    }
    srv.sessionsByUser[username] = s
    srv.sessions[sessionID] = s
    return s
}