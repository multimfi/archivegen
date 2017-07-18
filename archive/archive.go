package archive

import "io"

//go:generate stringer -type=FileType
type FileType uint

const (
	TypeDir FileType = iota
	TypeFifo
	TypeChar
	TypeBlock
	TypeRegular
	TypeSymlink
	TypeSocket
)

type Header struct {
	Name string   // name of header file entry.
	Mode int64    // permission and mode bits.
	Uid  int      // user id of owner.
	Gid  int      // group id of owner.
	Size int64    // length in bytes.
	Type FileType // filetype.
}

type Writer interface {
	io.WriteCloser

	WriteHeader(hdr *Header) error

	Symlink(src, dst string, uid, gid int) error
}
