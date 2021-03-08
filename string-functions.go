package main

import (
	"fmt"
	sx "strings"
)

var p = fmt.Println

func main() {

	p("Contains:  ", sx.Contains("test", "es"))
	p("Count:     ", sx.Count("test", "t"))
	p("HasPrefix: ", sx.HasPrefix("test", "te"))
	p("HasSuffix: ", sx.HasSuffix("test", "st"))
	p("Index:     ", sx.Index("test", "e"))
	p("Join:      ", sx.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", sx.Repeat("a", 5))
	p("Replace:   ", sx.Replace("foo", "o", "0", -1))
	p("Replace:   ", sx.Replace("foo", "o", "0", 1))
	p("Split:     ", sx.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", sx.ToLower("TEST"))
	p("ToUpper:   ", sx.ToUpper("test"))
	p()

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}