//Package router http router library. simple fast based on net/router .
//Copyright (C) 2020 To All Authors. All rights reserved.
//Author: Ron.
//Date: 2020-08-08
//Version: 1.0
package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Context http context.
type Context interface {
	Get(string) string
	Put(string)
	Body() string
	Method() string
	Response() http.ResponseWriter
	Request() *http.Request
}

//Controller - Current Controller.
type Controller interface {
	Path() string
	Execute(ctx Context)
}

type context struct {
	writer  http.ResponseWriter
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

//ParseJSON - unmarshal requet body to json.
func (c *context) ParseJSON(v interface{}) error {
	if c.request.Body != nil {
		arr, err := ioutil.ReadAll(c.request.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(arr, v)
	}
	return errors.New("Request Body Is Empty")
}

//JSON - Output format json
func (c *context) JSON(v interface{}) error {
	arr, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = c.writer.Write(arr)
	return err
}

//Failed - send failed message.
func (c *context) Failed(msg string) {
	c.writer.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(c.writer, msg)
}

//Cookie - query cookie.
func (c *context) Cookie(name string) (*http.Cookie, error) {
	return c.request.Cookie(name)
}

//AddCookie - add cookie.
func (c *context) AddCookie(name string, value string) {
	k := http.Cookie{
		Name:  name,
		Value: value,
	}
	c.writer.Header().Set("Set-Cookie", k.String())
}

//H - shutcut of map.
type H map[string]interface{}

func (c *context) Method() string {
	return c.Method()
}

func (c *context) Response() http.ResponseWriter {
	return c.writer
}

func (c *context) Request() *http.Request {
	return c.request
}

//Action -
type Action func(ctx Context)

//RegistryAction -
func RegistryAction(path string, action Action) {
	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		c := &context{
			writer:  writer,
			request: request,
		}
		action(c)
	})
}

//Registry -
func Registry(controller Controller) {
	p := controller.Path()
	http.HandleFunc(p, func(writer http.ResponseWriter, request *http.Request) {
		controller.Execute(&context{
			writer:  writer,
			request: request,
		})
	})
}

//StaticPath - static file.
func StaticPath(path string, dir string) {
	http.Handle(path, http.FileServer(http.Dir(dir)))
}

//Launch - run http server.
func Launch(addr string) {
	log.Fatal(http.ListenAndServe(addr, nil))
}
