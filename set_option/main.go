// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
// https://www.alexedwards.net/blog/organising-database-access

package main

import "fmt"

type option func(*F)

type F struct {
	K string
}

func (f *F) Option(opts ...option) {
	for _, opt := range opts {
		opt(f)
	}
}

func OptK(s string) option {
	return func(f *F) {
		f.K = s
	}
}

func main() {
	f := F{}
	fmt.Println(f)
	f.Option(OptK("aaa"))
	fmt.Println(f)
}
