package main

import (
	"flag"
	"fmt"

	"github.com/codegangsta/negroni"
)

var (
	host = flag.String("host", "", "The hostname/ip to listen on.")
	port = flag.Int("port", 3000, "The port number to listen on.")
)

func main() {
	flag.Parse()

	n := negroni.Classic()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	n.Run(addr)
}
