package app

import "github.com/justinas/alice"
import "github.com/bmizerany/pat"
import "net/http"

// Application.
type App struct {
	*pat.PatternServeMux
	chain alice.Chain
}

// New application.
func New() *App {
	return &App{
		PatternServeMux: pat.New(),
		chain:           alice.New(),
	}
}

// Use the given middleware.
func (a *App) Use(mw ...alice.Constructor) {
	a.chain = a.chain.Append(mw...)
}

// Head will register a pattern with a handler for HEAD requests.
func (a *App) Head(pat string, h interface{}) {
	a.PatternServeMux.Head(pat, handler(h))
}

// Get will register a pattern with a handler for GET requests.
// It also registers pat for HEAD requests.
func (a *App) Get(pat string, h interface{}) {
	a.PatternServeMux.Get(pat, handler(h))
}

// Post will register a pattern with a handler for POST requests.
func (a *App) Post(pat string, h interface{}) {
	a.PatternServeMux.Post(pat, handler(h))
}

// Put will register a pattern with a handler for PUT requests.
func (a *App) Put(pat string, h interface{}) {
	a.PatternServeMux.Put(pat, handler(h))
}

// Del will register a pattern with a handler for DELETE requests.
func (a *App) Del(pat string, h interface{}) {
	a.PatternServeMux.Del(pat, handler(h))
}

// Options will register a pattern with a handler for OPTIONS requests.
func (a *App) Options(pat string, h interface{}) {
	a.PatternServeMux.Options(pat, handler(h))
}

// Listen on `addr`.
func (a *App) Listen(addr string) error {
	handler := a.chain.Then(a.PatternServeMux)
	http.Handle("/", handler)
	return http.ListenAndServe(addr, nil)
}

// ServeHTTP implements http.Handler
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := a.chain.Then(a.PatternServeMux)
	handler.ServeHTTP(w, r)
}

// coerce handler into an http.Handler.
func handler(h interface{}) http.Handler {
	switch h.(type) {
	case func(w http.ResponseWriter, r *http.Request):
		return http.HandlerFunc(h.(func(w http.ResponseWriter, r *http.Request)))
	case http.Handler:
		return h.(http.Handler)
	default:
		panic("invalid handler")
	}
}
