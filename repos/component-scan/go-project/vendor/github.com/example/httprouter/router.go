// Package httprouter is a lightweight HTTP request router.
package httprouter

import (
	"fmt"
	"net/http"
	"strings"
)

type Handler func(http.ResponseWriter, *http.Request, map[string]string)

type Router struct {
	routes map[string]Handler
}

func New() *Router {
	return &Router{routes: make(map[string]Handler)}
}

func (r *Router) GET(path string, h Handler)  { r.routes["GET:"+path] = h }
func (r *Router) POST(path string, h Handler) { r.routes["POST:"+path] = h }

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := fmt.Sprintf("%s:%s", req.Method, strings.TrimRight(req.URL.Path, "/"))
	if h, ok := r.routes[key]; ok {
		h(w, req, nil)
		return
	}
	http.NotFound(w, req)
}
