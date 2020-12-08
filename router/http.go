package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Action - handle context
type Action func(hc *HTTPContext)

//Interceptor - http filter.
type Interceptor func(hc *HTTPContext) bool

//KV - map
type KV map[string]interface{}

//HTTPContext web request.
type HTTPContext struct {
	W  http.ResponseWriter
	R  *http.Request
	PS httprouter.Params
}

//Header - write header.
func (hc *HTTPContext) Header(key, value string) {
	hc.W.Header()[key] = []string{value}
}

//GetHeader - get header value.
func (hc *HTTPContext) GetHeader(key string) string {
	return hc.R.Header[key][0]
}

//BodyData - read body data.
func (hc *HTTPContext) BodyData() ([]byte, error) {
	defer hc.R.Body.Close()
	return ioutil.ReadAll(hc.R.Body)
}

//GetParam get request parameter or path parameter.
func (hc *HTTPContext) GetParam(key string) string {
	params := httprouter.ParamsFromContext(hc.R.Context())
	return params.ByName(key)
}

//Write - data.
func (hc *HTTPContext) Write(bs []byte) {
	hc.W.Write(bs)
	hc.W.WriteHeader(http.StatusOK)
}

//WriteCode - write http status.
func (hc *HTTPContext) WriteCode(code int) {
	hc.W.WriteHeader(code)
}

//WriteString - write string value.
func (hc *HTTPContext) WriteString(data string) {
	hc.W.Write([]byte(data))
}

//WriteJSON - write key value to json.
func (hc *HTTPContext) WriteJSON(kv KV) {
	bs, err := json.Marshal(kv)
	if err != nil {
		hc.W.WriteHeader(http.StatusBadRequest)
		hc.W.Write([]byte(err.Error()))
		return
	}
	hc.W.Write(bs)
}
