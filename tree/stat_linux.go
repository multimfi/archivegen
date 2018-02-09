package tree

import (
	"syscall"
	"time"
)

func stat(s *syscall.Stat_t) sysstat {
	return sysstat{
		ctime: time.Unix(s.Ctim.Sec, s.Ctim.Nsec),
		atime: time.Unix(s.Atim.Sec, s.Atim.Nsec),
	}
}
