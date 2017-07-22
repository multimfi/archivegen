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
    // dst is the filename in archive
    // 'l to from' = 'from -> to'

    l *dst *src uid gid

  File
    f *src dst mode uid gid

  Recursive
    // omitted dst will target archive root
    // src path is stripped from dst

    R *src *dst uid gid

  Create
    // all preceding ' ' and '\n' are stripped
    // from data and file is '\n' terminated
    // 'c file - - -	 foo  bar  ' = 'foo  bar  '

    c *dst mode uid gid *data

  Elf
    L *elf dst mode uid gid
```
