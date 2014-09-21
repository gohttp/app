package main

import "github.com/gohttp/app"
import "net/http"
import "fmt"

func main() {
	app := app.New()

	app.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	}))

	app.Listen(":3000")
}
