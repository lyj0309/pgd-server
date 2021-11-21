package main

import "net/http"

type fuck interface {
	Name() string
	Bind(*http.Request, interface{}) error
}
