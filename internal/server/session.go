package server

import (
	"net/http"
	"time"
)

type Session struct {
	name   string
	expiry time.Time
}

func (s *Session) Expired() bool {
	return s.expiry.Before(time.Now())
}

func (s *Server) GetSessionFromRequest(req *http.Request) (*Session, error) {

	cookie, err := req.Cookie("x-session")
	if err != nil {
		return nil, err
	}
	token, err := s.Queries.GetToken(req.Context(), cookie.Value)
	if err != nil {
		return nil, err
	}
	user, err := s.Queries.GetUser(req.Context(), token.UserID)
	if err != nil {
		return nil, err
	}

	return &Session{
		name:   user.Name,
		expiry: time.UnixMilli(token.ExpiresAt),
	}, nil

}
