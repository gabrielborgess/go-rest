package main

import (
	"net/http"
)

// Router ...
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}


func(r *Router)FindHandler(path string,method string) (http.HandlerFunc,bool,bool){

	_,exits:=r.rules[path]
	handler,methodExist:=r.rules[path][method]
	return handler,methodExist ,exits

}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler,methodExist,exits:=r.FindHandler(request.URL.Path,request.Method)

	if !exits{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist{
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	handler(w,request)
}
