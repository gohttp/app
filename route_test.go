package app

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/franela/go-supertest"
	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)

	describe := g.Describe
	it := g.It

	const hi = "hi"
	var sayhi = func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(hi))
	}
	var echo = func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		b, _ := ioutil.ReadAll(req.Body)
		w.Write(b)
	}

	a := New()
	a.Route("/").
		Head(sayhi).
		Get(sayhi).
		Post(echo).
		Put(echo).
		Del(echo).
		Options(sayhi)

	server := httptest.NewServer(a)
	defer server.Close()

	describe("app", func() {
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
