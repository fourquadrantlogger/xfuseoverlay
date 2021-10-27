package main

import (
	"github.com/c2h5oh/datasize"
	"path/filepath"
	"syscall"
)

type XRoot struct {
	XNode
	Layers []string

	size  uint64
	limit uint64
}

//IsDeleted when size=0 &&  mode_fstype=chardevice

func IsDeleted(name string) bool {
	var st syscall.Stat_t
	err := syscall.Stat(name, &st)
	ischar := (st.Mode & syscall.S_IFMT) == syscall.S_IFCHR
	return err == nil || (ischar && st.Size == 0)
}

func (f *XRoot) Size() (size datasize.ByteSize) {
	return datasize.ByteSize(f.size)
}
func (r *XRoot) getBranch(name string, st *syscall.Stat_t) int {
	if IsDeleted(filepath.Join(r.root().Layers[0], name)) {
		return -1
	}
	if st == nil {
		st = &syscall.Stat_t{}
	}
	for i, layer := range r.Layers {
		p := filepath.Join(layer, name)
		err := syscall.Lstat(p, st)
		if err == nil {
			return i
		}
	}
	return -1
}
