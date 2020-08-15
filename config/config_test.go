//Package config - config parse test.package config
//Copyright (C) 2020 to All Authors. all rights reserved.
//Author Ron
//Date 2020/08/15 12:29:27
//Version 1.0
package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	c := New("test.conf")
	_, ok := c.Get("abc")
	if ok {
		t.Fail()
	}
	v, _ := c.Get("db.mysql.ip")
	if v != "127.0.0.1" {
		t.FailNow()
	}

}
