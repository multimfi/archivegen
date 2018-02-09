package tree

import (
	"syscall"
	"time"
)

func stat(s *syscall.Stat_t) sysstat {
	return sysstat{
		ctime: time.Unix(s.Ctimespec.Sec, s.Ctimespec.Nsec),
		atime: time.Unix(s.Atimespec.Sec, s.Atimespec.Nsec),
	}
}
