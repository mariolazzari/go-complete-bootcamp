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

### Overflows

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	x++ // overflow
	fmt.Println(x)

	// a := int8(255 + 1) -> compile time overflow

	var b int8 = 127
	fmt.Printf("%d\n", b+1)

	b = -127
	b--
	fmt.Printf("%d\n", b)

	f := float32(math.MaxFloat32)
	fmt.Println(f)
}
```

### Converting types

```go
package main
 
import (
    "fmt"
    "strconv"
)
 
func main() {
 
    var x = 3   //int type
    var y = 3.2 //float type
 
    // x = x * y //compile error ->  mismatched types
 
    x = x * int(y) // converting float64 to int
    fmt.Println(x) // => 9
 
    y = float64(x) * y //converting int to float64
    fmt.Println(y)     // => 28.8
 
    x = int(float64(x) * y)
    fmt.Println(x) // => 259
 
    //In Go types with different names are different types.
    var a int = 5   // same size as int64 or int32 (platform specific)
    var b int64 = 2 // int and int64 are not the same type
 
    // a = b // error: cannot use b (type int64) as type int in assignment
    a = int(b) // converting int64 to int (explicit conversion required)
 
    // preventing unused variable error
    _ = a
 
    //** CONVERTING NUMBERS TO STRINGS AND STRINGS TO NUMBERS **//
 
    s := string(99)            // int to rune (Unicode code point)
    fmt.Println(s)             // => 99, the ascii code for symbol c
    fmt.Println(string(34234)) // => 34234 is the unicode code point for 薺
 
    // we cannot convert a float to a string similar to an int to a string
    // s1 := string(65.1) // error
 
    // converting float to string
    var myStr = fmt.Sprintf("%f", 5.12)
    fmt.Println(myStr) // => 5.120000
 
    // converting int to string
    var myStr1 = fmt.Sprintf("%d", 34234)
    fmt.Println(myStr1) // => 34234
 
    // converting string to float
    var result, err = strconv.ParseFloat("3.142", 64)
    if err == nil {
        fmt.Printf("Type: %T, Value: %v\n", result, result) // => Type: float64, Value: 3.142
    } else {
        fmt.Println("Cannot convert to float64!")
    }
 
    // Atoi(string to int) and Itoa(int to string).
    i, err := strconv.Atoi("-50")
    s = strconv.Itoa(20)
    fmt.Printf("i Type is %T, i value is %v\n", i, i) // => i Type is int, i value is -50
    fmt.Printf("s Type is %T, s value is %q\n", s, s) // => s Type is string, s value is "20"
}
```

### Named types

```go
package main
 
import "fmt"
 
type age int        //new type, int is the underlying type
type oldAge age     //new type, int (not age) is the underlying type
type veryOldAge age //new type, int (not age) is the underlying type
 
func main() {
 
    // new type speed (underlying type uint)
    type speed uint
 
    // s1, s2 of type speed
    var s1 speed = 10
    var s2 speed = 20
 
    // performing operations with the new types
    fmt.Println(s2 - s1) // -> 10
 
    // uint and speed are different types (they have different names)
    var x uint
 
    // x = s1  //error different types
 
    // correct
    x = uint(s1)
    _ = x
 
    // correct
    var s3 speed = speed(x)
    _ = s3
}
```

### Aliases

```go
package main
 
import "fmt"
 
func main() {
 
    // declaring a variable of type uint8
    var a uint8 = 10
    var b byte // byte is an alias to uit8
 
    // even though they have different names, byte and uit8 are the same type because they are aliases
    b = a // no error
    _ = b
 
    // declaring a new alias named second for uint
    // type alias_name = type_name
    type second = uint
 
    var hour second = 3600
    fmt.Printf("hour type: %T\n", hour) // => hour type: uint
 
    //no need to convert operations (same type)
    fmt.Printf("Minutes in an hour: %d\n", hour/60) // => Minutes in an hour: 60
}
```

## Coding challeges: basics

### Declare variables

#### Declare variables: ex1

```go
package main

import "fmt"

func main() {
	// Using the var keyword, declare 4 variables called a, b, c, d of type int, float64, bool and string.
	var a int
	var b float64
	var c bool
	var d string

	// Using short declaration syntax declare and assign these values to variables x, y and z:
	x := 20
	y := 15.5
	z := "Gopher!"

	// Using fmt.Println() print out the values of a, b, c, d, x, y and z.
	fmt.Println(a, b, c, d, y, x, z)
}
```

#### Declare variables: ex2

```go
package main

func main() {
	// Declare a, b, c, d using a single var keyword (multiple variable declaration) for better readability.
	var (
		a int
		b float64
		c bool
		d string
	)

	// Declare x, y and z on a single line -> multiple short declarations
	x, y, z := 20, 15.5, "Gopher!"

	// Remove the statement that prints out the variables. See the error!
	//fmt.Println(a, b, c, d, y, x, z)

	// Change the program to run without error using the blank identifier
	_ = a
	_ = b
	_ = c
	_ = d
	_ = x
	_ = y
	_ = z
}
```

#### Declare variables: ex3

```go
package main

func main() {
	var a float64 = 7.1

	x, y := true, 3.7

	// Try to identify the errors, change the code and run the program without errors.
	// a, x := 5.5, false
	a, x = 5.5, false

	_, _, _ = a, x, y
}
```

#### Declare variables: ex4

```go
package main

import "fmt"

// Try to identify the errors, change the code and run the program without errors.
// version := "3.1"
var version = "3.1"

func main() {
	//name := 'Golang'
	name := "Golang"
	fmt.Println(name)
}
```

### Constants

#### Constants ex1

```go
package main

func main() {
	// Using the const keyword declare and initialize the following constants:

	// 1. daysWeek with value 7
	const daysWeek = 7

	// 2. lightSpeed with value 299792458
	const lightSpeed = 299792458

	// 3. pi with value 3.14159
	const pi = 3.14159
}
````

#### Constants ex2

```go
package main

func main() {
	// Change the code from the previous exercise and declare all 3 constants as grouped constants.
	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)
}
```

#### Constants ex3

```go
package main

import "fmt"

// Calculate how many seconds are in a year.
func main() {
	// Declare secPerDay constant and initialize it to the number of seconds in a day
	const secPerDay = 60 * 60 * 24
	// Declare daysYear constant and initialize it to 365
	const daysYear = 365
	// Use fmt.Printf() to print out the total number of seconds in a year.
	fmt.Printf("%d\n", secPerDay*daysYear)
}
```

#### Constants ex4

```go
package main

func main() {
	const x int = 10

	// declaring a constant of type slice int ([]int)
	// const m = []int{1: 3, 4: 5, 6: 8}
	m := []int{1, 3, 4, 5, 6, 8}

	_ = m
}
```

#### Constants ex5

```go
package main

import "math"

// Try to identify the errors,
func main() {
	const a int = 7
	const b float64 = 5.6
	// const c = a * b
	const c = float64(a) * b

	// x := 8
	const x = 8
	const xc int = x

	// const noIPv6 = math.Pow(2, 128)
	noIPv6 := math.Pow(2, 128)
	_ = noIPv6
}
```

#### Constants ex6

```go
package main

// Using Iota declare the following months of the year: Jun, Jul and Aug
func main() {
	//Jun, Jul and Aug are constant and their value is 6, 7 and 8
	const (
		Jun = iota + 6
		Jul
		Aug
	)
}
```

### Fmt package

#### Fmt ex1

```go
package main

import "fmt"

// Write a Go program that converts i to float64 and f to int.
func main() {
	var i int = 3
	var f float64 = 3.2

	iF := float64(i)
	fI := int(f)

	// Print out the type and the value of the variables created after conversion.
	fmt.Printf("i type: %T, f type: %T\n", iF, fI)
}
```

#### Fmt ex2

```go
package main

import (
	"fmt"
	"strconv"
)

// Write a Go program that converts
// // Print the value and the type for each variable created after conversion.
func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// i to string (int to string)
	s := strconv.Itoa(i)
	fmt.Printf("s Type is %T, s value is %q\n", s, s)

	// s2 to int (string to int)
	is, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is)
	} else {
		fmt.Println("Can not convert string to int.")
	}

	// f to string (float64 to string)
	ss1 := fmt.Sprintf("%f", f)
	fmt.Printf("ss1's type: %T, ss1's value: %s\n", ss1, ss1)

	// s1 to float64  (string to float64)
	f1, err1 := strconv.ParseFloat(s1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1)
	} else {
		fmt.Println("Cannot convert string to float64.")
	}
}
```

#### Fmt ex3

```go
package main

import "fmt"

// There are some errors i
func main() {
	x, y := 4, 5.1

	// z := x * y
	z := float64(x) * y
	fmt.Println(z)

	a := 5
	// b := 6.2 * a
	b := 6.2 * float64(a)
	fmt.Println(b)
}
```

### Named types

#### Named types ex1

```go
package main

import "fmt"

// Declare a new type type called duration. Have the underlying type be an int.
type duration int

// Declare a variable of the new type called hour using the var keyword
var hour duration

func main() {
	// print out the value of the variable hour
	fmt.Printf("%d\n", hour)
	// print out the type of the variable hour
	fmt.Printf("%T\n", hour)
	// assign 3600 to the variable hour using the  = operator
	hour = 3600
	// print out the value of hour
	fmt.Printf("%d\n", hour)
}
```

#### Named types ex2

```go
package main

import "fmt"

type duration int

// There are some errors in the following
func main() {
	var hour duration = 3600
	minute := 60
	// fmt.Println(hour != minute)
	fmt.Println(hour != duration(minute))
}
```

#### Named types ex3

```go
package main

import "fmt"

// Declare two defined types called mile and kilometer.
// Have the underlying type be an float64.
type mile float64
type kilometer float64

// Declare a constant called m2km equals 1.609  ( 1mile=1.609km)
const m2km = 1.609

func main() {
	// declare a variable called mileBerlinToParis of type mile with value 655.3
	var mileBerlinToParis mile = 6555.3

	// declare a variable called kmBerlinToParis of type kilometer
	var kmBerlinToParis kilometer

	// calculate the distance between Berlin and Paris in km by multiplying mileBerlinToParis and m2km.
	// Assign the result to kmBerlinToParis
	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)

	// print out the distance in km between Berlin and Paris
	fmt.Println(kmBerlinToParis)
}
```

## Program flow

### if

```go
package main
 
import "fmt"
 
func main() {
 
    // if condition_that_evaluates_to_boolean{
    //      perform action1
    // }else if condition_that_evaluates_to_boolean{
    //      perform action2
    // }else{
    //      perform action3
    // }
 
    price, inStock := 100, true
 
    if price >= 80 { // parenthesis are no required to enclose the testing condition
        fmt.Println("Too Expensive")
    }
 
    if price <= 100 && inStock == true { //the same with: if price <= 100 && inStock { }
        fmt.Println("Buy it!")
    }
 
    // In Go there is not such a thing like the Truthiness of a variable.
    // Error:
    // if price {
    //  fmt.Println("We have price!")
    // }
 
    // only one if branch will be executed
    if price < 100 {
        fmt.Println("It's cheap!")
    } else if price == 100 {
        fmt.Println("On the edge")
    } else { //executed only once if all the if branches are false (it's optional)
        fmt.Println("It's Expensive!")
    }
}
```

### Command line args

```go
package main
 
import (
    "fmt"
    "os"
)
 
func main() {
    fmt.Println("os.Args:", os.Args) // os.Args is slice of strings ([]string)
 
    // accessing command line arguments using indexes
    fmt.Println("Path:", os.Args[0])
    fmt.Println("1st Argument:", os.Args[1])
    fmt.Println("2nd Argument:", os.Args[2])
    fmt.Println("No. of items inside os.Args:", len(os.Args))
}
```

### Simple if

```go
package main
 
import (
    "fmt"
    "strconv"
)
 
func main() {
    // converting string to int:
    i, err := strconv.Atoi("45")
 
    // error handling
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(i)
 
    }
 
    // simple (short) statement ->  the same effect as the above code
    // i and err are variables scoped to the if statement only
    if i, err := strconv.Atoi("34"); err == nil {
        fmt.Println("No error. i is ", i)
    } else {
        fmt.Println(err)
    }
}
```

### For loops

```go
package main
 
import "fmt"
 
func main() {
 
    // printing numbers from 0 to 9
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
 
    // has the same effect as a while loop in other languages
    // there is no while loop in Go
    j := 10
    for j >= 0 {
        fmt.Println(j)
        j--
    }
 
    // handling of multiple variables in a for loop
    for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
        fmt.Printf("i = %v, j = %v\n", i, j)
    }
 
    // infinite loop
    // sum := 0
    // for {
    //  sum++
    // }
    // fmt.Println(sum) //this line is never reached
}
```

### Continue & break

```go
package main
 
import "fmt"
 
func main() {
 
    //** CONTINUE STATEMENT **//
 
    // It works just the same as in C,  Java or Python.
    // The continue statement rejects all the remaining statements in the current iteration of the loop
    // and moves the control back to the top of the loop.
 
 
    // printing even numbers less than or equal to 10
    for i := 1; i <= 10; i++ {
        if i%2 != 0 {
            continue    // skipping the remaining code in this iteration
        }
        fmt.Println(i)
    }
 
 
    // **BREAK STATEMENT **//
 
    // It is used to terminate the innermost for or switch statement.
    // It works just the same as in C,  Java or Python.
    
    // finding 10 numbers divisible by 13 
    count := 0 
    for i := 0; true; i++ {
        if i%13 == 0 {
            fmt.Printf("%d is divisible by 13\n", i)
            count++
        }
 
        if count == 10 { //if 10 numbers were found, break!
            break //it breaks the current loop (inner loop if there are more loops)
        }
    }
 
    // the break statement is not terminating the program entirely;
    fmt.Println("Just a message after the for loop")
}
```

### Labels

```go
package main
 
import (
    "fmt"
)
 
func main() {
    //** LABEL STATEMENT **//
 
    // declaring a variable
    // there is no conflict name between variable and label
    outer := 19  
    _ = outer
 
    // declaring two arrays
    people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Michael"}
    friends := [2]string{"Mark", "Marry"}
 
    // searching for a single friend in a list of people.
    
outer: //label, it doesn't conflict with other names
    // iterating over the array.
    for index, name := range people {  // range returns both the index and the elements of the array one by one
        for _, friend := range friends { //iterating over the second array
            if name == friend {
                fmt.Printf("FOUND A FRIEND: %q at index %d\n", friend, index)
                break outer //breaking outside the outer loop which terminates
            }
        }
    }
 
    fmt.Println("Next instruction after the break.")
 
 
    // **GOTO STATEMENT **//
    
    //the following piece of code creates a loop like a for statement does
    i := 0
loop: // label
    if i < 5 {
        fmt.Println(i)
        i++
        goto loop
    }
 
    //  goto todo //ERROR it's not permitted to jump over the declaration of x
    //  x := 5
    // todo:
    //  fmt.Println("something here")
}
```

### Switch

```go
package main
 
import "fmt"
 
func main() {
 
    language := "golang"
 
    switch language {
    case "Python": //values must be comparable (compare string to string)
        fmt.Println("You are learning Python! You don't use { } but indentation !! ")
        // an implicit break is added here
    case "Go", "golang": //compare language with "Go" OR "golang"
        fmt.Println("Good, Go for Go!. You are using {}!")
    default:
        // the default clause the equivalent of the else clause of an if statement
        // and gets executed if no testing condition is true.
        fmt.Println("Any other programming language is a good start!")
    }
 
    n := 5
    // comparing the result of an expression which is bool to another bool value
    switch true {
    case n%2 == 0:
        fmt.Println("Even!")
    case n%2 != 0:
        fmt.Println("Odd!")
    default:
        fmt.Println("Never here!")
    }
 
    //** Switch simple statement **//
 
    // Syntax: statement (n:=10), semicolon and a switch condition
    //(true in this case, we are comparing boolean expressions that return true)
    // we can remove the word "true" because it's the default
    switch n := 10; true {
    case n > 0:
        fmt.Println("Positive")
    case n < 0:
        fmt.Println("Negative")
    default:
        fmt.Println("Zero")
    }
}
```

### Scopes

```go
package main
 
// import statements are file scoped
import (
    "fmt"
 
    // import "fmt" -> error, within the same scope, unique names
 
    // importing as another name (alias) is permitted
    f "fmt"
)
 
// variables or constant declared outside any function are package scoped
const done = false
 
func main() { // package scoped
 
    // block scoped: visible until the end of the block "}"
    var task = "Running:"
    fmt.Println(task, done) // => Running: false (this is done from package scope)
    f.Println("Bye bye!")
 
    // names must be unique only within the same scope
    const done = true                        // local scoped
    fmt.Printf("done in main(): %v\n", done) // => done in main(): true
    f1()
}
 
func f1() {
    fmt.Printf("done in f(): %v\n", done) //this is done from package scope
}
```

## Flow exercises

### Flow ex1

```go
package main

import "fmt"

// Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.
func main() {
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}
```

### Flow ex2

```go
package main

import "fmt"

// Change the code from the previous exercise and use the continue statement to print out all the numbers divisible by 7 between 1 and 50.
func main() {
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
```

### Flow ex3

```go
package main

import "fmt"

// Change the code from the previous exercise and use the break statement to print out only the first 3 numbers divisible by 7 between 1 and 50.
func main() {
	j := 1
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i)

		j++
		if j == 4 {
			break
		}
	}
}
```

### Flow ex4

```go
package main

import "fmt"

// Using a for loop, an if statement and the logical and operator print out all the numbers between 1 and 500 that divisible both by 7 and 5.
func main() {
	for i := 1; i <= 50; i++ {
		if i%5 != 0 && i%7 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
```

### Flow ex5

```go
package main

import (
	"fmt"
	"time"
)

// Using a for loop print out all the years from your birthday to the current year.
// Use a variant of for loop where the post statement is moved inside the for block of code.
func main() {
	year := 1975
	for year <= time.Now().Year() {
		fmt.Println(year)
		year++
	}
}
```

### Flow ex6

```go
package main

import "fmt"

func main() {
	age := -9
	if age < 0 || age > 100 {
		fmt.Println("Invalid Age")
	} else if age < 18 {
		fmt.Println("You are minor!")
	} else if age == 18 {
		fmt.Println("Congratulations! You've just become major!")
	} else {
		fmt.Println("You are major!")
	}

	// Change the code and use a switch statement instead of an if statement.
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
```

## Arrays
