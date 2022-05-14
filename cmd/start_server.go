package main

import (
	"echoapp/application"
	http "echoapp/transport/http/echo"
)

func main() {
	app := application.NewApplication()
	http.NewServer(app)
}
