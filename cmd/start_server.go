package main

import (
	"github.com/Ethan3600/funwithgolang/application"
	http "github.com/Ethan3600/funwithgolang/transport/http/echo"
)

func main() {
	app := application.NewApplication()
	http.NewServer(app)
}
