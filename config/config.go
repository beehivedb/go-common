//Package config Simple Config File.
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

//SimpleConfig - config interface
type SimpleConfig interface {
	Get(key string, value string) string
	GetArray(key string) []string
}

type config struct {
	content map[string][]string
}

//Get - get config by key. if unset then return default.
func (c *config) Get(key string, value string) string {
	v, ok := c.content[key]
	if ok {
		return v[0]
	}
	return value
}

func (c *config) GetArray(key string) []string {
	arr, ok := c.content[key]
	if ok {
		return arr
	}
	return nil
}

//New - create new config.
func New(file string) SimpleConfig {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	content := make(map[string][]string)

	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			break
		}

		t := strings.Trim(string(line), " ")
		if t == "" || strings.HasPrefix(t, "#") {
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
