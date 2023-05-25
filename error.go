package main

import "errors"

var (
	errResourceNotSupport = errors.New("your specified resource is not supported")
	errOptionNotSupport   = errors.New("this option is not yet supported")
)
