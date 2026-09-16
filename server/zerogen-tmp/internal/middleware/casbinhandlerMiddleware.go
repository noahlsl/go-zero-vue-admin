// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package middleware

import "net/http"

type CasbinHandlerMiddleware struct {
}

func NewCasbinHandlerMiddleware() *CasbinHandlerMiddleware {
	return &CasbinHandlerMiddleware{}
}

func (m *CasbinHandlerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
