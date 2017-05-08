package main

import (
	"log"
	"flag"
	"strconv"

	"github.com/emersion/go-imap/server"

	"github.com/donghui/imap-mock-server/backend/memory"
)

var port int

func init() {
	log.Println("init()")
	flag.IntVar(&port, "port", 1143, "Port number for connection")
	flag.Parse()
}
func main() {
	// Create a memory backend
	be := memory.New()

	// Create a new server
	s := server.New(be)
	s.Addr = ":" + strconv.Itoa(port)
	// Since we will use this server for testing only, we can allow plain text
	// authentication over unencrypted connections
	s.AllowInsecureAuth = true

	log.Printf("Starting IMAP server at localhost:%d\n", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

