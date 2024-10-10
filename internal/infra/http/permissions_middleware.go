package http

import "net/http"

// PermissionsMiddleware is a middleware that checks for the presence of the correct API key
func PermissionsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") != "Ap1K3y" {
			forbiddenResponse(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
