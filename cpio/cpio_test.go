package cpio

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	type entry struct {
		header   *Header
		contents string
	}

	entries := []*entry{
		{
			header: &Header{
				Name: "regular",
				Uid:  1234,
				Gid:  4321,
				Size: 12,
				Mode: 0640,
				Type: TypeRegular,
			},
			contents: "regular file",
		},
		{
			header: &Header{
				Name: "dst",
				Uid:  1234,
				Gid:  4321,
				Size: 3, // src size
				Mode: 0777,
				Type: TypeSymlink,
			},
			contents: "src",
		},
		{
			header: &Header{
				Name: "dir",
				Uid:  1234,
				Gid:  4321,
				Size: 0,
				Mode: 0755,
				Type: TypeDir,
			},
		},
	}

	b := new(bytes.Buffer)
	w := NewWriter(b)
	for _, v := range entries {
		if err := w.WriteHeader(v.header); err != nil {
			t.Fatal(err)
		}
		if v.contents == "" {
			continue
		}
		if _, err := w.Write([]byte(v.contents)); err != nil {
			t.Fatal(err)
		}
	}
	w.Close()

	if !bytes.Equal(golden, b.Bytes()) {
		t.Fatal("have != golden")
	}
}
