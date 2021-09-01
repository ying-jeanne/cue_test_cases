package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue"
)

var cueObj1 string = `
AType: *4 | int | string
Foo: {
	Baz: AType | *7
}`

var cueObj2 string = `
AType: *4 | int
BType: *7 | int
Foo: {
	Baz: AType | BType
}`

var cueObj3 string = `
Foo: {
	Baz: *4 | *7
}`

func main() {

	var r cue.Runtime
	i, err := r.Compile("testcases1.cue", cueObj1)
	if err != nil {
		log.Fatal(err)
	}
	ops, vvals := i.Value().Lookup("Foo", "Baz").Expr()
	fmt.Printf("case 1 ........ the v expr() is: %v and the length is: %d, ops is : %v\n", vvals, len(vvals), ops)

	i, err = r.Compile("testcases2.cue", cueObj2)
	if err != nil {
		log.Fatal(err)
	}
	ops, vvals = i.Value().Lookup("Foo", "Baz").Expr()
	fmt.Printf("case 2 ........ the v expr() is: %v and the length is: %d, ops is : %v\n", vvals, len(vvals), ops)

	i, err = r.Compile("testcases3.cue", cueObj3)
	if err != nil {
		log.Fatal(err)
	}
	ops, vvals = i.Value().Lookup("Foo", "Baz").Expr()
	fmt.Printf("case 3 ........ the v expr() is: %v and the length is: %d, ops is : %v\n", vvals, len(vvals), ops)

}
