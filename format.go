package main

// TODO: ignore/user/group/mode/copy/hardlink/rename/prefix mask.

const helpFormat = `
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
    // elf is prefixed with rootfs when it is not omitted
    L *elf dst mode uid gid rootfs`
