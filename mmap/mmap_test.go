//Package mmap for test
//Copyright (C) 2020 to All Authors. all rights reserved.
//Author Ron
//Date 2020/08/23 08:26:14
//Version 1.0
package mmap

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

var (
	data = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	file = "test.db"
)

func before(t *testing.T) *os.File {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("error in opening file :: %v", err)
	}

	if _, err := f.Write([]byte(data)); err != nil {
		t.Fatalf("error in writing to file :: %v", err)
	}

	if err := f.Sync(); err != nil {
		t.Fatalf("flush to file error : %s", err)
	}

	return f
}

func after(f *os.File, t *testing.T) {
	if err := f.Close(); err != nil {
		t.Fatalf("error in closing file :: %v", err)
	}

	if err := os.Remove(file); err != nil {
		t.Fatalf("error in deleting file :: %v", err)
	}
}

func TestUnmap(t *testing.T) {
	f := before(t)
	defer after(f, t)

	m, err := NewRO(f)
	if err != nil {
		t.Fatalf("error in create new instance : %v", err)
	}

	if err := m.Open(0, len(data)); err != nil {
		t.Fatalf("error in mapping :: %v", err)
	}

	if err := m.Close(); err != nil {
		t.Fatalf("error in unmapping :: %v", err)
	}

}

func TestReadWrite(t *testing.T) {
	f := before(t)
	defer after(f, t)

	m, err := NewRW(f)
	if err != nil {
		t.Fatalf("error in new mmap: %v", err)
	}

	if err := m.Open(0, len(data)); err != nil {
		t.Fatalf("error in mapping :: %v", err)
	}
	defer func() {
		if err := m.Close(); err != nil {
			t.Fatalf("error in unmapping :: %v", err)
		}
	}()

	_, err = m.WriteAt([]byte("X"), 9)
	if err != nil {
		t.Fatalf("write at error: %v", err)
	}
	if err := m.Flush(); err != nil {
		t.Fatalf("flush to file error: %s", err)
	}

	f1, _ := os.OpenFile(file, os.O_RDONLY, 0644)
	b, err := ioutil.ReadAll(f1)
	if err != nil {
		t.Fatalf("error reading file: %s", err)
	}

	if !bytes.Equal(b, []byte("012345678XABCDEFGHIJKLMNOPQRSTUVWXYZ")) {
		t.Fatalf("file wan't modify. %d", len(b))
	}
	f1.Close()
}
