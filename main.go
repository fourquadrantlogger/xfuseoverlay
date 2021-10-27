package main

import (
	"flag"
	"github.com/hanwen/go-fuse/v2/fs"
	"log"
)

type XRoot struct {
	fs.Inode
}

func main() {
	opts := &fs.Options{}
	server, err := fs.Mount(flag.Arg(0), &XRoot{}, opts)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	server.Wait()
}
