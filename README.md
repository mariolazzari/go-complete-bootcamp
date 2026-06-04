# Go complete bootcamp

## Getting started

### Go playground

[Go playground](https://go.dev/blog/playground)

### Go program structure

```go
/////////////////////////////////
// The Typical Structure of a Go Application
// Go Playground: https://play.golang.org/p/vY_IeYBb1GN
/////////////////////////////////

// a package clause starts every source file
// main is a special name declaring an executable rather than a library (package)
package main

// import declaration declares packages used in this file
import "fmt"

// package scoped variables and constants
var x int = 100

const y = 0

// a function declaration. main is a special function name
// it is the entry point for the executable program
// Go uses brace brackets to delimit code blocks
func main() {

    // Local Scoped Variables and Constants Declarations, calling functions etc
    var a int = 7
    var b float64 = 3.5
    const c int = 10

    // Println() function prints out a line to stdout
    // It belongs to package fmt
    fmt.Println("Hello Go world!") // => Hello Go world!
    fmt.Println(a, b, c)           // => 7 3.5 10

}
```

### Multiple declarations

[Multi vars](https://go.dev/blog/declaration-syntax)

### Declarations

```go
/////////////////////////////////
// Variables and Declarations
// Go Playground: https://play.golang.org/p/PKdAxUp8mNT
/////////////////////////////////

package main

import "fmt"

func main() {

    //** DECLARING VARIABLES **///

    // Syntax: var var_name type
    var s1 string
    s1 = "Learning Go!"
    fmt.Println(s1) // printing string s1

    //** TYPE INFERENCE **//

    // Go deduces automatically the type of the variable by looking at the initial value (bool, int, string etc)

    var k int = 6 // not necessary to say the type (int). It is inferred from the literal on the right side of =
    var i = 5     // type int
    var j = 5.6   // type float64

    // printing i, j and k
    fmt.Println("i:", i, "j:", j, "k:", k)

    // ii == jj  // -> error: cannot assign float to int (Go is a strong typed language)

    // declaring and initializing a new variable of type string (type inference)
    var s2 = "Go!"
    _ = s2 //in Go each variable must be used or there is a compile-time error
    // _ is the Blank Identifier and mutes the error of unused variables
    // _ can be only on the left hand side of the = operator

    // multiple assignments
    var ii, jj int
    ii, jj = 5, 8 // -> tuple assignment. It allows several variables to be assigned at once

    // swapping two variables using multiple assignments
    ii, jj = jj, ii

    fmt.Println(ii, jj)

    //** Short Declaration (works only in Block Scope!) **//

    // := (colon equals syntax) used only when declaring a new variable (or at least a new variable)
    // := tells go we are going to create a new variable and go figures out what type it will be
    s3 := "Learning golang!"
    _ = s3

    // can't use short declaration at Package Scope (outside main() or other function)
    // all statements at package scope must start with a Go keyword (package, var, import, func etc)

    // multiple short declaration
    car, cost := "Audi", 50000
    fmt.Println(car, cost)

    // redeclaration with short declaration syntax
    // at least one variable must be NEW on the left side of :=
    var deleted = false
    deleted, file := true, "a.txt"
    _, _ = deleted, file

    // expressions in short declarations are allowed
    sum := 5 + 2.5
    fmt.Println(sum)

    // multiple declaration is good for readability
    var (
        age       float64
        firstName string
        gender    bool
    )
    _, _, _ = age, firstName, gender

    // a concise way to declare multiple variables that have the same type
    var a, b int
    _, _ = a, b

}
```

### Zero values

```go
package main

import "fmt"

func main() {

    // you must provide a type for each variable you declare or Go should be able to infer it
    var a int = 10
    var b = 15.5      // type inference (deduction)
    c := "Gopher"     // short declaration, type inference
    _, _, _ = a, b, c // Blank Identifier (_) to get rid of unused variable error

    // Go is a Statically and Strong Typed Programming Language
    // a = 3.14 -> error. A variable cannot change it's type
    // a = b    -> error. It's not allowed to assign a type to another type

    //** ZERO VALUES **//

    // An uninitialized variable or empty variable  will get the so called ZERO VALUE
    // The zero-value mechanism of Go ensures that a variable always holds a well defined value of its type
    var value int                         // initialized with 0
    var price float64                     // initialized with 0.0
    var name string                       // initialized with empty string -> ""
    var done bool                         // initialized with false
    fmt.Println(value, price, name, done) // -> 0 0.0 ""  false
}
```

### Naming conventions

[Go specifications](https://go.dev/ref/spec#Keywords)

```go
package main
 
//** COMMENTS **//
 
// this is a single line comment
 
/*
 This is a block comment.
 a := 10
 fmt.Println(a)
*/
 
var name = "John Wick" // inline comment
 
//** NAMING CONVENTIONS IN GO **//
 
// Naming Conventions are important for code readability and maintainability.
 
// use short, concise names especially in shorter scopes
// common names for common types:
var s string   //string
var i int      //index
var num int    //number
var msg string //message
var v string   //value
var err error  //error value
var done bool  //bool, has been done?
 
// use mixedCase a.k.a camelCase instead of snake_case (variables and  functions)
var maxValue = 100  // recommended (camelCase)
var max_value = 100 // not recommended (snake_case)
 
// recommended
func writeToFile() {
}
 
// not recommended
func write_to_file() {
}
 
// write acronyms in all caps
var writeToDB = true // recommended
var writeToDb = true // not recommended
 
func main() {
 
    // use fewer letters, don’t be too verbose especially in smaller scopes
    var packetsReceived int // NOT OK, too verbose
    var n int               // OK
    _, _ = packetsReceived, n
 
    // an uppercase first letter has special significance to go (it will be exported in other packages)
}
```

### Package fmt

[fmt](https://pkg.go.dev/fmt#Printf)

```go
package main
 
// Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
// It's used mainly to print out to stdout
import "fmt"
 
func main() {
 
    // fmt.Println() writes to standard output.
    // spaces are always added between operands and a newline is appended.
    fmt.Println("Hello Go World!") // => Hello Go World!
 
    var name, age = "Andrei", 35
    fmt.Println(name, "is", age, "years old.") // => Andrei is 35 years old.
 
    //** fmt.Printf() **//
 
    // fmt.Printf() prints out to stdout according to a format specifier called verb.
    // It doesn't add a newline (\n)
 
    // VERBS:
    // %d -> decimal
    // %f -> float
    // %s -> string
    // %q -> double-quoted string
    // %v -> value (any)
    // %#v -> a Go-syntax representation of the value
    // %T -> value Type
    // %t -> bool (true or false)
    // %p -> pointer (address in base 16, with leading 0x)
    // %c -> char (rune) represented by the corresponding Unicode code point
 
    a, b, c := 10, 15.5, "Gophers"
    grades := []int{10, 20, 30}
 
    fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c)    // => a is 10, b is 15.500000, c is Gophers
    fmt.Printf("%q\n", c)                      // => "Gophers"
    fmt.Printf("%v\n", grades)                 // => [10 20 30]
    fmt.Printf("%#v\n", grades)                // => b is of type float64 and grades is of type []int
    fmt.Printf("b is of type %T and grades is of type %T\n", b, grades) 
    // => b is of type float64 and grades is of type []int
    fmt.Printf("The address of a: %p\n", &a)    // => The address of a: 0xc000016128
    fmt.Printf("%c and %c\n", 100, 51011)       // =>  d and 읃  (runes for code points 101 and 51011)
 
    const pi float64 = 3.14159265359
    fmt.Printf("pi is %.4f\n", pi) // => formatting with 4 decimal points
 
    // %b -> base 2
    // %x -> base 16
    fmt.Printf("255 in base 2 is %b\n", 255)  //  => 255 in base 2 is 11111111
    fmt.Printf("101 in base 16 is %x\n", 101) // => 101 in base 16 is 65
 
    // fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()
    s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)
    fmt.Println(s) // => a is 10, b is 15.500000, c is Gophers
}
```

### Constants

```go
package main

import "fmt"

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi = 3.14
	fmt.Println(days, pi)

	const a, b = 5, 0
	// fmt.Println(a/b) -> error detected
	fmt.Println(a, b)

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)
	fmt.Println(min1, min2, min3)

	// group constant
	const (
		max1 = 100
		max2
		max3
	)
	fmt.Println(max1, max2, max3)

}
```

### Untype constants

[Constants](https://go.dev/blog/constants)

```go
package main

import "fmt"

func main() {
	// typed constant
	const a float64 = 7.7
	// untyped constant
	const b = 6.6

	const c = a * b
	const name = "Mario" + "Lazzari"

	fmt.Println(c, name)

	// const x int = 5
	// const y float64 = 7 * x -> error: type mismatched

	const x = 5
	const y = x * 2.2

	fmt.Printf("%T\n", y)

	const i int = x     // untyped x changes to int
	const j float64 = x // untyped x changes to float
	fmt.Printf("%T %T\n", i, j)

}
```

### Iota

```go
package main
 
import "fmt"
 
func main() {
    // To declare a constant and give it a name, we use the const keyword
    // Constants need to be initialized when declared
    const days int = 7 // typed constant
    const pi = 3.14    // untyped constant
    
    // There are ONLY boolean constants, rune constants, integer constants, 
    // floating-point constants, complex constants, and string constants.
 
    // Declaring multiple (grouped) constants
    const (
        a         = 5   // untyped constant
        b float64 = 0.1 // typed constant
    )
 
    const n, m int = 4, 5
 
    const (
        min1 = -500
        max1 //gets its type and value form the previous constant. It's 500
        max2 //in a grouped constants, a constant repeats the previous one -> 500
    )
 
    // CONSTANTS RULES
 
    // 1. You cannot change a constant
    const temp int = 100
    // temp = 50 //compile-time error
 
    // 2. You cannot initiate a constant at runtime (constants belong to compile-time)
    // const power = math.Pow(2, 3) //error, functions calls belong to runtime
 
    // 3. You cannot use a variable to initialize a constant
    t := 5
    // error, variables belong to runtime and you cannot initialize a const to runtime values
    // const tc = t 
    
 
    // 4. You can use a function like len() to initialize a const if it has as argument
    // a constant string literal (not a variable)
    // a string literal is an untyped constant
 
    const l1 = len("Hello") //ok
 
    str := "Hello"
    // const l2 = len(str) //error, str is a variable and belongs to runtime
 
    _, _ = t, str
 
    // UNTYPED CONSTANTS
    const x = 5
    const y float64 = 1.1
 
    var v1 int = 5
    var v2 float64 = 1.1
 
    fmt.Println(x * y) 
    // => 5.5, No Error because x is untyped and gets its type when its used first time (float64).
    
    // fmt.Println(v1 * v2)  
    // => Error: invalid operation: v1 * v2 (mismatched types int and float64)
    _, _ = v1, v2
 
    // IOTA
    // iota is number generator for constants which starts from zero 
    // and is incremented by 1 automatically.
 
    const (
        c1 = iota
        c2 = iota
        c3 = iota
    )
    fmt.Println(c1, c2, c3) // => 0 1 2
 
    const (
        North = iota //by default 0
        East         //omitting type and value means, repeating its type and value so East = iota = 1 (it increments by 1 automatically)
        South        // -> 2
        West         // -> 3
    )
 
    // Initializing the constants using a step:
    const (
        c11 = iota * 2 // -> 0
        c22            // -> 2
        c33            // -> 4
    )
}
```

### Types

```go
package main
 
import "fmt"
 
func main() {
 
    //** NUMERIC TYPES **//
 
    // uint8      the set of all unsigned  8-bit integers (0 to 255)
    // uint16      the set of all unsigned 16-bit integers (0 to 65535)
    // uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
    // uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
 
    // int8        the set of all signed  8-bit integers (-128 to 127)
    // int16       the set of all signed 16-bit integers (-32768 to 32767)
    // int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
    // int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
 
    // uint     either 32 or 64 bits
    // int      same size as uint
 
    // float32     the set of all IEEE-754 32-bit floating-point numbers
    // float64     the set of all IEEE-754 64-bit floating-point numbers
    // complex64   the set of all complex numbers with float32 real and imaginary parts
    // complex128  the set of all complex numbers with float64 real and imaginary parts
 
    // byte        alias for uint8
    // rune        alias for int32
 
    //int type
    var i1 int8 = -128     //min value
    fmt.Printf("%T\n", i1) // => int8
 
    var i2 uint16 = 65535  //max value
    fmt.Printf("%T\n", i2) // => int16
 
    var i3 int64 = -324_567_345  // underscores are used to write large numbers for a better readability
    fmt.Printf("%T\n", i3)       // => int64
    fmt.Printf("i3 is %d\n", i3) // => i3 is -324567345 (underscores are ignored)
 
    //float64 type
    var f1, f2, f3 float64 = 1.1, -.2, 5. // trailing and leading zeros can be ignored
    fmt.Printf("%T %T %T\n", f1, f2, f3)
 
    //rune type
    var r rune = 'f'
    fmt.Printf("%T\n", r) // => int32 (rune is an alias to int32)
    fmt.Printf("%x\n", r) // => 66,  the hexadecimal ascii code for 'f'
    fmt.Printf("%c\n", r) // => f
 
    //bool type
    var b bool = true
    fmt.Printf("%T\n", b) // => bool
 
    //string type
    var s string = "Hello Go!"
    fmt.Printf("%T\n", s) // => string
 
    //array type
    var numbers = [4]int{4, 5, -9, 100}
    fmt.Printf("%T\n", numbers) // =>  [4]int
 
    //slice type
    var cities = []string{"London", "Bucharest", "Tokyo", "New York"}
    fmt.Printf("%T\n", cities) // => []string
 
    //map type
    balances := map[string]float64{
        "USD": 233.11,
        "EUR": 555.11,
    }
    fmt.Printf("%T\n", balances) // => map[string]float64
 
    //struct type
    type Person struct {
        name string
        age  int
    }
    var you Person
    fmt.Printf("%T\n", you) // => main.Person
 
    //pointer type
    var x int = 2
    ptr := &x                                                 // pointer to int
    fmt.Printf("ptr is of type %T with value %v\n", ptr, ptr) // => ptr is of type *int with value 0xc000016168
 
    //function type
    fmt.Printf("%T\n", f) // => func()
 
}
 
func f() {
}
```

### Operators

```go
package main
 
import "fmt"
 
func main() {
    a, b := 10, 5.5
 
    //** ARITHMETIC OPERATORS **//
    //  +       sum
    // -        difference
    // *        product
    // /        quotient
    // %        remainder
    // there is no power operator in Go. Use math.Pow(a, b) for raising to a power.
 
    fmt.Println(a + 5)   // => 15
    fmt.Println(3.1 - b) // => -2.4
    fmt.Println(a * a)   // => 100
    fmt.Println(a / a)   // => 1
    fmt.Println(11 / 5)  // => 2
 
    // Go is a Strong Typed Language
    // fmt.Println(a * b)       // =>  invalid operation: a * b (mismatched types int and float64)
    fmt.Println(a * int(b))     // => 50
    fmt.Println(float64(a) * b) // => 55
 
    // IncDec Statements
    // The "++" and "--" statements increment or decrement their operands by the untyped constant 1.
    x := 10
    x++ // x is 11. Same as: x += 1
    x-- // x is 10. Same as: x -= 1
 
    //** ASSIGNMENT OPERATORS **//
    //  =   (simple assignment)
    // +=   (increment assignment)
    // -=   (decrement assignment)
    // *=   (multiplication assignment)
    // /=   (division assignment)
    // %=   (modulus assignment)
 
    a = 10
    a += 2 // => a is 12
    a -= 3 // => a is 9
    a *= 2 // => a is 18
    a /= 3 // => a is 6
    a %= 5 // => a is 1
 
    //** COMPARISON OPERATORS **//
    //  ==      equal values
    // !=       not equal
    // >        left operand is greater than right operand
    // <        left operand is less than right operand
    // >=       left operand is greater than or equal to right operand
    // <=       left operand is less than or equal to right operand
 
    fmt.Println(5 == 6)   // => false
    fmt.Println(5 != 6)   // => true
    fmt.Println(10 > 10)  // => false
    fmt.Println(10 >= 10) // => true
    fmt.Println(5 < 5)    // => false
    fmt.Println(5 <= 5)   // => true
 
    //** LOGICAL OPERATORS **//
    // &&       logical and
    // ||       logical or
    // !        logical negation
 
    fmt.Println(0 < 2 && 4 > 1) // => true
    fmt.Println(1 > 5 || 4 > 5) // => false
    fmt.Println(!(1 > 2))       // => true
 
}
```
