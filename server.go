package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	host = flag.String("host", "", "The hostname/ip to listen on.")
	port = flag.Int("port", 3000, "The port number to listen on.")
)

func main() {
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	panic(http.ListenAndServe(addr, http.FileServer(http.Dir("./public"))))
}
