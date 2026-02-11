package service

import (
	"gateway-service/internal/proxy"
	"net/http"
)

type Service struct {
	AuthURL   string
	ChatURL   string
	FeedURL   string
	JWTSecret string
}

func NewService(AuthURL, ChatURL, FeedURL, JWTSecret string) *Service {
	return &Service{
		AuthURL:   AuthURL,
		ChatURL:   ChatURL,
		FeedURL:   FeedURL,
		JWTSecret: JWTSecret,
	}
}

func (s *Service) AuthProxy() http.Handler {
	return proxy.NewProxy(s.AuthURL)
}

func (s *Service) ChatProxy() http.Handler {
	return proxy.NewProxy(s.ChatURL)
}

func (s *Service) FeedProxy() http.Handler {
	return proxy.NewProxy(s.FeedURL)
}

