package app

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/franela/go-supertest"
	. "github.com/franela/goblin"
)

func TestApp(t *testing.T) {
	g := Goblin(t)

	a := New()
	a.Head("/", sayhi)
	a.Get("/", sayhi)
	a.Post("/", echo)
	a.Put("/", echo)
	a.Del("/", echo)
	a.Options("/", sayhi)

	testApp(a, g, "app")
}

const hi = "hi"

func sayhi(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(hi))
}

func echo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	b, _ := ioutil.ReadAll(req.Body)
	w.Write(b)
}

func testApp(a *App, g *G, name string) {
	server := httptest.NewServer(a)
	defer server.Close()

	describe := g.Describe
	it := g.It

	describe(name, func() {
		it("HEAD should respond 200", func(done Done) {
			NewRequest(server.URL).
				Head("/").
				Expect(200, done)
		})

		it("GET should say hi with status 200", func(done Done) {
			NewRequest(server.URL).
				Get("/").
				Expect(200, hi, done)
		})

		it("POST should respond ok with status 200", func(done Done) {
			NewRequest(server.URL).
				Post("/").
				Send("ok").
				Expect(200, "ok", done)
		})

		it("PUT should respond ok with status 200", func(done Done) {
			NewRequest(server.URL).
				Put("/").
				Send("ok").
				Expect(200, "ok", done)
		})

		it("DELETE should respond ok with status 200", func(done Done) {
			NewRequest(server.URL).
				Delete("/").
				Send("ok").
				Expect(200, "ok", done)
		})

		it("OPTIONS should say hi with status 200", func(done Done) {
			NewRequest(server.URL).
				Options("/").
				Send(hi).
				Expect(200, hi, done)
		})
	})
}
