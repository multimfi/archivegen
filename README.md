# archivegen

```
archivegen [OPTIONS...] [FILES...]

  -fmt string
    	file format, cpio/tar (default "tar")
  -out string
    	output file (default "out.archive")
  -print
    	print resolved tree in archivegen format
  -stdout
    	output to stdout
  -version
    	version

Format:
  '*' required
  '-' omit

  Directory
    d *dst mode uid gid

  Symlink
    l *dst *src uid gid

  File
    f *src dst mode uid gid

  Recursive
    // Omitted dst will target archive root.
    R *src *dst uid gid

  Create
    // Data is everything after gid + ' ' or '\t'.
    c *dst mode uid gid *data...

  Elf
    L *elf dst mode uid gid
```
