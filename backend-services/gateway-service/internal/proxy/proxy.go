package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func NewProxy(authURL, chatURL string) http.Handler {
	authTarget, _ := url.Parse(authURL)
	chatTarget, _ := url.Parse(chatURL)

	authProxy := httputil.NewSingleHostReverseProxy(authTarget)
	chatProxy := httputil.NewSingleHostReverseProxy(chatTarget)
	

	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/api/auth"):
			authProxy.ServeHTTP(w, r)
		case strings.HasPrefix(r.URL.Path, "/api/message"):
			chatProxy.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}