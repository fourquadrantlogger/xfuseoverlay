package main

import (
	"flag"
	"github.com/hanwen/go-fuse/v2/fs"
	"log"
)

func main() {
	flag.Parse()
	opts := &fs.Options{}
	opts.MountOptions.Debug = false

	server, err := fs.Mount(flag.Arg(0), &XRoot{
		Layers: []string{"/home/timeloveboy/test/layer1"},
	}, opts)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	server.Wait()
}
