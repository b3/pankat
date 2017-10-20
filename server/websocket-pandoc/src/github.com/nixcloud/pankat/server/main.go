package main

import (
	"log"
	"net/http"

	"github.com/nixcloud/pankat/chat"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// websocket server
	server := chat.NewServer("/entry")
	go server.Listen()

	// static files
	// FIXME: vm-specific absolute path
	http.Handle("/", http.FileServer(http.Dir("/root/hostFS/pankat/server/websocket-pandoc/webroot")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
