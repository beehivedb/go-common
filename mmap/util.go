//Package mmap memory map uitls
//Copyright (C) 2020 to All Authors. all rights reserved.
//Author Ron
//Date 2020/08/24 10:53:20
//Version 1.0
package mmap

import (
	"golang.org/x/sys/unix"
)

//Random Access Memory Random.
func Random(data []byte) error {
	return unix.Madvise(data, unix.MADV_RANDOM)
}

//Lock locks all the mapped memory to RAM, preventing the pages from swapping out.
func Lock(data []byte) error {
	return unix.Mlock(data)
}

//Unlock unlocks the mapped memory from RAM, enabling swapping out of RAM if required.
func Unlock(data []byte) error {
	return unix.Munlock(data)
}
