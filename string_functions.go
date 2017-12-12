package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test","es"))
	p("Count: ",s.Count("test","t"))
	p("HasPrefix: ",s.HasPrefix("test","te"))
	p("Index: ",s.Index("test","e"))
	p("ToLower: ",s.ToLower("TEST"))
	p("ToUpper: ",s.ToUpper("test"))
	p()
	p("Len: ",len("Rashmi is learning Go"))
}
