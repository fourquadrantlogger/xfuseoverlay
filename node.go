package main

import (
	"context"
	"fmt"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
	"syscall"
)

type XNode struct {
	fs.Inode
}

func (n *XNode) root() *XRoot {
	return n.Root().Operations().(*XRoot)
}

var _ = (fs.NodeGetattrer)((*XNode)(nil))

func (n *XNode) Getattr(ctx context.Context, f fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	fmt.Println("Getattr")
	return 0
}
