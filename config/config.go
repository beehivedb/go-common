//Simple Config File.
//Copyright (C) 2020 Ron. all rights reserved.
//Version 1.0
//Date : 2020-08-09 10:25:00
//Package config - simple conf file parse.
//conf file format
//key = value in a line.
//namespace is encode by key . like db.mysql.ip = 127.0.0.1
//group by , like db.mysql.user = admin, user,
package config

import (
	"bufio"
	"os"
	"strings"
)

type SimpleConfig interface {
	Get(key string) (string, bool)
	GetArray(key string) ([]string, bool)
}

type config struct {
	content map[string][]string
}

func (c *config) Get(key string) (string, bool) {
	v, ok := c.content[key]
	if ok {
		return v[0], ok
	}
	return "", false
}

func (c *config) GetArray(key string) ([]string, bool) {
	arr, ok := c.content[key]
	if ok {
		return arr, ok
	}
	return nil, false
}

func New(file string) SimpleConfig {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	content := make(map[string][]string)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := strings.Trim(sc.Text(), " ")
		if t == "" {
			continue
		}
		arr := strings.Split(t, "=")
		if len(arr) != 2 {
			panic("SimpleConfig File Format Error, line not has = ")
		}
		key := strings.Trim(arr[0], " ")
		value := strings.Trim(arr[1], " ")

		if key == "" || value == "" {
			panic("SimpleConfig File Format Error, key or value empty.")
		}

		content[key] = strings.Split(value, ",")
	}
	return &config{content: content}
}