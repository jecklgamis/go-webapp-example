package main

import (
	"flag"
	"github.com/jecklgamis/go-api-server-template/pkg/server"
)

var (
	port = flag.Int("port", 8080, "The HTTP port")
)

func main() {
	flag.Parse()
	server.Start(*port)
}
