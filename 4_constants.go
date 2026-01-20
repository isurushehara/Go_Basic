// Constants
/*
If a variable should have a fixed value that cannot be changed, you can use the const keyword.
The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
*/

// const CONSTNAME type = value

package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}

// Constant Rules
/*
Constant names follow the same naming rules as variables
Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
Constants can be declared both inside and outside of a function
*/

// Constant Types
/*
There are two types of constants:

- Typed constants: When a constant is declared with a specific type.

package main
import ("fmt")

const A int = 1

func main() {
  fmt.Println(A)
}

- Untyped constants: When a constant is declared without a specific type.

package main
import ("fmt")

const A int = 1

func main() {
  fmt.Println(A)
}
  

- Constants: Unchangeable and Read-only
When a constant is declared, it is not possible to change the value later:

Example
package main
import ("fmt")

func main() {
  const A = 1
  A = 2
  fmt.Println(A)
}
Result:
./prog.go:8:7: cannot assign to A


*/



