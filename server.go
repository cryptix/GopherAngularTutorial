package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	host = flag.String("host", "", "The hostname/ip to listen on.")
	port = flag.Int("port", 3000, "The port number to listen on.")
)

func main() {
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	http.Handle("/gopath/", http.StripPrefix("/gopath", http.FileServer(http.Dir(os.Getenv("GOPATH")))))
	http.Handle("/", http.FileServer(http.Dir("./public")))

	fmt.Println("Listeing on", addr)
	panic(http.ListenAndServe(addr, nil))
}
