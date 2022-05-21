package main

import (
	"github.com/Ethan3600/funwithgolang/application"
	http "github.com/Ethan3600/funwithgolang/transport/http/chi"
)

func main() {
	app := application.NewApplication()
	http.NewServer(app)
}
