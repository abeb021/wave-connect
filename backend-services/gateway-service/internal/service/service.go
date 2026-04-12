package service

import (
	"gateway-service/internal/proxy"
	"net/http"
)

type Service struct {
	AuthURL    string
	ChatURL    string
	FeedURL    string
	ProfileURL string
	JWTSecret  string
}

func NewService(AuthURL, ChatURL, FeedURL, ProfileURL, JWTSecret string) *Service {
	return &Service{
		AuthURL:    AuthURL,
		ChatURL:    ChatURL,
		FeedURL:    FeedURL,
		ProfileURL: ProfileURL,
		JWTSecret:  JWTSecret,
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

func (s *Service) ProfileProxy() http.Handler {
	return proxy.NewProxy(s.ProfileURL)
}
