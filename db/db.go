//Package db key/value database.
//Copyright (C) 2020 to All Authors. all rights reserved.
//Author Ron
//Date 2020/08/22 21:59:36
//Version 1.0
package db

import (
	"os"
)

const (
	//default max file size 1G
	defaultMaxFileSize = 1 << 30
	//default map size 64M
	defaultMemMapSize = 64 * (1 << 20)
)

//DB represents a collection of buckets presisted to a file on disk.
type DB struct {
	file *os.File
	data []byte
	size int
}
