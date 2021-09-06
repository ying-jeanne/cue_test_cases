package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue"
)

var cueObjtest2 string = `
AType: *4 | int
BType: *7 | int
_Foo: {
	Baz: AType | BType
}`

func test2() {
	var r cue.Runtime
	i, err := r.Compile("testcases1.cue", cueObjtest2)
	if err != nil {
		log.Fatal(err)
	}
	fields, _ := i.Value().Fields(cue.Hidden(true))
	for fields != nil && fields.Next() {
		ishidden1 := fields.IsHidden()
		ishidden2 := fields.Selector().PkgPath() != ""
		fmt.Printf("the result of two %v is hidden:\n %v, %v, \n the pkgpath is: %v\n", fields.Label(), ishidden1, ishidden2, fields.Selector().PkgPath())
	}
}
