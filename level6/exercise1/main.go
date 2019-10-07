// Review
// * functions
// * purpose of functions
//     + abstract code
//     + code reusability
//             - DRY - Don't Repeat Yourself
// * func, receiver, identifier, params, returns
// * parameters vs arguments
// * variadic funcs
//     + multiple "variadic" params
//     + multiple "variadic" args
// * returns
//     + multiple returns
//     + named return - yuck!
// * func expressions
//     + assigning a func to a variable
// * callback
//     + passing a func into another func as an argument
// * closure
//     + one scope enclosing another
//     + variable declared in the outer scope are accessible in the inner scopes
//     + closure helps us limit the scope of variables
// * recursion
//     + a func that calls itself
//     + factorial
// * life philosophy
//     + focus on what's important; not upon what's urgent

package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "God is Great."
}

func main() {
	f := foo()
	bi, bs := bar()
	fmt.Println(f)
	fmt.Println(bi, bs)
}
