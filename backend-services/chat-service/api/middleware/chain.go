package middleware 

import "net/http"

type Middleware func (http.Handler) http.Handler

func Chain(middlewairs ...Middleware) Middleware{
	return func (next http.Handler) http.Handler{
		for i:=len(middlewairs) - 1; i>=0; i--{
			next=middlewairs[i](next)
		}
		return next
	}
}

