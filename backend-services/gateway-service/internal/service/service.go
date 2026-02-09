package service

import (
	"net/http"
	"gateway-service/internal/proxy"
)


type Service struct {
	AuthURL   string
	ChatURL   string
	JWTSecret string
}

func NewService(AuthURL, ChatURL, JWTSecret string) *Service{
	return &Service{
		AuthURL: AuthURL,
		ChatURL: ChatURL,
		JWTSecret: JWTSecret,
	}
}

func (s *Service)AuthProxy() http.Handler{
	return proxy.NewProxy(s.AuthURL)
}

func (s *Service)ChatProxy() http.Handler{
	return proxy.NewProxy(s.ChatURL)
}