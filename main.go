package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue"
)

var cueObj1 string = `
AType: *4 | int
Foo: {
	Baz: AType | *"call"
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

var cueObj4 string = `
AType: *4 | int
Foo: {
	Baz: AType | *7
}`

func main() {

	var r cue.Runtime
	i, err := r.Compile("testcases1.cue", cueObj1)
	if err != nil {
		log.Fatal(err)
	}

	ops, bval := i.Value().Lookup("Foo", "Baz").Eval().Expr()
	hasDuplication := checkDuplication(i.Value().Lookup("Foo", "Baz"))
	fmt.Printf("case 1 ........ the v expr() is: %v and the length is: %d, ops is : %v, has duplication %v \n", bval, len(bval), ops, hasDuplication)

	ops, vvals := i.Value().Lookup("Foo", "Baz").Expr()
	hasDuplication = checkDuplication(i.Value().Lookup("Foo", "Baz"))
	fmt.Printf("case 1 ........ the v expr() is: %v and the length is: %d, ops is : %v, has duplication %v \n", vvals, len(vvals), ops, hasDuplication)

	i, err = r.Compile("testcases2.cue", cueObj2)
	if err != nil {
		log.Fatal(err)
	}
	ops, vvals = i.Value().Lookup("Foo", "Baz").Expr()
	hasDuplication = checkDuplication(i.Value().Lookup("Foo", "Baz"))
	fmt.Printf("case 2 ........ the v expr() is: %v and the length is: %d, ops is : %v, has duplication %v \n", vvals, len(vvals), ops, hasDuplication)

	i, err = r.Compile("testcases3.cue", cueObj3)
	if err != nil {
		log.Fatal(err)
	}
	ops, vvals = i.Value().Lookup("Foo", "Baz").Expr()
	hasDuplication = checkDuplication(i.Value().Lookup("Foo", "Baz"))
	fmt.Printf("case 3 ........ the v expr() is: %v and the length is: %d, ops is : %v, has duplication %v \n", vvals, len(vvals), ops, hasDuplication)

	i, err = r.Compile("testcases4.cue", cueObj3)
	if err != nil {
		log.Fatal(err)
	}
	ops, vvals = i.Value().Lookup("Foo", "Baz").Expr()
	hasDuplication = checkDuplication(i.Value().Lookup("Foo", "Baz"))
	fmt.Printf("case 3 ........ the v expr() is: %v and the length is: %d, ops is : %v, has duplication %v \n", vvals, len(vvals), ops, hasDuplication)

}

func checkDuplication(v cue.Value) bool {
	var defs []cue.Value

	// To capture precision when disjuct int with *4, there is something wierd happens
	ops, bval := v.Eval().Expr()
	if ops == cue.OrOp {
		for _, val := range bval {
			v, ok := val.Default()
			fmt.Printf(".........the default value is %v \n", v)
			if ok && v.IsConcrete() {
				defs = append(defs, v)
			}
		}
	}

	// To cover other cases
	ops, vvals := v.Expr()
	if ops == cue.OrOp {
		for _, vval := range vvals {
			if inst, path := vval.Reference(); len(path) > 0 {
				if def, ok := inst.Lookup(path...).Default(); ok {
					defs = append(defs, def)
				}
			}
		}
	}

	def, ok := v.Default()
	if ok {
		defs = append(defs, def)
		op, dvals := def.Expr()
		if len(dvals) > 1 && op == cue.OrOp {
			return true
		}
	}
	if len(defs) >= 1 {
		for _, def := range defs[1:] {
			if !defs[0].Equals(def) {
				return true
			}
		}
	}
	return false
}
