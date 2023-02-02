package main

import (
    "github.com/suedoh/go-cache/cache"
)

func main()  {
    opts := ServerOpts{
        ListenAddr: ":3000",
        IsLeader: true,
    }

    server := NewServer(opts, cache.New())
    server.Start()
}
