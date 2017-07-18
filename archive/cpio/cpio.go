package cpio

import (
	"io"

	"bitbucket.org/multimfi/archivegen/archive"
	"bitbucket.org/multimfi/archivegen/cpio"
)

const max32 = int64(^uint32(0))

type writer struct {
	cw *cpio.Writer
}

func NewWriter(w io.Writer) archive.Writer {
	cw := cpio.NewWriter(w)
	return &writer{
		cw: cw,
	}
}

func (w *writer) Close() error {
	return w.cw.Close()
}

func (w *writer) Write(b []byte) (int, error) {
	return w.cw.Write(b)
}

func typeconv(t archive.FileType) int {
	switch t {
	case archive.TypeDir:
		return cpio.TypeDir
	case archive.TypeFifo:
		return cpio.TypeFifo
	case archive.TypeChar:
		return cpio.TypeChar
	case archive.TypeBlock:
		return cpio.TypeBlock
	case archive.TypeRegular:
		return cpio.TypeRegular
	case archive.TypeSymlink:
		return cpio.TypeSymlink
	case archive.TypeSocket:
		return cpio.TypeSocket
	default:
		panic("unknown type " + t.String())
	}

}

func hdrconv(a *archive.Header) *cpio.Header {
	if a.Size >= max32 {
		panic("filesize")
	}
	if a.Mode >= max32 {
		panic("filemode")
	}

	return &cpio.Header{
		Name: a.Name,
		Uid:  a.Uid,
		Gid:  a.Gid,
		Size: a.Size,
		Mode: int(a.Mode),
		Type: typeconv(a.Type),
	}
}

func (w *writer) WriteHeader(hdr *archive.Header) error {
	return w.cw.WriteHeader(hdrconv(hdr))
}

func (w *writer) Symlink(src, dst string, uid, gid int) error {
	hdr := &cpio.Header{
		Name: dst,
		Size: int64(len(src)),
		Mode: 0777,
		Uid:  uid,
		Gid:  gid,
		Type: cpio.TypeSymlink,
	}
	if err := w.cw.WriteHeader(hdr); err != nil {
		return err
	}
	if _, err := w.Write([]byte(src)); err != nil {
		return err
	}
	return nil
}
