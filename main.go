package main

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/veracode/example-go-goget/sub"
	"github.com/veracode/example-go-goget/sub2"
)

type Options struct {
	Query   string `url:"q"`
	ShowAll bool   `url:"all"`
	Page    int    `url:"page"`
}

func main() {
	fmt.Println("Intro to Go!")
	opt := Options{"foo", true, 2}
	v, _ := query.Values(opt)
	fmt.Println(v.Encode()) // will output: "q=foo&all=true&page=2"

	sub.Foo()

	sub2.Bar()
}
