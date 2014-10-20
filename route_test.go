package app

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestRoute(t *testing.T) {
	g := Goblin(t)

	a := New()
	a.Route("/").
		Head(sayhi).
		Get(sayhi).
		Post(echo).
		Put(echo).
		Del(echo).
		Options(sayhi)

	testApp(a, g, "route-app")
}
