package main

import (
	"log"
	"net/http"

	"github.com/nixcloud/pankat/chat"

	// FIXME: vm-specific hack. remove this in production
	_ "github.com/wader/disable_sendfile_vbox_linux"
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
