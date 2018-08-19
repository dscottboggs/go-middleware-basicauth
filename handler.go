package auth

import (
	"net/http"
)

// BasicAuth is the middleware function. it passes the request through if the
// request is authenticated. If the request is unauthenticated, it writes the
// 401 (forbidden) result.
func BasicAuth(
	w http.ResponseWriter,
	r *http.Request,
) (
	http.ResponseWriter, *http.Request,
) {
	if IsUnauthenticatedEndpoint(r.URL.RawQuery) {
		return w, r
	}
	user, pass, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusForbidden)
	}
	if u := User(user); !u.IsAuthenticatedBy(pass) {
		w.WriteHeader(http.StatusForbidden)
	}
	return w, r
}
