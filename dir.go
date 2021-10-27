package main

import (
	"context"
	"fmt"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
	"path/filepath"
	"syscall"
)

var _ = (fs.NodeLookuper)((*XNode)(nil))

func (n *XNode) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	fmt.Println("Lookup", name)
	if n.IsRoot() && IsDeleted(filepath.Join(n.root().Layers[0], name)) {
		return nil, syscall.ENOENT
	}
	var st syscall.Stat_t

	p := filepath.Join(n.Path(nil), name)
	idx := n.root().getBranch(p, &st)
	if idx >= 0 {
		// XXX use idx in Ino?
		ch := n.NewInode(ctx, &XNode{}, fs.StableAttr{Mode: st.Mode, Ino: st.Ino})
		out.FromStat(&st)
		out.Mode |= 0111
		return ch, 0
	}
	return nil, syscall.ENOENT
}

var _ = (fs.NodeReaddirer)((*XNode)(nil))

func (n *XNode) Readdir(ctx context.Context) (fs.DirStream, syscall.Errno) {
	fmt.Println("Readdir")

	names := map[string]uint32{}

	result := make([]fuse.DirEntry, 0, len(names))
	for nm, mode := range names {
		result = append(result, fuse.DirEntry{
			Name: nm,
			Mode: mode,
		})
	}

	return fs.NewListDirStream(result), 0
}
