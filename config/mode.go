package config

import "os"

const digits = "01234567"

func mode(f os.FileInfo) int {
	m := f.Mode()
	r := m.Perm()

	if m&os.ModeSticky != 0 {
		r |= modeSticky
	}
	if m&os.ModeSetgid != 0 {
		r |= modeSetgid
	}
	if m&os.ModeSetuid != 0 {
		r |= modeSetuid
	}
	return int(r)
}
