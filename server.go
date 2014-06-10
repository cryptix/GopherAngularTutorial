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

	// gopherjs generates source maps for your go code, with this you can naviagte call stacks
	// and even set breakpoints (which than map to the generated .js autoamtically)
	http.Handle("/gopath/", http.StripPrefix("/gopath", http.FileServer(http.Dir(os.Getenv("GOPATH")))))

	http.Handle("/", http.FileServer(http.Dir("./public")))

	addr := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Println("Listeing on", addr)
	panic(http.ListenAndServe(addr, nil))
}
