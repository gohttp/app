package app

// Route provides fluent API for App
type Route struct {
	app *App
	pat string
}

// Route returns new Route instance for a pattern
func (a *App) Route(pat string) *Route {
	return &Route{a, pat}
}

// Head will register a handler for HEAD requests.
func (r *Route) Head(h interface{}) *Route {
	r.app.Head(r.pat, h)
	return r
}

// Get will register a handler for GET requests.
// It also registers pat for HEAD requests.
func (r *Route) Get(h interface{}) *Route {
	r.app.Get(r.pat, h)
	return r
}

// Post will register a handler for POST requests.
func (r *Route) Post(h interface{}) *Route {
	r.app.Post(r.pat, h)
	return r
}

// Put will register a handler for PUT requests.
func (r *Route) Put(h interface{}) *Route {
	r.app.Put(r.pat, h)
	return r
}

// Del will register a handler for DELETE requests.
func (r *Route) Del(h interface{}) *Route {
	r.app.Del(r.pat, h)
	return r
}

// Options will register a handler for OPTIONS requests.
func (r *Route) Options(h interface{}) *Route {
	r.app.Options(r.pat, h)
	return r
}
