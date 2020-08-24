//Package mmap memory map file
//Copyright (C) 2020 to All Authors. all rights reserved.
//Author Ron
//Date 2020/08/24 09:03:51
//Version 1.0
package mmap

import (
	"errors"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

//Mmap map file to memory.
type Mmap interface {
	Open(offset int64, size int) error
	Close() error
	Flush() error
	ReadAt(dest []byte, offset int64) (int, error)
	WriteAt(src []byte, offset int64) (int, error)
}

type mmap struct {
	file     *os.File
	data     []byte
	filesz   int64
	mapsz    int
	readOnly bool
}

//Open map file.
func (m *mmap) Open(offset int64, size int) error {
	var flag int
	if m.readOnly {
		flag = syscall.PROT_READ
	} else {
		flag = syscall.PROT_READ | syscall.PROT_WRITE
	}
	b, err := unix.Mmap(int(m.file.Fd()), offset, size, flag, syscall.MAP_SHARED)
	m.data = b
	m.mapsz = size
	return err
}

//Close unmap the data.
func (m *mmap) Close() error {
	if m.data == nil {
		return nil
	}
	err := unix.Munmap(m.data)
	m.data = nil
	m.mapsz = 0
	return err
}

//Flush flush data to disk.
func (m *mmap) Flush() error {
	return unix.Msync(m.data, unix.MS_SYNC)
}

func (m *mmap) boundCheck(offset, numberBytes int64) bool {
	if m.data == nil {
		return false
	}
	if offset+numberBytes > int64(m.mapsz) || offset < 0 {
		return false
	}
	return true
}

//readAt read at offset from db.
func (m *mmap) ReadAt(dest []byte, offset int64) (int, error) {
	if m.boundCheck(offset, 1) {
		return copy(dest, m.data[offset:]), nil
	}
	return 0, errors.New("Index Out Of Bound")
}

//writeAt write data at position of db.
func (m *mmap) WriteAt(src []byte, offset int64) (int, error) {
	if m.boundCheck(offset, 1) {
		return copy(m.data[offset:], src), nil
	}
	return 0, errors.New("Index Out Of Bound")
}

func newMmap(f *os.File, ro bool) (Mmap, error) {
	if f == nil {
		return nil, errors.New("file is nil")
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return &mmap{
		file:     f,
		filesz:   fi.Size(),
		readOnly: ro,
	}, nil
}

//NewRO create new Read Only Mmap instance.
func NewRO(f *os.File) (Mmap, error) {
	return newMmap(f, true)
}

//NewRW create new read and write mmap.
func NewRW(f *os.File) (Mmap, error) {
	return newMmap(f, false)
}
