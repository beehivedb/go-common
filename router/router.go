//Copyright (C) 2020 To All Authors. All rights reserved.
//Author: Ron.
//Date: 2020-08-08
//Version: 1.0
//Package router http router library. simple fast based on net/router .
package router

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Context interface {
	Get(string) string
	Put(string)
	Body() string
	Method() string
	Response() http.ResponseWriter
	Request() *http.Request
}

type Controller interface {
	Path() string
	Execute(ctx Context)
}

type context struct {
	writer http.ResponseWriter
	request *http.Request
}

func (c *context) Get(key string) string {
	err := c.request.ParseForm()
	if err == nil {
		return c.request.Form.Get(key)
	}
	return ""
}

func (c *context) Put(value string) {
	_, err := fmt.Fprint(c.writer, value)
	if err != nil {
		panic(err)
	}
}

func (c *context) Body() string {
	bytes, err := ioutil.ReadAll(c.request.Body)
	if err == nil {
		return string(bytes)
	}
	return ""
}

func (c *context) Method() string {
	return c.Method()
}

func (c *context) Response() http.ResponseWriter {
	return c.writer
}

func (c *context) Request() *http.Request {
	return c.request
}

type Action func(ctx Context)

func RegistryAction(path string, action Action) {
	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		c := &context{
			writer: writer,
			request: request,
		}
		action(c)
	})
}

func Registry(controller Controller){
	p := controller.Path()
	http.HandleFunc(p, func(writer http.ResponseWriter, request *http.Request) {
		controller.Execute(&context{
			writer: writer,
			request: request,
		})
	})
}

func StaticPath(path string, dir string) {
	http.Handle(path, http.FileServer(http.Dir(dir)))
}

func Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, nil))
}

