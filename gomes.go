// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
)

var (
	fcgiMode = flag.Bool("fcgi", false, "fcgi mode (default: false)")
	host     = flag.String("host", "localhost", "host name to which "+
		"application modules should bind (default: localhost)")
	port = flag.Int("port", 8080, "lowest port to which application "+
		"modules should bind (default: 8080)")
	path = flag.String("path", "./", "path to serve files (default: ./)")
)

func main() {
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("%s: Running...", addr)

	if *fcgiMode {
		l, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(fcgi.Serve(l, http.FileServer(http.Dir(*path))))
	} else {
		log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir(*path))))
	}
}
