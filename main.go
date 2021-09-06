package main

import (
	"fmt"

	"cuelang.org/go/cue"
)

var cueObj1 string = `
AType: *4 | int
BType: [...int]
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
	Baz: *7 | AType
}`

func main() {
	test2()
	// var r cue.Runtime
	// i, err := r.Compile("testcases1.cue", `BType: [...int]`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// v, ok := i.Value().Lookup("BType").Default()
	// fmt.Printf("the default value %v, %v", ok, v)
	// hasDuplication := checkDuplication(i.Value().Lookup("Foo", "Baz"))
	// fmt.Printf("case 1 ........ the v expr() is: %v and the length is: %d, ops is : %v, has duplication %v \n", bval, len(bval), ops, hasDuplication)

	// ops, vvals := i.Value().Lookup("Foo", "Baz").Expr()
	// hasDuplication = checkDuplication(i.Value().Lookup("Foo", "Baz"))
	// fmt.Printf("case 1 ........ the v expr() is: %v and the length is: %d, ops is : %v, has duplication %v \n", vvals, len(vvals), ops, hasDuplication)

	// i, err = r.Compile("testcases2.cue", cueObj2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ops, vvals = i.Value().Lookup("Foo", "Baz").Expr()
	// hasDuplication = checkDuplication(i.Value().Lookup("Foo", "Baz"))
	// fmt.Printf("case 2 ........ the v expr() is: %v and the length is: %d, ops is : %v, has duplication %v \n", vvals, len(vvals), ops, hasDuplication)

	// i, err = r.Compile("testcases3.cue", cueObj3)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ops, vvals = i.Value().Lookup("Foo", "Baz").Expr()
	// hasDuplication = checkDuplication(i.Value().Lookup("Foo", "Baz"))
	// fmt.Printf("case 3 ........ the v expr() is: %v and the length is: %d, ops is : %v, has duplication %v \n", vvals, len(vvals), ops, hasDuplication)

	// i, err = r.Compile("testcases4.cue", cueObj4)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ops, vvals = i.Value().Lookup("Foo", "Baz").Expr()
	// hasDuplication = checkDuplication(i.Value().Lookup("Foo", "Baz"))
	// fmt.Printf("case 3 ........ the v expr() is: %v and the length is: %d, ops is : %v, has duplication %v \n", vvals, len(vvals), ops, hasDuplication)

}

func checkDuplication(v cue.Value) bool {

	// To capture precision when disjuct int with *4, there is something wierd happens
	// hasConcret := false
	// ops, bval := v.Eval().Expr()
	// if ops == cue.OrOp {
	// 	for _, val := range bval {
	// 		v, ok := val.Default()
	// 		if ok && v.IsConcrete() {
	// 			if hasConcret {
	// 				fmt.Printf(".........the default value is %v \n", v)
	// 				return true
	// 			} else {
	// 				hasConcret = true
	// 			}
	// 		}
	// 	}
	// }

	// To cover other cases
	var defs []cue.Value
	ops, vvals := v.Expr()
	if ops == cue.OrOp {
		for _, vval := range vvals {
			if inst, path := vval.Reference(); len(path) > 0 {
				if def, ok := inst.Lookup(path...).Default(); ok {
					defs = append(defs, def)
				} else {
					if ops, vvvals := vval.Expr(); ops == cue.OrOp && len(vvvals) > 1 {

					}
				}
			}
		}
	} else {
		if len(vvals) > 0 {
			if v, ok := vvals[0].Default(); ok {
				defs = append(defs, v)
			}
		}
	}

	def, ok := v.Default()
	if ok {
		defs = append(defs, def)
		op, dvals := def.Expr()
		fmt.Printf("the expression of default value is : ........ %v.......", dvals)
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
