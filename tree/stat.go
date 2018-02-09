package tree

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

type sysstat struct {
	ctime time.Time
	atime time.Time
}

func statt(file os.FileInfo) (sysstat, error) {
	r, ok := file.Sys().(*syscall.Stat_t)
	if !ok {
		return sysstat{}, fmt.Errorf("!= *syscall.Stat_t, %#v", file.Sys())
	}
	return stat(r), nil
}
