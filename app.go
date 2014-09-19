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

// Listen on `addr`.
func (a *App) Listen(addr string) error {
	handler := a.chain.Then(a)
	http.Handle("/", handler)
	return http.ListenAndServe(addr, nil)
}
