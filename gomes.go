// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gomes

import (
	"flag"
	"fmt"
	// "net/http"
	// "net/http/fcgi"
)

var (
	fcgi = flag.Bool("fcgi", false, "fcgi mode (default: false)")
	host = flag.String("host", "localhost", "host name to which "+
		"application modules should bind (default: localhost)")
	port = flag.Int("port", 8080, "lowest port to which application "+
		"modules should bind (default: 8080)")
)

func init() {
	flag.Parse()
}

type Server struct {
	Host string
}

func NewServer() *Server {
	return &Server{
		Host: "test..",
	}
}

func (s *Server) Run() {
	fmt.Println(s.Host)
}
