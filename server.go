package main

import (
	"fmt"
	"log"
	"net"

	"github.com/suedoh/go-cache/cache"
)

type ServerOpts struct {
    ListenAddr string
    IsLeader bool
}

type Server struct {
    ServerOpts 

    cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
    return &Server{
        ServerOpts: opts,
        cache: c,
    }
}

func (s *Server) Start() error  {
    ln, err := net.Listen("tcp", s.ListenAddr)
    if err != nil {
        return fmt.Errorf("listen error: %s", err)
    }

    log.Printf("server starting on %s\n", s.ListenAddr)

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Printf("accept error %s\n", err)
            continue
        }

        go s.handleConn(conn)
    }

    return nil
}

func (s *Server) handleConn(conn net.Conn)  {
    buf := make([]byte, 2048)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            log.Printf("conn read error:%s\n", err)
            break
        }

        msg := buf[:n]
        fmt.Println(string(msg))
    }
}
