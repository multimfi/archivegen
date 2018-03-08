# archivegen

```
archivegen [OPTIONS...] [FILES...]
  -X value
    	variable
    	e.g. '-X foo=bar -X a=b'
  -fmt string
    	file format, cpio/tar (default "tar")
  -format
    	print archive format
  -out string
    	output file (default "out.archive")
  -print
    	print resolved tree in archivegen format
  -rootfs string
    	ELF rootfs
  -stdout
    	output to stdout
  -timestamp
    	preserve file timestamps
  -version
    	version
```

## file format
```
Format:
  '*' required
  '-' omit

  Variable
    // variables do not apply across files
    $ *name value

  Directory
    d *dst mode uid gid

  Symlink
    // dst is the filename in archive
    // 'l to from' = 'from -> to'

    l *dst *src uid gid

  File
    f *src dst mode uid gid
    // fr is relative

  Recursive
    // omitted dst will target archive root
    // src path is stripped from dst

    R *src *dst uid gid
    // Rr is relative

  Create
    // all preceding ' ' and '\t' are stripped
    // from data and file is '\n' terminated
    // 'c file - - -	 foo  bar  ' = 'foo  bar  '

    c *dst mode uid gid *data

  Elf
    // elf is prefixed with rootfs when it is not omitted
    L *elf dst mode uid gid rootfs

Masks:
  Mode
    mm *idx *regexp mode uid gid

  Ignore
    mi *idx *regexp
    // mI is reversed

  Rename
    mr *idx *regexp *dst

  Clear
    mc idx
```
