package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProxy(target string) http.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Printf("Invalid target %s, err: %v", target, err)
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}
