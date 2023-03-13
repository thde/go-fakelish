package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	flag "github.com/spf13/pflag"
)

const defaultDir = "./public"

var port = flag.IntP("port", "p", 8080, "Port")

func main() {
	flag.Parse()
	dir := defaultDir
	if flag.Arg(0) != "" {
		dir = flag.Arg(0)
	}

	addr := fmt.Sprintf("localhost:%d", *port)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen on %s error: %s", addr, err.Error())
	}
	fs := http.FileServer(http.Dir(dir))

	log.Print("serving " + dir + " on http://" + addr)
	http.Serve(ln, logRequest(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("cache-control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	})))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
