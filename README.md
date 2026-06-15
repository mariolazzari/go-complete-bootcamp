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
```

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

### Intro to arrays

```go
package main

import (
    "fmt"
)

func main() {

    // declaring an array with four elements of type int
    var numbers [4]int

    // array zero value is zeroed value elements
    fmt.Printf("%v\n", numbers)  // -> [0 0 0 0]
    fmt.Printf("%#v\n", numbers) // -> [4]int{0, 0, 0, 0}

    // declaring an array and initialize it using an array literal
    var a1 = [4]float64{}                           //initialized with defaults (0)
    var a2 = [3]int{5, -3, 5}                       //initialized with the given values
    a3 := [4]string{"Dan", "Diana", "Paul", "John"} //short declaration syntax
    a4 := [4]string{"x", "y"}                       //initializing only the first 2 elements

    // the ellipsis operator (...) finds out automatically the length of the array
    a5 := [...]int{1, 4, 5}
    fmt.Println("The length of a5 is: ", len(a5)) // len is 3

    // declare an array on multiple lines for better readability
    a6 := [...]int{1,
        2,
        3,
    } //the ending comma is mandatory when initializing the array on multiple lines and the closing curly brace is on its own line

    _, _, _, _, _, _ = a1, a2, a3, a4, a5, a6

    // changing an array
    // we can't add or remove elements from the array (it's fixed length)
    numbers[0] = 7              //changing first element (index 0)
    fmt.Printf("%v\n", numbers) // -> [7 0 0 0]

    // compile-time error
    // numbers[5] = 8  // invalid array index 5 (out of bounds for 4-element array)

    // getting an element
    x := numbers[0]
    fmt.Println("x is ", x) // => x is  7

    // iterating over an array (2-ways)
    for i, v := range numbers { // range is a language keyword used for iteration
        fmt.Println("index:", i, "value: ", v)

    }

    // iterating over an array (C/C++, Java Style)
    for i := 0; i < len(numbers); i++ {
        fmt.Println("index:", i, "value: ", numbers[i])
    }

    // declaring a multi-dimensional arrays (array of arrays or matrix)
    balances := [2][3]int{
        [3]int{5, 6, 7}, //[3]int is optional
        {8, 9, 10},
    }

    for _, arr := range balances {
        for _, value := range arr {
            fmt.Printf("%d ", value)
        }
        fmt.Println("")
    }

    //  = operator creates a copy of an array.
    // the arrays are not connected and are saved in different memory locations
    m := [3]int{1, 2, 3}
    n := m //n is a copy of m

    fmt.Println("n is equal to m: ", n == m) // => true
    m[0] = -1                                //only m is modified
    fmt.Println("n is equal to m: ", n == m) // => false

}
```

### Keyed elements

```go
package main

import "fmt"

func main() {

    // each key corresponds to an index of the array
    grades := [3]int{ //the keyed elements can be in any order
        1: 10,
        0: 5,
        2: 7,
    }
    fmt.Println(grades) // -> [5 10 7]

    accounts := [3]int{
        2: 50,
    }
    fmt.Println(accounts) //[0 0 50]

    names := [...]string{
        4: "Dan",
    }
    fmt.Println(len(names)) // -> 5

    // un unkeyed element gets its index from the last keyed element
    cities := [...]string{
        5:        "Paris",
        "London", // this is at index 6
        1:        "NYC",
    }
    fmt.Printf("%#v\n", cities) // -> [7]string{"", "NYC", "", "", "", "Paris", "London"}

    weekend := [7]bool{5: true, 6: true}
    fmt.Println(weekend) // => [false false false false false true true]
}
```

### Arrays exercises

#### Array ex1

```go
package main

import "fmt"

func main() {
	// Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array.
	var cities [2]string
	// Print out the cities array and notice the value of its elements.
	fmt.Println(cities)

	// Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
	var grades [3]float64 = [3]float64{10, 20}
	// Print out the grades array and notice the value of its elements.
	fmt.Println(grades)

	// Declare an array called taskDone using the ellipsis operator. The elements are of type bool. Print out taskDone.
	var taskDone = [...]bool{true, false, true}
	fmt.Println(taskDone)

	// Iterate over the array called cities using the classical for loop syntax and the len function and print out element by element.
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	// Iterate over grades using the range keyword and print out element by element.
	for _, grade := range grades {
		fmt.Println(grade)
	}
}
```

#### Array ex2

```go
package main

import "fmt"

func main() {
	// Consider the following array declaration:
	nums := [...]int{30, -1, -6, 90, -6}

	// Write a Go program that counts the number of positive even numbers in the array.
	count := 0
	for _, n := range nums {
		if n > 0 && n%2 == 0 {
			count++
		}
	}
	fmt.Printf("Positive even numbers: %d\n", count)
}
```

#### Array ex3

```go
package main

import "fmt"

func main() {
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10

	// There are some errors in the following
	// myArray[0] = a
	myArray[0] = float64(a)
	// myArray[3] = 10.10
	myArray[2] = 10.10

	fmt.Println(myArray)
}
```

## Slices

### Slices intro

```go
package main

import "fmt"

func main() {

	// declaring a string slice, by default is initialized with nil or uninitialized
	var cities []string

	fmt.Println("cities is equal to nil: ", cities == nil) // -> cities is equal to nil:  true
	fmt.Printf("cities: %#v\n", cities)                    // -> cities: []string(nil)

	// we can not assign elements to nil slice:
	// cities[0] = "Paris" // -> runtime error

	// declaring a slice using a slice literal
	numbers := []int{2, 3, 4, 5} // on the right hand-side of the equal sign is a slice literal
	fmt.Println(numbers)         // => [2 3 4 5]

	// creating a slice using the make() built-in function
	// creating a slice with 2 int elements initialized with zero.
	nums := make([]int, 2) // the same as []int{0, 0}.
	fmt.Println(nums)      // => [0 0]

	// declaring a slice using a slice literal
	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	// getting an element from the slice
	x := numbers[0]
	fmt.Println("x is", x) // => x is 2

	// modifying an element of the slice
	numbers[1] = 200
	fmt.Printf("%#v\n", numbers) // => []int{2, 200, 4, 5}

	// iterating over a slice
	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	//iterating over a slice (C/C++, Java Style)
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index: %v, value: %v\n", i, numbers[i])

	}

	// slices with the same element type can be assigned to each other
	var n []int
	n = numbers
	_ = n

	//** COMPARING SLICES **//
	// a Go slice can only be compared to nil

	// uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil) // true

	// empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil) //false

	// we can not compare slices using the equal (=) operator
	// fmt.Println(nn == mm) //error -> slice can only be compared to nil

	// to compare 2 slices use a for loop to iterate over the slices and compare element by element
    // it's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)
    a, b := []int{1, 2, 3}, []int{1, 2, 3}
    var eq bool = true
	if len(a) != len(b) {
		eq = false
	}

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false // don't check further, break!
			break
		}
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}
}
```

### Append slices

```go
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Printf("%#v\n", nums)

	nums = append(nums, 4)
	fmt.Printf("%#v\n", nums)

	nums = append(nums, 5, 6, 7)
	fmt.Printf("%#v\n", nums)

	nums2 := []int{100, 200}
	nums = append(nums, nums2...)
	fmt.Printf("%#v\n", nums)

	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Printf("src = %#v, dst = %#v, nn = %d\n\n", src, dst, nn)

}
```

### Slice expresisons

```go
package main

import "fmt"

func main() {

    // arrays, slices and strings are sliceable
    // slicing doesn't modify the array or the slice, it returns a new one

    // declaring an [5]int array
    a := [5]int{1, 2, 3, 4, 5}

    // a slice expression is formed by specifying a start or a low bound and a stop or high bound like  a[start:stop].
    // this selects a range of elements which includes the element at index start, but excludes the element at index stop.

    // slicing an array returns a slice, not an array
    b := a[1:3]                                 // 1 is called start (included), 3 is called stop (excluded)
    fmt.Printf("Type: %T , Value: %#v\n", b, b) // => Type: []int , Value: []int{2, 3}

    // declaring a slice
    s1 := []int{1, 2, 3, 4, 5, 6}

    s2 := s1[1:3]   //start included, stop excluded
    fmt.Println(s2) //[2 3]

    //for convenience, any of the indexis may be omitted.
    // a missing low index defaults to zero; a missing high index defaults to the length of the sliced operand.
    fmt.Println(s1[2:])       // => [3 4 5 6], same as s1[2:len(s1)]
    fmt.Println(s1[:3])       // => [1 2 3], same as s1[0:3]
    fmt.Println(s1[:])        // => [1 2 3 4 5 6], same with s1[0:len(s1)]
    fmt.Println(s1[:len(s1)]) // => => [1 2 3 4 5 6], returns the entire slice
    // fmt.Println(s1[:45])   //panic: runtime error: slice bounds out of range

    s1 = append(s1[:4], 100) // adds 100 after index 4 (excluded)
    fmt.Println(s1)          // -> [1 2 3 4 100]

    s1 = append(s1[:4], 200) // overwrites the last element
    fmt.Println(s1)          // -> [1 2 3 4 200]
}
```

### Slices internals

[Slice intro](https://go.dev/blog/slices-intro)

```go
// Go implements a slice as data structure called Slice Header.
// Slice Header contains 3 fields:
// - the address of the backing array (pointer).
// - the length of the slice.  The built-in function len() returns it.
// - the capacity of the slice. The size of the backing array after the slice first element. cap() built-in function returns it.

// A nil slice doesn't have backing array, so all the fields in the slice header are equal to zero.

package main

import (
    "fmt"
)

func main() {

    // a slice expression doesn't create a new backing array. The original and the returned slice are connected!
    s1 := []int{10, 20, 30, 40, 50}
    s3, s4 := s1[0:2], s1[1:3] //s3, s4 share the same backing array with s1

    s3[1] = 600     // modifying the backing array so s1, s3 and s4 are in fact modified!!
    fmt.Println(s1) // -> [10 600 30 40 50]
    fmt.Println(s4) // -> [600 30]

    // when a slice is created by slicing an array, that array becomes the backing array of the new slice.
    arr1 := [4]int{10, 20, 30, 40}
    slice1, slice2 := arr1[0:2], arr1[1:3]
    arr1[1] = 2                       // modifying the array
    fmt.Println(arr1, slice1, slice2) // -> [10 2 30 40] [10 2] [2 30]

    // append() function creates a complete new slice from an existing slice
    cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
    newCars := []string{}

    // newCars doesn't share the same backing array with cars
    newCars = append(newCars, cars[0:2]...)

    cars[0] = "Nissan"                              // only cars is modified
    fmt.Println("cars:", cars, "newCars:", newCars) // => cars: [Nissan Honda Audi Range Rover] newCars: [Ford Honda

}
```

### Length and capacity

[Slices](https://go.dev/blog/slices)

```go
package main

import "fmt"

func main() {
	numbers := []int{2, 3}

	// append() returns a new slice after appending a value to its end
	numbers = append(numbers, 10)
	fmt.Println(numbers) //-> [2 3 10]

	// appending more elements at once
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers) //-> [2 3 10 20 30 40]

	// appending all elements of a slice to another slice
	n := []int{100, 200, 300}
	numbers = append(numbers, n...) // ... is the ellipsis operator
	fmt.Println(numbers)            // -> [2 3 10 20 30 40 100 200 300]

	//** Slice's Length and Capacity **//

	nums := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 1, Capacity: 1

	nums = append(nums, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 2, Capacity: 2

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 3, Capacity: 4
	// the capacity of the new backing array is now larger than the length
	// to avoid creating a new backing array when the next append() is called.

	nums = append(nums, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 5, Capacity: 8

	// copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.
	// if the slices don't have the same no of elements, it copies the minimum of length of the slices
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(src, dst, nn) // => [10 20 30] [10 20 30] 3
}
```

## Slices exercises

### Slices ex1

```go
package main

import "fmt"

func main() {
	// Using a composite literal declare and initialize a slice of type string with 3 elements.
	names := []string{"Mario", "Maria", "Mariarosa"}
	// Iterate over the slice and print each element in the slice and its index.
	for i, name := range names {
		fmt.Printf("%d - %s\n", i, name)
	}
}
```

### Slice ex2

```go
package main

import "fmt"

func main() {
	mySlice := []float64{1.2, 5.6}

	// There are some errors
	mySlice[0] = 6

	// There are some errors
	a := float64(10)
	mySlice[0] = a

	// There are some errors
	// mySlice[3] = 10.10
	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)
}
```

### Slices ex3

```go
package main

import "fmt"

func main() {
	// 1. Declare a slice called nums with three float64 numbers.
	nums := []float64{1.1, 2.2, 3.3}
	// 2. Append the value 10.1 to the slice
	nums = append(nums, 1.1)
	// 3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
	nums = append(nums, 4.1, 5.5, 6.6)
	// 4. Print out the slice
	fmt.Printf("%#v\n", nums)
	// 5. Declare a slice called n with two float64 values
	n := []float64{7.7, 8.8}
	// 6. Append n to nums
	nums = append(nums, n...)
	// 7. Print out the nums slice
	fmt.Printf("%#v\n", nums)
}
```

### Slices ex4

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

// Create a Go program that reads some numbers from the command line
// and then calculates the sum and the product
// of all the numbers given at the command line.

func main() {
	sum, prod := 0., 1.

	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.ParseFloat(os.Args[i], 64)
		if err != nil {
			continue
		}

		if n < 2 || n > 10 {
			fmt.Println("Please enter between 2 and 10 numbers!")
			return
		}

		sum += n
		prod *= n
	}

	fmt.Printf("sum = %.2f, prod = %.2f\n", sum, prod)
}
```

### Slices ex5

```go
package main

import "fmt"

func main() {
	// Consider the following slice declaration:
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	// Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.
	var sum int

	// Print those elements and their sum.
	for _, num := range nums[1 : len(nums)-2] {
		fmt.Println(num)
		sum += num
	}
	fmt.Println("Sum:", sum)
}
```

### Slices ex6

```go
package main

import "fmt"

func main() {
	// Consider the following slice declaration:
	friends := []string{"Marry", "John", "Paul", "Diana"}

	// Using copy() function create a copy of the slice.
	friendsCopy := make([]string, len(friends))
	_ = copy(friendsCopy, friends)

	// Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
	friends[0] = "Mario"
	fmt.Printf("original: %#v\n", friends)
	fmt.Printf("copy: %#v\n", friendsCopy)
}
```

## Strings

### Strings intro

[Strings](https://go.dev/blog/strings)

```go
package main

import (
    "fmt"
)

func main() {

    // Strings are defined between double quotes "..."
    // Strings in Go are UTF-8 encoded by default
    // A string is in fact a slice of bytes in Go

    // declaring a string
    s1 := "Hi there  Go!"

    // printing a string
    fmt.Printf("%s\n", s1) // => Hi there  Go!
    fmt.Printf("%q\n", s1) // => "Hi there  Go!"

    // using double-quotes inside double quotes
    fmt.Println("He say: \"Hello!\"")

    // double quotes inside backticks (backquote)
    fmt.Println(`He say: "Hello!"`)

    // a string literal inclosed in backticks is called a raw string and it is interpreted literally.
    // backslashes or \n  have no special meaning
    s2 := `Hi there Go!`
    fmt.Println(s2)

    // declaring a multiline string
    fmt.Println("Price: 100 \nBrand: Nike")

    //the same with:
    fmt.Println(`
Price: 100
Brand: Nike`)

    // using backslashes inside a string:
    fmt.Println(`C:\Users\Andrei`)
    fmt.Println("C:\\Users\\Andrei")

    // concatenating strings (+)
    // Go creates a new string because strings are immutable in Go (this is not efficient).
    var s3 = "I love " + "Go " + "Programming"
    fmt.Println(s3 + "!") // -> I love Go Programming!

    // getting an element (byte) of a string:
    fmt.Println("Element at index zero:", s3[0]) // => 73 (ascii code for I)
    //  a string is in fact a slice of bytes in Go

    // strings are immutable and can't be changed
    // s3[5] = 'x' // => error: Cannon assign to s3[5].
}
```

### Runes

```go
package main

import (
    "fmt"
    "strings"
    "unicode/utf8"
)

func main() {

    // characters or rune literals are expressed in Go by enclosing them in single quotes
    // declaring a variable of type rune (alias to int32)
    var1 := 'a'
    fmt.Printf("Type: %T, Value:%d \n", var1, var1) // => Type: int32, Value:97
    // value is 97 (the code point for a)

    // declaring a string value that contains non-ascii characters
    str := "ţară" // ţară means country in Romanian
    // 't', 'a' ,'r' and 'a' are runes and each rune occupies beetween 1 and 4 bytes.

    //The len() built-in function returns the no. of bytes not runes or chars.
    fmt.Println(len(str)) // -> 6,  4 runes in the string but the length is 6

    // returning the number of runes in the string
    m := utf8.RuneCountInString(str)
    fmt.Println(m) // => 4

    // by using indexes we get the byte at that positioin, not rune.
    fmt.Println("Byte (not rune) at position 1:", str[1]) // -> 163

    // decoding a string byte by byte
    for i := 0; i < len(str); i++ {
        fmt.Printf("%c", str[i]) // -> Å£arÄ
    }

    fmt.Println("\n" + strings.Repeat("#", 10))

    // decoding a string rune by rune manually:
    for i := 0; i < len(str); {
        r, size := utf8.DecodeRuneInString(str[i:]) // it returns the rune in string in variable r
        //and its size in bytes in variable size

        // printing out each rune
        fmt.Printf("%c", r) // -> ţară
        i += size           // incrementing i by the size of the rune in bytes
    }

    fmt.Println("\n" + strings.Repeat("#", 10))

    // decoding a string rune by rune automatically:
    for i, r := range str { //the first value returned by range is the index of the byte in string where rune starts
        fmt.Printf("%d -> %c", i, r) // => ţară
    }

}
```

### Slicing strings

```go
package main

import "fmt"

func main() {

    // Slicing a string is efficient because it reuses the same backing array
    // Slicing returns bytes not runes

    s1 := "abcdefghijkl"
    fmt.Println(s1[2:5]) // -> cde, bytes from 2 (included) to 5 (excluded)

    s2 := "中文维基是世界上"
    fmt.Println(s2[0:2]) // -> � - the unicode representation of bytes from index 0 and 1.

    // returning a slice of runes
    // 1st step: converting string to rune slice
    rs := []rune(s2)
    fmt.Printf("%T\n", rs) // => []int32

    // 2st step: slicing the rune slice
    fmt.Println(string(rs[0:3])) // => 中文维
}
```

### Strings package

```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
 
    // declaring a variable of type func to call the Println function easier.
    p := fmt.Println
 
    // it returns true whether a substr is within a string
    result := strings.Contains("I love Go Programming!", "love")
    p(result) // -> True
 
    // it returns true whether any Unicode code points are within our string, and false otherwise.
    result = strings.ContainsAny("success", "xy")
    p(result) // => false
 
    // it reports whether a rune is within a string.
    result = strings.ContainsRune("golang", 'g')
    p(result) // => true
 
    // it returns the number of instances of a substring in a string
    n := strings.Count("cheese", "e")
    p(n) // => 3
 
    // if the substr is an empty string Count() returns 1 + the number of runes in the string
    n = strings.Count("five", "")
    p(n) // => 5 (1 + 4 runes)
 
    // ToUpper() and ToLower() return a new string with all the letters
    // of the original string converted to uppercase or lowercase.
    p(strings.ToLower("Go Python Java")) // -> go python java
    p(strings.ToUpper("Go Python Java")) // -> GO PYTHON JAVA
 
    // comparing strings (case matters)
    p("go" == "go") // -> true
    p("Go" == "go") // -> false
 
    // comparing strings (case doesn't matter) -> it is not efficient
    p(strings.ToLower("Go") == strings.ToLower("go")) // -> true
 
    // EqualFold() compares strings (case doesn't matter) -> it's efficient
    p(strings.EqualFold("Go", "gO")) // -> true
 
    // it returns a new string consisting of n copies of the string that is passed as the first argument
    myStr := strings.Repeat("ab", 10)
    p(myStr) // => abababababababababab
 
    // it returns a copy of a string by replacing a substring (old) by another substring (new)
    myStr = strings.Replace("192.168.0.1", ".", ":", 2) // it replaces the first 2 occurrences
    p(myStr)                                            // -> 192:168:0.1
 
    // if the last argument is -1 it replaces all occurrences of old by new
    myStr = strings.Replace("192.168.0.1", ".", ":", -1)
    p(myStr) // -> 192:168:0:1
 
    // ReplaceAll() returns a copy of the string s with all non-overlapping instances of old replaced by new.
    myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
    p(myStr) // -> 192:168:0:1
 
    // it slices a string into all substrings separated by separator and returns a slice of the substrings between those separators.
    s := strings.Split("a,b,c", ",")
    fmt.Printf("%T\n", s)                  // -> []string
    fmt.Printf("strings.Split():%#v\n", s) // => strings.Split():[]string{"a", "b", "c"}
 
    // If separator is empty Split function splits after each UTF-8 rune literal.
    s = strings.Split("Go for Go!", "")
    fmt.Printf("strings.Split():%#v\n", s) // -> []string{"G", "o", " ", "f", "o", "r", " ", "G", "o", "!"}
 
    // Join() concatenates the elements of a slice of strings to create a single string.
    // The separator string is placed between elements in the resulting string.
    s = []string{"I", "learn", "Golang"}
    j := strings.Join(s, "-")
    fmt.Printf("%T\n", j) // -> string
    p(j)                  // -> I-learn-Golang
 
    // splitting a string by whitespaces and newlines.
    myStr = "Orange Green \n Blue Yellow"
    fields := strings.Fields(myStr) // it returns a slice of strings
    fmt.Printf("%T\n", fields)      // -> []string
    fmt.Printf("%#v\n", fields)     // -> []string{"Orange", "Green", "Blue", "Yellow"}
 
    // TrimSpace() removes leading and trailing whitespaces and tabs.
    s1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux!\n ")
    fmt.Printf("%q\n", s1) // "Goodbye Windows, Welcome Linux!"
 
    // To remove other leading and trailing characters, use Trim()
    s2 := strings.Trim("...Hello, Gophers!!!?", ".!?")
    fmt.Printf("%q\n", s2) // "Hello, Gophers"
 
}
```

## Strings exercises

### Strings ex1

```go
package main

import "fmt"

func main() {
	// 1. Using the var keyword declare a string called name and initialize it with your name.
	var name = "Mario"
	// 2. Using short declaration syntax declare a string called country and assign the country you are living in to the string variable.
	country := "Italia"

	// 3. Print the following string on multiple lines like this:
	// Your name: `here the value of name variable`
	// Country: `here the value of country variable`
	fmt.Printf("Your name: '%s'\nCountry: '%s'\n", name, country)

	// 4. Print out the following strings:
	// a) He says: "Hello"
	// b) C:\Users\a.txt
	fmt.Println(`He says: "Hello"
C:\Users\a.txt`)
}
```

### Strings ex2

```go
package main

import (
	"fmt"
)

func main() {
	// 1. Declare a rune called r that stores the non-ascii letter ă
	r := 'ă'

	// 2. Print out the type of r
	fmt.Printf("%v\n", r)
	fmt.Printf("%T\n", r)

	// 3. Declare 2 strings that contains the values ma and m
	str1, str2 := "ma", "m"

	// 4. Concatenate the strings and the rune in a new string
	// (the new string will hold the value mamă ).
	str3 := str1 + str2 + string(r)
	fmt.Println(str3)
}
```

### Strings ex3

```go
package main

import "fmt"

func main() {
	// Consider the following string declaration:
	s1 := "țară means country in Romanian"

	// 1. Iterate over the string and print out byte by byte
	for i := range s1 {
		fmt.Printf("%b\n", s1[i])
	}

	fmt.Println()

	// 2. Iterate  over the string and print out rune by rune and the byte index where the rune starts in the string
	for i, c := range s1 {
		fmt.Printf("%d %d\n", i, c)
	}
}
```

### Strings ex4

```go
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	// There are some errors
	// s1[0] = 'x'
	strings.Replace(s1, s1, "G", 'x')

	// printing the number of runes in the string
	fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))
}
```

### Strings ex5

```go
package main

import "fmt"

func main() {
	// Consider the following string declaration:
	s := "你好 Go!"

	// 1. Convert the string to a byte slice.
	b := []byte(s)

	// 2. Print out the byte slice
	fmt.Printf("%b\n", b)

	// 3. Iterate over the byte slice and print out each index and byte in the byte slice
	for i, v := range b {
		fmt.Printf("%d %b\n", i, v)
	}
}
```

### Strings ex6

```go
package main

import "fmt"

func main() {
	// Consider the following string declaration:
	s := "你好 Go!"

	// 1. Convert the string to a rune slice.
	r := []rune(s)

	// 2. Print out the rune slice
	fmt.Printf("%v\n", r)

	// 3. Iterate over the rune slice and print out each index and rune in the rune slice
	for i, v := range r {
		fmt.Printf("%d %c\n", i, v)
	}
}
```

## Maps

### Maps in action

[Maps](https://go.dev/blog/maps)

```go
package main
 
import "fmt"
 
func main() {
 
    // declaring a map with keys of type string and values of type string
    var employees map[string]string
    //the zero value of a map is nil
 
    // the zero value of a map is nil
    fmt.Printf("%#v\n", employees) // -> map[string]string(nil).
 
    fmt.Printf("No. of elements: %d\n", len(employees)) // => No. of elements: 0
 
    // getting the corresponding value of a key
    // if the key doesn't exist or the map is not initialized it returns the zero value for the value type
    fmt.Printf("The value for key %q is %q \n", "Dan", employees["Dan"]) // => The value for key "Dan" is ""
 
    // key must be of comparable types
    // var m1 map[[]int]string // error -> invalid map key type []int (slices are not comparable)
 
    // inserting a key:value pair in a nil map
    // employees["Dan"] = "Programmer" // error -> panic: assignment to entry in nil map
 
    // declaring and initializing a map using a map literal
    people := map[string]float64{} // empty map
 
    // inserting key:value pairs in a map
    people["John"] = 30.5
    people["Marry"] = 22
 
    fmt.Printf("%v\n", people) // => map[John:30.5 Marry:22]
 
    // declaring and initializing a map using the make() function:
    map1 := make(map[int]int)
    fmt.Printf("map1: %#v\n", map1) // -> map1: map[int]int{}
    map1[4] = 8
 
    // declaring and initializing a map using a map literal
    balances := map[string]float64{
        "USD": 233.11,
        "EUR": 555.11,
        //50: "ABC"  //illegal, all keys have the same type which is string and all values have the same type which is float64
        "CHF": 600, //this last comma (,) is mandatory when declaring the map on multiple lines
    }
    fmt.Println(balances) // => map[CHF:600 EUR:555.11 USD:233.11]
 
    //If we declare and initialize a map on a single line, it's not mandatory to use a comma after the last element
    m := map[int]int{1: 3, 4: 5, 6: 8}
    _ = m
 
    // initializing a map with duplicate keys
    // n := map[int]int{1: 3, 4: 5, 6: 8, 1: 4} // error -> duplicate key 1 in map literal
 
    // if the key exists it updates its value and if the key doesn't exist it adds the key: value pair
    balances["USD"] = 500.5
    balances["GBP"] = 800.8
    fmt.Println(balances) // => map[CHF:600 EUR:555.11 GBP:800.8 USD:500.5]
 
    // "comma ok" idiom is used to distinguish between a missing key:value pair and an existing key with value zero
    v, ok := balances["RON"]
 
    //v is the key's corresponding value
    // ok is bool type value which is true if the key exists and false otherwise
 
    if ok {
        fmt.Println("The RON Balance is: ", v)
    } else {
        fmt.Println("The RON key doesn't exist in the map!")
    }
 
    // iterating over a map
    for k, v := range balances {
        fmt.Printf("Key: %#v, Value: %#v\n", k, v)
    }
 
    //starting with go 1.12 fmt.Printf() function prints out the map sorted by key.
    fmt.Printf("balances: %v\n", balances) // => balances: map[CHF:600 EUR:555.11 GBP:800.8 USD:500.5]
 
    // deleting a key:value pair from the map
    delete(balances, "USD")
 
    //** COMPARING MAPS **//
 
    // Maps cannot be compared using == operator. A map can be compared only to nil.
    a := map[string]string{"A": "X"}
    b := map[string]string{"B": "X"}
 
    // fmt.Println(a == b) // error -> invalid operation: a == b (map can only be compared to nil)
 
    // to compare 2 maps that have the keys and values of type string
    // we get a string representation of the maps and compare those strings.
 
    // getting a string representation of maps called a and b
    s1 := fmt.Sprintf("%s", a)
    s2 := fmt.Sprintf("%s", b)
 
    if s1 == s2 {
        fmt.Println("Maps are equal")
    } else {
        fmt.Println("Maps are not equal")
    }
 
    //** CLONING A MAP **//
 
    // When creating a map variable Go creates a pointer to a map header value in memory.
    // The key: value pairs of the map are not stored directly into the map.
    // They are stored in memory at the address referenced by the map header.
 
    friends := map[string]int{"Dan": 40, "Maria": 35}
 
    // neighbors is not a copy of friends.
    // both maps reference the same data structure in memory
    neighbors := friends
 
    // modifying friends AND neighbors
    friends["Dan"] = 30
 
    fmt.Println(neighbors) // -> map[Dan:30 Maria:35]
 
    // How to clone a map?
    // 1. initialize a new map
    colleagues := make(map[string]int)
 
    // colleagues = friends // -> ERROR, illegal with maps!
 
    // 2. use a for loop to copy each element into the new map
    for k, v := range friends {
        colleagues[k] = v
    }
 
    // colleagues and friends are different maps in memory!
}
```

## Maps exercises

### Maps ex1

```go
package main

import "fmt"

func main() {
	// 1. Declare a map called m1 which value is nil. Print out its type and value.
	var m1 map[int]string
	fmt.Printf("%#v\n", m1)

	// 2. Declare a map called m2. It's keys are of type int and values of type string.
	// Initialize the map using  a map literal with two key:value pairs.
	m2 := map[int]string{1: "Mario"}
	fmt.Printf("%#v\n", m2)

	// 3. Add the following key: value to the map: 10: "Abba"
	m2[10] = "Abba"

	// 4. Retrieve the value of an existing key and the value of a non existing key. Think about the results.
	val, ok := m2[1]
	if ok {
		fmt.Printf("Val 1 exists: %s\n", val)
	} else {
		fmt.Println("Val 1 does not exist")
	}

	val, ok = m2[11]
	if ok {
		fmt.Printf("Val 11 exists: %s\n", val)
	} else {
		fmt.Println("Val 11 does not exist")
	}
}
```

### Maps ex2

```go
package main

import "fmt"

func main() {
	// There are some errors
	// var m1 map[int]bool
	var m1 = make(map[int]bool)
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	// There are some errors
	// fmt.Println(m2 == m3)
	m2Str := fmt.Sprintf("%v", m2)
	m3Str := fmt.Sprintf("%v", m3)
	fmt.Println(m2Str == m3Str)
}
```

### Maps ex3

```go
package main

import "fmt"

func main() {
	// Consider the following map declaration:
	// m := map[int]bool{"1": true, 2: false, 3: false}

	// 1. The above map declaration has an error. Solve the error!
	m := map[int]bool{1: true, 2: false, 3: false}

	// 2. Delete a key:value pair from the map.
	delete(m, 1)

	// 3. Iterate over the map and print out each key and value.
	for k, v := range m {
		fmt.Printf("%d -> %t\n", k, v)
	}
}
```

## Files

### File operations

```go
package main
 
import (
    "fmt"
    "log"
    "os"
)
 
func main() {
 
    //** Use valid paths according to your OS. **//
 
    // CREATING A FILE
 
    // os.Create() function creates a file if it doesn't already exist. If it exists, the file is truncated.
    // it returns a file descriptor which is a pointer to os.File and an error value.
    newFile, err := os.Create("a.txt")
 
    // error handling
    if err != nil {
        // log the error and exit the program
        log.Fatal(err) // the idiomatic way to handle errors
 
    }
 
    // TRUNCATING A FILE
    err = os.Truncate("a.txt", 0) //0 means completely empty the file.
 
    // error handling
    if err != nil {
        log.Fatal(err)
    }
 
    // CLOSING THE FILE
    newFile.Close()
 
    // OPEN AND CLOSE AN EXISTING FILE
    file, err := os.Open("a.txt") // open in read-only mode
 
    // error handling
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
 
    //OPENING a FILE WITH MORE OPTIONS
    file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)
    // We can Use opening attributes individually or combined
    // using an OR between them
    // e.g. os.O_CREATE|os.O_APPEND
    // or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
    // os.O_RDONLY // Read only
    // os.O_WRONLY // Write only
    // os.O_RDWR // Read and write
    // os.O_APPEND // Append to end of file
    // os.O_CREATE // Create is none exist
    // os.O_TRUNC // Truncate file when opening
 
    // error handling
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
 
    // GETTING FILE INFO
    var fileInfo os.FileInfo
    fileInfo, err = os.Stat("a.txt")
 
    p := fmt.Println
    p("File Name:", fileInfo.Name())        // => File Name: a.txt
    p("Size in bytes:", fileInfo.Size())    // => Size in bytes: 0
    p("Last modified:", fileInfo.ModTime()) // => Last modified: 2019-10-21 16:16:00.325037748 +0300 EEST
    p("Is Directory? ", fileInfo.IsDir())   // => Is Directory?  false
    p("Pemissions:", fileInfo.Mode())       // => Pemissions: -rw-r-----
 
    // CHECKING IF FILE EXISTS
    fileInfo, err = os.Stat("b.txt")
    // error handling
    if err != nil {
        if os.IsNotExist(err) {
            log.Fatal("The file does not exist")
        }
    }
 
    // RENAMING AND MOVING A FILE
    oldPath := "a.txt"
    newPath := "aaa.txt"
    err = os.Rename(oldPath, newPath)
    // error handling
    if err != nil {
        log.Fatal(err)
    }
 
    // REMOVING A FILE
    err = os.Remove("aa.txt")
    // error handling
    if err != nil {
        log.Fatal(err)
    }
 
}
```

### Writing bytes

```go
package main
 
import (
    "io/ioutil"
    "log"
    "os"
)
 
func main() {
 
    // opening the file in write-only mode if the file exists and then it truncates the file.
    // if the file doesn't exist it creates the file with 0644 permissions
    file, err := os.OpenFile(
        "b.txt",
        os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
        0644,
    )
    // error handling
    if err != nil {
        log.Fatal(err)
    }
    // defer closing the file
    defer file.Close()
 
    // WRITING BYTES TO FILE
 
    byteSlice := []byte("I learn Golang! 传")   // converting a string to a bytes slice
    bytesWritten, err := file.Write(byteSlice) // writing bytes to file.
    // It returns the no. of bytes written and an error value
    // error handling
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bytes written: %d\n", bytesWritten) // => 2019/10/21 16:26:16 Bytes written: 19
 
    // WRITING BYTES TO FILE USING ioutil.WriteFile()
 
    // ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
    // if the file doesn't exist WriteFile() creates it
    // and if it already exists the function will truncate it before writing to file.
 
    bs := []byte("Go Programming is cool!")
    // err = ioutil.WriteFile("c.txt", bs, 0644) //=> deprecated
    err = os.WriteFile("c.txt", bs, 0644)
    // error handling
    if err != nil {
        log.Fatal(err)
    }
}
```

### Buffer writer

```go
package main
 
import (
    "bufio"
    "log"
    "os"
)
 
func main() {
 
    // Opening the file for writing
    file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
    // error handling
    if err != nil {
        log.Fatal(err)
    }
    // defer closing the file
    defer file.Close()
 
    // Creating a buffered writer from the file variable using bufio.NewWriter()
    bufferedWriter := bufio.NewWriter(file)
 
    // declaring a byte slice
    bs := []byte{97, 98, 99}
 
    // writing the byte slice to the buffer in memory
    bytesWritten, err := bufferedWriter.Write(bs)
 
    // error handling
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)
    // => 2019/10/21 16:30:59 Bytes written to buffer (not file): 3
 
    // checking the available buffer
    bytesAvailable := bufferedWriter.Available()
    log.Printf("Bytes available in buffer: %d\n", bytesAvailable)
    // => 2019/10/21 16:30:59 Bytes available in buffer: 4093
 
    // writing a string (not a byte slice) to the buffer in memory
    bytesWritten, err = bufferedWriter.WriteString("\nJust a random string")
 
    // error handling
    if err != nil {
        log.Fatal(err)
    }
 
    // checking how much data is stored in buffer, just waiting to be written to disk
    unflushedBufferSize := bufferedWriter.Buffered()
    log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
    // -> 24 (3 bytes in the byte slice + 21 bytes in the ASCII string; here each rune is 1 byte because the string is ASCII)
 
    // The bytes have been written to buffer, not yet to file.
    // Writing from buffer to file.
    bufferedWriter.Flush()
}
```

### Buffer reader

```go
package main
 
import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "strings"
)
 
func main() {
 
    //** READING INTO A BYTE SLICE USING io.ReadFull() **//
 
    // Opening the file in read-only mode. The file must exist (in the current working directory)
    // Use a valid path!
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
 
    // declaring a byte slice and initializing it with a length of 2
    byteSlice := make([]byte, 2)
 
    // io.ReadFull() returns an error if the file is smaller than the byte slice.
    // it reads the file into the byte slice up to its length
    numberBytesRead, err := io.ReadFull(file, byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Number of bytes read: %d\n", numberBytesRead)
    log.Printf("Data read: %s\n", byteSlice)
 
    fmt.Println(strings.Repeat("#", 20))
 
    //** READING WHOLE FILE INTO A BYTESLICE USING ioutil.ReadAll() **//
 
    // Opening another file (from the current working directory)
    file, err = os.Open("main.go")
    if err != nil {
        log.Fatal(err)
    }
 
    // ioutil.ReadAll() reads every byte from the file and return a slice of unknown size
    // data, err := ioutil.ReadAll(file)  //=> deprecated
    data, err := io.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }
 
    fmt.Printf("Data as string: %s\n", data)
    fmt.Println("Number of bytes read:", len(data))
 
    //** READING WHOLE FILE INTO MEMORY USING ioutil.ReadFile() **//
 
    // ioutil.ReadFile() reads a file into byte slice
    // this function handles opening and closing the file.
    // data, err = ioutil.ReadFile("test.txt")  //=> deprecated
    data, err = os.ReadFile("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Data read: %s\n", data)
}
```

### Read by line

```go
package main
 
import (
    "bufio"
    "fmt"
    "log"
    "os"
)
 
func main() {
 
    // opening the file in read-only mode. The file must exist (in the current working directory)
    // use a valid path!
    file, err := os.Open("my_file.txt")
    // error handling
    if err != nil {
        log.Fatal(err)
    }
    // defer closing the file
    defer file.Close()
 
    // the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
    scanner := bufio.NewScanner(file)
 
    // the default scanner is bufio.ScanLines and that means it will scan a file line by line.
    // there are also bufio.ScanWords and bufio.ScanRunes.
    // scanner.Split(bufio.ScanLines)
 
    // scanning for next token in this case \n which is the line delimiter.
    success := scanner.Scan() //read a line
    if success == false {
        // false on error or EOF. Check for errors
        err = scanner.Err()
        if err == nil {
            log.Println("Scan was completed and it reached End Of File.")
        } else {
            log.Fatal(err)
        }
    }
 
    // Getting the data from the scanner with Bytes() or Text()
    fmt.Println("First Line found:", scanner.Text())
    //If we want the next token, so the next line or \n, we call scanner.Scan() again
 
    // Reading the whole remaining part of the file:
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
 
    // Checking for any possible errors:
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
```

### User input

```go
package main
 
import (
    "bufio"
    "fmt"
    "log"
    "os"
)
 
func main() {
 
    // creating a scanner
    scanner := bufio.NewScanner(os.Stdin) //os.Stdin reads the output from the command line
    fmt.Printf("%T\n", scanner)           //pointer to bufio.scanner
 
    scanner.Scan() //it waits for the input and buffers the input untill a new line
 
    // gettting the scanned data
    text := scanner.Text()   // string type
    bytes := scanner.Bytes() // uint8[] slice type
 
    fmt.Println("Input text:", text)
    fmt.Println("Input bytes:", bytes)
 
    // reading the input continously until a specific string is scanned
    for scanner.Scan() {
        text = scanner.Text()
        fmt.Println("You entered:", text)
        if text == "exit" {
            fmt.Println("Exiting the scanning ...")
            break
        }
    }
 
    // error handling
    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
}
```

## Files exercises

### File ex1

```go
package main

import (
	"log"
	"os"
)

func main() {
	// Create a new file in the current working directory called info.txt and then close the file.
	// If the file cannot be created (no permissions, wrong path etc) then print out the error and stop the program (do error handling).
	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
```

### File ex2

```go
package main

import (
	"log"
	"os"
)

func main() {
	// Rename the file created at Exercise #1 to data.txt
	// Check if file exists before renaming it.
	// If it doesn't exist print a message and stop the program.

	oldName := "info.txt"
	newName := "data.txt"

	_, err := os.Stat(oldName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("File %s does not exist\n", oldName)
		}
	}

	err = os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}
```

### File ex3

```go
package main

import (
	"log"
	"os"
)

func main() {
	// Remove the file created at exercise #1.
	// Take care that the file is now called data.txt (it has been renamed at exercise #2).
	// Perform error handling.

	err := os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
```

### File ex4

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create a Go Program that reads the entire contents of a file called info.txt into a string.
	// You can use ioutil.ReadAll() or any other function you want.
	// The file exists in the current working directory.

	file, err := os.OpenFile("README.md", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	fmt.Println(string(fileBytes))
}
```

### File ex5

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a Go Program that reads the entire contents of a file called info.txt  using a scanner (bufio package) line by line.
	// The file exists in the current working directory.

	file, err := os.OpenFile("README.md", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
```

### File ex6

```go
package main

import (
	"io/ioutil"
	"log"
)

func main() {
		bs := []byte("The Go gopher is an iconic mascot!")
	err := ioutil.WriteFile("info.txt", bs, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
```

## Structs

### Creating structs

```go
package main
 
import "fmt"
 
func main() {
 
    // creating a struct type
    type book struct {
        title  string //the fields of the book struct
        author string //each field must be unique inside a struct
        year   int
    }
 
    // combining different fields of the same type on the same line
    type book1 struct {
        title, author string
        year, pages   int
    }
 
    // declaring, initializing and assigning a new book value, all in one step
    lastBook := book{"The Divine Comedy", "Dante Aligheri", 1320} //this is a struct literal and order matters
    fmt.Println(lastBook)
 
    // Declaring a new book value by specifying field: value (order doesn't matter)
    bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945}
    _ = bestBook
 
    //if we create a new struct value by omitting some fields they will be zero-valued according to their type
    aBook := book{title: "Just a random book"}
    fmt.Printf("%#v\n", aBook) // => main.book{title:"Just a random book", author:"", year:0}
 
    // retrieving the value of a struct field
    fmt.Println(lastBook.title) // => The Divine Comedy
 
    // selecting a field that doesn't exist raises an error
    // pages := lastBook.pages // error -> lastBook.pages undefined (type book has no field or method pages)
 
    // updating a field
    lastBook.author = "The Best"
    lastBook.year = 2020
    fmt.Printf("lastBook: %+v\n", lastBook) // => lastBook: {title:The Divine Comedy author:The Best year:2020}
    // + modifier with %v  printed out both the field names and their values
 
    // comparing struct values
    // two struct values are equal if their corresponding fields are equal.
    randomBook := book{title: "Random Title", author: "John Doe", year: 100}
    fmt.Println(randomBook == lastBook) // => false
 
    // = creates a copy of a struct
    myBook := randomBook
    myBook.year = 2020              // modifying only myBook
    fmt.Println(myBook, randomBook) // => {Random Title John Doe 2020} {Random Title John Doe 100}
 
}
```

### Working with structs

```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
 
    // an anonymous struct is a struct with no explicitly defined struct type alias.
    diana := struct {
        firstName, lastName string
        age                 int
    }{
        firstName: "Diana",
        lastName:  "Muller",
        age:       30,
    }
 
    fmt.Printf("%#v\n", diana)
    // =>struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"Muller", age:30
 
    //** ANONYMOUS FIELDS **//
 
    // fields type becomes fields name.
    type Book struct {
        string
        float64
        bool
    }
 
    b1 := Book{"1984 by George Orwell", 10.2, false}
    fmt.Printf("%#v\n", b1) // => main.Book{string:"1984 by George Orwell", float64:10.2, bool:false}
 
    fmt.Println(b1.string) // => 1984 by George Orwell
 
    // mixing anonymous with named fields:
    type Employee1 struct {
        name   string
        salary int
        bool
    }
 
    e := Employee1{"John", 40000, false}
    fmt.Printf("%#v\n", e) // => main.Employee1{name:"John", salary:40000, bool:false}
    e.bool = true          // changing a field
 
    fmt.Println(strings.Repeat("#", 10))
 
    //** EMBEDDED STRUCTS **//
    // An embedded struct is just a struct that acts like a field inside another struct.
 
    // define a new struct type
    type Contact struct {
        email, address string
        phone          int
    }
 
    // define a struct type that contains another struct as a field
    type Employee struct {
        name        string
        salary      int
        contactInfo Contact
    }
 
    // declaring a value of type Employee
    john := Employee{
        name:   "John Keller",
        salary: 3000,
        contactInfo: Contact{
            email:   "jkeller@company.com",
            address: "Street 20, London",
            phone:   042324234,
        },
    }
 
    fmt.Printf("%+v\n", john)
    // => {name:John Keller salary:3000 contactInfo:{email:jkeller@company.com address:Street 20, London phone:295619381404}}
 
    // accessing a field
    fmt.Printf("Employee's salary: %d\n", john.salary)
 
    // accessing a field from the embedded struct
    fmt.Printf("Employee's email:%s\n", john.contactInfo.email) // => Employee's email:jkeller@company.com
 
    // updating a field
    john.contactInfo.email = "new_email@thecompany.com"
    fmt.Printf("Employee's new email address:%s\n", john.contactInfo.email)
    // =>  Employee's new email address:new_email@thecompany.com
}
```

## Structs exercises

### Structs ex1

```go
package main

import "fmt"

func main() {
	// 1. Create your own struct type called person that stores the following data:
	// name, age and a list with favorite colors.
	type person struct {
		name   string
		age    int
		colors []string
	}

	// 2. Declare and initialize two values of type person, one called me and another called you.
	me := person{name: "Mario", age: 51, colors: []string{"black", "blue"}}
	you := person{name: "Mariarosa", age: 50, colors: []string{"green"}}

	// 3. Print out the struct values.
	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)
}
```

### Structs ex2

```go
package main

import "fmt"

func main() {
	// Consider the code from the previous exercise and:
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{name: "Mario", age: 51, colors: []string{"black", "blue"}}
	you := person{name: "Mariarosa", age: 50, colors: []string{"green"}}

	// 1. Change the name or the struct value called me to "Andrei"
	me.name = "Andrei"

	// 2. Take in a new variable the favorites colors of struct value called you.
	// Print out the type and the value of the new variable.
	yourColors := you.colors
	fmt.Printf("Your colors: %#v\n", yourColors)

	// 3. Add a new favorite color to the second struct value called you.
	yourColors = append(yourColors, "pink")
	you.colors = yourColors

	// 4. Print out the struct values.
	fmt.Printf("Your colors: %#v\n", you.colors)
}
```

### Structs ex3

```go
package main

import "fmt"

func main() {
	// Consider the code from Exercise #1.
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{name: "Mario", age: 51, colors: []string{"black", "blue"}}

	// Iterate and print out the favorite colors of the struct value called me.
	for _, color := range me.colors {
		fmt.Println(color)
	}
}
```

### Structs ex4

```go
package main

import "fmt"

func main() {
	// Change the code from Exercise #1 and:

	// 1. Create a new struct type called grades with  2 fields: grade and course
	type grades struct {
		grade  int
		course string
	}

	// 2. Add another field of type grades to person struct type (embedded struct).
	type person struct {
		name   string
		age    int
		colors []string
		grades grades
	}

	// 3. Change the initialization of the struct values called me and you to contain information about the grades.
	me := person{name: "Mario", age: 51, colors: []string{"black", "blue"}, grades: grades{grade: 100, course: "a"}}

	// 4. Change the grade and the course of one struct value to "Golang" and 98.
	me.grades.course = "Golang"
	me.grades.grade = 98

	// 5. Print out the struct values.
	fmt.Printf("me: %#v\n", me)
}
```

## Functions

### Function

```go
package main

import "fmt"

func main() {
	f1()
}

func f1() {
	fmt.Println("F1")
}
```

### Function parameters

[No pass by reference](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go)

```go
package main

import (
	"fmt"
	"math"
)

// defining a function with no parameters
func f1() {
	fmt.Println("This is f1() function")
}

// defining a function with 2 parameters, a and b
func f2(a int, b int) {
	//a and b are local to the function
	fmt.Println("Sum:", a+b)
}

// defining a function using shorthand parameters notation
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// defining a function that have one parameter of type float64 and returns a value of type float64
func f4(a float64) float64 {
	return math.Pow(a, a)
	//any statements below the return statement are never executed

}

// defining a function that have two parameters of type int and returns two values of type int
func f5(a, b int) (int, int) {
	return a * b, a + b
}

// defining a function that have one parameter of type int and returns a "named parameter"
func sum(a, b int) (s int) {
	fmt.Println("s:", s) // -> s is a variable with the zero value inside the function
	s = a + b

	// it automatically return s
	return // This is known as a "naked" return.
}

func main() {
	// calling functions
	f1() // => This is f1() function

	f3(3, 4, 5, 4., 5.5, "ss") // => 3 4 5 4 5.5 ss

	fmt.Println(f4(5.1))

	a, b := f5(6, 7)
	fmt.Printf("a:%d, b:%d\n", a, b) // => a:42, b:13

	ss := sum(4, 5)
	fmt.Println(ss) // -> 9

	// There are no default arguments in Go //
}
```

### Variadic functions

```go
package main
 
import (
    "fmt"
    "strings"
)
 
// creating a variadic function
func f1(a ...int) {
    fmt.Printf("%T\n", a) // => []int, slice of int
    fmt.Printf("%#v\n", a)
}
 
// variadicfunction that modifies one of the arguments passed.
func f2(a ...int) {
    a[0] = 50
}
 
// creating a variadic function that calculates and returns the sum and product of its arguments
func sumAndProduct(a ...float64) (float64, float64) {
    sum := 0.
    product := 1.
 
    for _, v := range a {
        sum += v
        product *= v
    }
 
    return sum, product
}
 
// mixing variadic and non-variadic parameters is allowed
// non-variadic parameters are always before the variadic parameter
func personInformation(age int, names ...string) string {
    fullName := strings.Join(names, " ")
    returnString := fmt.Sprintf("Age: %d, Full Name:%s", age, fullName)
    return returnString
}
 
func main() {
 
    // calling variadic functions
    // a variadic function can be invoked with zero or more arguments
    f1(1, 2, 3, 4)
 
    f1() // a is: []int(nil)
 
    // an example of a well-known variadic function is append() built-in function.
    // it appends items to an existing slice and returns back the same slice.
    nums := []int{1, 2}
    nums = append(nums, 3, 4)
 
    s, p := sumAndProduct(2., 5., 10.)
    fmt.Println(s, p) // -> 17 100
 
    info := personInformation(35, "Wolfgang", "Amadeus", "Mozart")
    fmt.Println(info) // => Age: 35, Full Name:Wolfgang Amadeus Mozart
}
```

### Defer

```go
package main
 
import (
    "fmt"
)
 
func foo() {
    fmt.Println("This is foo()!")
}
 
func bar() {
    fmt.Println("This is bar()!")
}
 
func foobar() {
    fmt.Println("This is foobar()!")
}
 
func main() {
 
    // a defer statement defers or postpones the execution of a function until the surrounding function returns.
 
    // by deferring foo() it will execute it just before exiting the surrounding function which is main()
    defer foo()
    bar()
 
    fmt.Println("Just a string after deferring foo() and calling bar()")
 
    // if there are more functions deferred, Go will execute them in the reverse order they were deferred
    defer foobar()
 
    /*
        When executing the program the fallowing output is printed out:
 
        This is bar()!
        Just a string after deferring foo() and calling bar()
        This is foobar()!
        This is foo()!
    */
 
}
```

### Anonymous functions

```go
package main
 
import "fmt"
 
// function that takes an int as an argument and returns another function that returns an int
func increment(x int) func() int {
    return func() int {
        x++
        return x
    }
}
 
func main() {
    // declaring an anonymous functions
    func(msg string) {
        fmt.Println(msg)
    }("I'm an anonymous function!") // calling the anonymous function
 
    // calling the increment function. It returns an anonymous function
    a := increment(10)
    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(a())
}
```

## Functions exercises

### Function ex1

```go
package main

import (
	"fmt"
)

// Create a function called cube() that takes a parameter of type float64 and returns the cube of that parameter
// (the parameter to the power of 3).
func cube(n float64) float64 {
	return n * n * n
}

func main() {
	fmt.Printf("%.3f\n", cube(3))
}
```

### Function ex2

```go
package main

import "fmt"

// Create a Go program with a function called f1() that takes a parameter of type uint and returns 2 values:
// a) the factorial of n
// b) the sum of all integer numbers greater than zero (>0) and less than or equal to n (<=n)

func f1(n uint) (uint, uint) {
	var fact uint = 1
	var sum uint

	// factorial of n
	var i uint = n
	for i > 0 {
		fact *= i
		i--
	}

	// sum of [0,n]
	i = 0
	for i <= n {
		sum += i
		i++
	}

	return fact, sum
}

func main() {
	f, s := f1(1)
	fmt.Printf("fact = %d, sum = %d\n", f, s)
	f, s = f1(2)
	fmt.Printf("fact = %d, sum = %d\n", f, s)
	f, s = f1(3)
	fmt.Printf("fact = %d, sum = %d\n", f, s)
	f, s = f1(4)
	fmt.Printf("fact = %d, sum = %d\n", f, s)
}
```

### Function ex3

```go
package main

import (
	"fmt"
	"log"
	"strconv"
)

// Write a function called myFunc() that takes exactly one argument which is an int number written between double quotes
// (this is in fact a string).
//
// If the argument is integer 'n', the function should return the result of n + nn + nnn
// Example: myFunc('5') returns 5 + 55 + 555 which is 615 and myFunc('9') returns 9 + 99 + 999 which is 1107

func myFunc(n string) int {
	num, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}
	return num * (1 + 11 + 111)
}

func main() {
	fmt.Println(myFunc("5"))
	fmt.Println(myFunc("9"))
}
```

### Function ex4

```go
package main

import "fmt"

// Create a function with the identifier sum that takes in a variadic parameter of type int
// and returns the sum of all values of type int passed in.
func sum(n ...int) int {
	total := 0
	for _, num := range n {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	fmt.Println(sum())
}
```

### Function ex5

```go
package main

import (
	"fmt"
	"strings"
)

// Create a function called searchItem() that takes 2 parameters:
// a) a string slice and
// b) a string.

// The function should search for the string (the second parameter) in the slice
// (the first parameter) and returns true if it finds the string in the slice and false otherwise.
// The function does a case-sensitive search.

func searchItem(vals []string, search string) bool {
	for _, val := range vals {
		if strings.Contains(val, search) {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false
}
```

### Function ex6

```go
package main

import (
	"fmt"
	"strings"
)

// Create a function called searchItem() that takes 2 parameters:
// a) a string slice and
// b) a string.

// The function should search for the string (the second parameter) in the slice
// (the first parameter) and returns true if it finds the string in the slice and false otherwise.
// The function does a case-sensitive search.

func searchItem(vals []string, search string) bool {
	for _, val := range vals {
		if strings.Contains(val, search) {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false
}
```

### Function ex7

```go
package main

import (
	"fmt"
	"strings"
)

// Change the function from the previous exercise to do a case-insensitive search.

func searchItem(vals []string, search string) bool {
	for _, val := range vals {
		if strings.EqualFold(val, search) {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "Bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "Pig")
	fmt.Println(result) // => false
}
```

### Function ex8

```go
package main

import "fmt"

// Modify only the line in the main() body function where the print() function is invoked
// so that the program will print out Hello, Go playground! and then
// The Go gopher is the iconic mascot of the Go project.

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}
```

### Function ex9

```go
package main

import "fmt"

// Create a function that takes in an int value and prints out that value.
// Assign the function to a variable, print out the type of the variable and then call that function through the variable name.

func printN(n int) {
	fmt.Println(n)
}

func main() {
	p := printN
	fmt.Printf("%T\n", p)
	p(1)
}
```

## Pointers

### Pointers intro

```go
package main
 
import "fmt"
 
func main() {
 
    // the & (ampersand) operator also known as address of operator returns the memory address of a variable.
    name := "Andrei"
    fmt.Println(&name) // -> 0xc0000101e0
 
    //** DECLARING AND INITIALIZING POINTERS **//
 
    var x int = 2
    // the expression &x means the address of x and creates a pointer to an integer variable,
    // ptr is of type *int, which is pronounced "pointer to int".
    ptr := &x
    fmt.Printf("ptr is of type %T with value %v and address %p\n", ptr, ptr, &ptr) // -> p is of type *int with value 0xc000014140.
 
    // declaring a pointer without initializing it
    // its zero value is nil
    var ptr1 *float64
    _ = ptr1
 
    // creating a pointer using new() built-in function.
    p := new(int) // it creates a pointer called p that is a pointer to an int type
 
    x = 100
    p = &x // initializing p
 
    fmt.Printf("p is of type %T with value %v\n", p, p) // => p is of type *int with value 0xc000014140
    fmt.Printf("address of x is %p\n", &x)              // => address of x is 0xc000016120
 
    //** THE DEREFERENCING OPERATOR **//
 
    // * in front of a pointer is called the dereferencing operator
    *p = 90 //equivalent to x = 90 because *p is x
    // x and *p is the same thing.
 
    fmt.Println(*p)                  // => 90
    fmt.Println("*p == x:", *p == x) // => *p == x: true
 
    fmt.Println("Value of x:", *p) // => Value of x: 90 , equivalent to fmt.Println(x)
 
    *p = 10        // If I write *p = 10, this is equivalent to x = 10
    *p = *p / 2    //dividing x through the pointer
    fmt.Println(x) // -> 5
 
    // In a nutshell:
    // &value => pointer -> if you have a value you turn it into an address or pointer by using the ampersand operator
    // *pointer => value  -> and if you have pointer you turn it into value value by using the star operator
 
}
```

### Pointers and functions

```go
package main
 
import "fmt"
 
// declaring a function that takes an int, a float, a string and a bool value.
// the function works on copy so the changes are not seen outside (pass by value)
func changeValues(quantity int, price float64, name string, sold bool) {
    quantity = 3
    price = 500.5
    name = "Mobile Phone"
    sold = false
}
 
// declaring a function that takes in a pointer to int, a pointer to float, a pointer to string and a pointer to bool.
// the function makes a copy of each pointer but they point to the same address as the originals
func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
    //changing the values the pointers point to is seen outside the function
    *quantity = 3
    *price = 500.5
    *name = "Mobile Phone"
    *sold = false
}
 
// declaring struct type
type Product struct {
    price       float64
    productName string
}
 
// declaring a function that takes in a struct value and modifies it
func changeProduct(p Product) {
    p.price = 300
    p.productName = "Bicycle"
    // the changes are not seen to the outside world
}
 
// declaring a function that takes in a pointer to struct value and modifies the value
func changeProductByPointer(p *Product) {
    (*p).price = 300
    p.productName = "Bicycle"
    // the changes are seen to the outside world
 
}
 
// declaring a function that takes in a slice
func changeSlice(s []int) {
    for i := range s {
        s[i]++
    }
    // the changes are seen to the outside world
}
 
// declaring a function that takes in a map
func changeMap(m map[string]int) {
    m["a"] = 10
    m["b"] = 20
    m["x"] = 30
    // the changes are seen to the outside world
}
 
func main() {
 
    // declaring some variables
    quantity, price, name, sold := 5, 300.2, "Laptop", true
    fmt.Println("BEFORE calling changeValues():", quantity, price, name, sold)
    // => BEFORE calling changeValues(): 5 300.2 Laptop true
 
    // invoking the function has no effect on the variables.
    // the function works on and modifies copies, not originals.
    changeValues(quantity, price, name, sold)
    fmt.Println("AFTER calling changeValues():", quantity, price, name, sold)
    // => AFTER calling changeValues(): 5 300.2 Laptop true
 
    // the function modifies the values.
    changeValuesByPointer(&quantity, &price, &name, &sold)
    fmt.Println("AFTER calling changeValuesByPointer():", quantity, price, name, sold)
    // => AFTER calling changeValuesByPointer(): 3 500.5 Mobile Phone false
 
    // declaring a struct value
    present := Product{
        price:       100,
        productName: "Watch",
    }
 
    // invoking the function has no effect on the struct value.
    // the function works on and modifies a copy, not the original.
    changeProduct(present)
    fmt.Println(present) // => {100 Watch}
 
    // the function modifies the struct value.
    changeProductByPointer(&present)
    fmt.Println("AFTER calling changeProductByPointer:", present)
    // => AFTER calling changeProductByPointer: {300 Bicycle}
 
    // declaring a slice
    prices := []int{10, 20, 30}
 
    // When a function changes a slice or a map the actual data is changed.
    changeSlice(prices)
    fmt.Println("prices slice after calling changeSlice():", prices)
    // => prices slice after calling changeSlice(): [11 21 31]
 
    // declaring a map
    myMap := map[string]int{"a": 1, "b": 2}
    // When a function changes a slice or a map the actual data is changed.
 
    changeMap(myMap)
    fmt.Println("myMap after calling changeMap():", myMap)
    // => myMap after calling changeMap(): map[a:10 b:20 x:30
 
    // slices and maps are not meant to be used with pointers.
}
```

## Pointers exercises

### Pointer ex1

```go
package main

import "fmt"

func main() {
	// Consider the following variable declaration
	x := 10.10

	// 1. Print out the address of x
	fmt.Printf("%p\n", &x)

	// 2. Declare a pointer called ptr that stores the address of x.
	ptr := &x

	// 3. Print out the type and the value of ptr.
	fmt.Printf("%T %v\n", ptr, ptr)

	// 4. Print the address of the pointer and the value of x though the pointer (use the dereferencing operator).
	fmt.Printf("%T %v\n", ptr, *ptr)
}
```

### Pointer ex2

```go
package main

import "fmt"

func main() {
	// Consider the following variable declarations:
	x, y := 10, 2
	ptrx, ptry := &x, &y

	// Declare a new variable called z and initialize the variable by dividing x by y through the pointers.
	z := *ptrx / *ptry
	fmt.Printf("z = %v", z)
}
```

### Pointer ex3

```go
package main

import "fmt"

func main() {
	// Consider the following variable declaration:
	x, y := 5.5, 8.8

	// Write a function that swaps the values of x and y.
	// After calling the function x will be 8.8 and y will 5.5

	fmt.Printf("x = %f, y = %f\n", x, y)
	swap(&x, &y)
	fmt.Printf("x = %f, y = %f\n", x, y)
}

func swap(a, b *float64) {
	*a, *b = *b, *a
}
```

## OOP

### Receiver

```go
package main
 
import (
    "fmt"
    "time"
)
 
// declaring a new type
type names []string
 
// declaring a method (function receiver)
func (n names) print() {
    // n is called method's receiver
    // n is the actual copy of the names we're working with and is available in the function.
    // n is like this or self from OOP.
    // any variable of type names can call this function on itself like variable_name.print()
 
    // iterating and printing all names
    for i, name := range n {
        fmt.Println(i, name)
    }
}
 
func main() {
    // Go doesn't have classes, but you can define methods on defined types.
    // a type may have a method set associated with it which enhances the type with extra behaviour.
 
    const day = 24 * time.Hour
    fmt.Printf("%T\n", day) // it's type is time.Duration
 
    // calling a method on time.Duration type
    // Seconds() is a method aka a receiver function.
    seconds := day.Seconds()
 
    // Seconds() returns the duration as a floating point number of seconds.
    fmt.Printf("%T\n", seconds)               //its type is float64
    fmt.Println("Seconds in a day:", seconds) // Seconds in a day: 86400
 
    // declaring a value of type names
    friends := names{"Dan", "Marry"}
    // calling the print() method on friends variable
    friends.print()
 
    // another way to call a method on a type
    names.print(friends) // not very common
}
```

### Pointer receivers

```go
package main

import "fmt"

// declaring a new struct type
type car struct {
	brand string
	price int
}

// declaring a method for car type
// it changes the value it works on
func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
	// the changes are not seen to the outside world (pass by value)
}

// declaring a method with a pointer receiver
func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
	// the changes are seen the outside world
}

// method declarations are not permitted on named types that are themselves pointer types
type distance *int

// ERROR ->  invalid receiver type *distance (distance is a pointer type)
// func (d *distance) f() {
//  fmt.Println("say something")
// }

func main() {
	// declaring a car value
	myCar := car{brand: "Audi", price: 40000}

	// calling the method with a value receiver
	myCar.changeCar1("Opel", 21000)

	// no change due to the same pass by value mechanism  !!!
	fmt.Println(myCar) // => {Audi 40000}

	// calling the method with a pointer receiver. It changes the value!
	(&myCar).changeCar2("Seat", 25000)
	fmt.Println(myCar) // -> {Seat 25000}

	// declaring a pointer to car
	var yourCar *car
	yourCar = &myCar      // it stores the address of myCar
	fmt.Println(*yourCar) // -> {Seat 25000}

	// calling the method on pointer type
	// valid ways to call the method:
	yourCar.changeCar2("VW", 30000)
	(*yourCar).changeCar2("VW", 30000)
	fmt.Println(*yourCar) // => {VW 30000}

	// in the above example both myCar and yourCar variables have been modified.
	fmt.Println(myCar) // => {VW 30000}
}
```

### Interfaces

```go
package main
 
import (
    "fmt"
    "math"
)
 
// declaring an interface type called shape
// an interface contains only the signatures of the methods, but not their implementation
type shape interface {
    area() float64
    perimeter() float64
}
 
// declaring 2 struct types that represent geometrical shapes: rectangle and circle
 
type rectangle struct {
    width, height float64
}
type circle struct {
    radius float64
}
 
// method that calculates circle's area
func (c circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}
 
// method that calculates rectangle's area
func (r rectangle) area() float64 {
    return r.height * r.width
}
 
// method that calculates circle's perimeter
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
 
// method that calculates rectangle's perimeter
func (r rectangle) perimeter() float64 {
    return 2 * (r.height + r.width)
}
 
// any type that implements the interface is also of type of the interface
// rectangle and circle values are also of type shape
func print(s shape) {
    fmt.Printf("Shape: %#v\n", s)
    fmt.Printf("Area: %v\n", s.area())
    fmt.Printf("Perimeter: %v\n", s.perimeter())
}
 
func main() {
 
    // declaring a circle and a rectangle values
    c1 := circle{radius: 5}
    r1 := rectangle{width: 3, height: 2}
 
    // circle and rectangle both implements the geometry interface  because they implement all methods of the interface
    // an interface is implicitly implemented in Go
    print(c1)
    print(r1)
 
    a := c1.area()
    fmt.Println("Circle Area:", a)
}
```

### Type assertions

```go
package main
 
import (
    "fmt"
    "math"
)
 
// declaring an interface type called shape
type shape interface {
    area() float64
    perimeter() float64
}
 
// declaring a struct type
type rectangle struct {
    width, height float64
}
 
// declaring a struct type
type circle struct {
    radius float64
}
 
// declaring a method for circle type
func (c circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}
 
// declaring a method for circle type
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
 
// declaring a method for circle type
func (c circle) volume() float64 {
    return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}
 
// declaring a method for rectangle type
func (r rectangle) area() float64 {
    return r.height * r.width
}
 
// declaring a method for rectangle type
func (r rectangle) perimeter() float64 {
    return 2 * (r.height + r.width)
}
 
// declaring a function that takes an interface value
func print(s shape) {
 
    fmt.Printf("Shape: %#v\n", s)
    fmt.Printf("Area: %v\n", s.area())
    fmt.Printf("Perimeter: %v\n", s.perimeter())
}
 
func main() {
 
    // declaring an interface value that holds a circle type value
    var s shape = circle{radius: 2.5}
 
    fmt.Printf("%T\n", s) //interface dynamic type is circle
 
    // no direct access to interface's dynamic values
    // s.volume() -> error
 
    // there is access only to the methods that are defined inside the interface
    fmt.Printf("Circle Area:%v\n", s.area())
 
    // an interface value hides its dynamic value.
    // use type assertion to extract and return the dynamic value of the interface value.
    fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())
 
    // checking if the assertion succeded or not
    ball, ok := s.(circle)
    if ok == true {
        fmt.Printf("Ball Volume:%v\n", ball.volume())
    }
 
    //** TYPE SWITCHES **//
 
    // it permits several type assertions in series
    switch value := s.(type) {
    case circle:
        fmt.Printf("%#v has circle type\n", value)
    case rectangle:
        fmt.Printf("%#v has rectangle type\n", value)
 
    }
}
```

### Extending interfaces

```go
package main

import (
	"fmt"
	"math"
)

// declaring an interface type called shape
type shape interface {
	area() float64
	perimeter() float64
}

// declaring a struct type
type rectangle struct {
	width, height float64
}

// declaring a struct type
type circle struct {
	radius float64
}

// declaring a method for circle type
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// declaring a method for circle type
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// declaring a method for circle type
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// declaring a method for rectangle type
func (r rectangle) area() float64 {
	return r.height * r.width
}

// declaring a method for rectangle type
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// declaring a function that takes an interface value
func print(s shape) {

	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {

	// declaring an interface value that holds a circle type value
	var s shape = circle{radius: 2.5}

	fmt.Printf("%T\n", s) //interface dynamic type is circle

	// no direct access to interface's dynamic values
	// s.volume() -> error

	// there is access only to the methods that are defined inside the interface
	fmt.Printf("Circle Area:%v\n", s.area())

	// an interface value hides its dynamic value.
	// use type assertion to extract and return the dynamic value of the interface value.
	fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())

	// checking if the assertion succeded or not
	ball, ok := s.(circle)
	if ok == true {
		fmt.Printf("Ball Volume:%v\n", ball.volume())
	}

	//** TYPE SWITCHES **//

	// it permits several type assertions in series
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)

	}
}
```

### Empty interface

```go
package main
 
import "fmt"
 
// declaring an empty interface
type emptyInterface interface {
}
 
// declaring a new struct type which has one field of type empty interface
type person struct {
    info interface{}
}
 
func main() {
 
    // declaring an empty interface value
    var empty interface{}
 
    // an empty interface may hold values of any type
    // storing an int in the empty interface
    empty = 5
    fmt.Println(empty) // => 5
 
    // storing a string in the empty interface
    empty = "Go"
    fmt.Println(empty) // => Go
 
    // storing a slice in the empty interface
    empty = []int{2, 34, 4}
    fmt.Println(empty) // => [2 34 4]
 
    // fmt.Println(len(empty)) -> error, and it doesn't work!
 
    // retrieving the dynamic value using an assertion
    fmt.Println(len(empty.([]int))) // => 3
 
    // declaring person value
    you := person{}
 
    // assigning any value to empty interface field
    you.info = "You name"
    you.info = 40
    you.info = []float64{4.5, 6., 8.1}
 
    fmt.Println(you.info)
}
```

## Methods and interfaces exercises

### Method exercises

#### Methods ex1

```go
package main

import "fmt"

// 1. Create a new type called money. Its underlying type is float64.
type money float64

// 2. Create a method called print that prints out the money value with exact 2 decimal points.
func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func main() {
	var m money = 1.127
	m.print()
}
```

#### Methods ex2

```go
package main

import "fmt"

// Consider the money type declared at exercise #1.
// Create a new method for the money type called printStr
// that returns the money value as a string with 2 decimal points.

type money float64

// 2. Create a method called print that prints out the money value with exact 2 decimal points.
func (m money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

func main() {
	var m money = 1.127
	fmt.Println(m.printStr())
}
```

#### Methods ex3

```go
package main

import "fmt"

// 1. Create a new struct type called book with 2 fields:
// price and title of type float64 and string.
type book struct {
	price float64
	title string
}

// 2. Create a method for the newly defined type called
// vat that returns the vat value if vat is 9%.
func (b book) vat() float64 {
	return b.price * 0.09
}

// 3. Create a method called discount that takes a book value as receiver and
// decreases the price of the book by 10%.
func (b book) discount() float64 {
	return b.price * 0.9
}

func main() {
	book := book{title: "The Trial", price: 9.9}
	fmt.Printf("Vat:%v\n", book.vat())
	fmt.Printf("%#v\n", book.discount())
}
```

#### Methods ex4

```go
package main

import "fmt"

type book struct {
	title string
	price float64
}

// method for book type
func (b *book) changePrice(p float64) {
	b.price = p
}

func main() {
	// book value
	bestBook := book{title: "The Trial by Kafka", price: 9.9}

	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price)
}
```

### Interface exercises

#### Interface ex1

```go
package main

import "fmt"

// Consider the following type and interface declaration.

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

// 1. Create a Go program where car type implements the vehicle interface.

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {
	// 2. Create a variable of type vehicle that holds a car struct value.
	var v vehicle = car{
		licenceNo: "l1",
		brand:     "car1",
	}

	// 3. Call the methods (Licence and Name) on the interface value declared at step 2
	fmt.Println(v.License())
	fmt.Println(v.Name())
}
```

#### Interface ex2

```go
package main

import "fmt"

func main() {
	// 1. Declare a variable called empty of type empty interface. Print out its type.
	var empty any
	fmt.Printf("%T\n", empty)

	// 2. Assign an int value to the variable called empty. Print out its type.
	empty = 42
	fmt.Printf("%T\n", empty)

	// 3. Assign a float64 value to empty. Print out its type.
	empty = 42.42
	fmt.Printf("%T\n", empty)

	// 4. Assign an int slice value to empty. Print out its type.
	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)

	// 5. Type Assertion
	empty = append(empty.([]int), 10)

	// 6. Print out its type.
	fmt.Printf("%v\n", empty)
}
```

#### Interface ex3

```go
package main

import "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x any
	x = cube{edge: 5}
	// There is an error in the following
	// v := volume(x)

	c, ok := x.(cube)
	if ok {
		fmt.Printf("Cube Volume: %.2f\n", volume(c))
	}
}
```

## Concurrency

### Concurrency vs parallelism

[Proverb](https://go.dev/blog/waza-talk)
[Go blog](https://opensource.googleblog.com/2009/11/hey-ho-lets-go.html)

### Goroutines intro

```go
package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("Do something...")
}

func main() {
	doSomething()
	go doSomething()

	time.Sleep(2 * time.Second)
}
```

### Go keyword

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1() start")
	for i := range 3 {
		fmt.Println("f1() i =", i)
	}
	fmt.Println("f1() end")
}

func f2() {
	fmt.Println("f1() start")
	for i := range 8 {
		fmt.Println("f1() i =", i)
	}
	fmt.Println("f1() end")
}

func main() {
	fmt.Println("Main execution started")

	fmt.Println("Logical CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("Max Procs:", runtime.GOMAXPROCS(0))

	go f1()
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	go f2()
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Millisecond * 200)
	fmt.Println("Main execution ended")
}
```

### Wait groups

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func f1(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("f1() start")
	for i := range 3 {
		fmt.Println("f1() i =", i)
	}
	fmt.Println("f1() end")
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("f2() start")
	for i := range 8 {
		fmt.Println("f2() i =", i)
	}
	fmt.Println("f2() end")
}

func main() {
	fmt.Println("Main execution started")

	var wg sync.WaitGroup

	wg.Add(1)
	go f1(&wg)
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Add(1)
	go f2(&wg)
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Main execution ended")
}
```

### URL checker

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		fmt.Printf("%s is dow\n", url)
		return
	}

	defer resp.Body.Close()
	fmt.Printf("%s -> %d\n", url, resp.StatusCode)
	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		file := strings.Split(url, "//")[1]
		fmt.Printf("Writing response to %s\n", file)
		err = os.WriteFile(file, body, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	urls := []string{"https://mariolazzari.it", "https://mariafilippini.it", "https://golang.org"}
	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
	}

	fmt.Printf("Goroutines running: %d\n", runtime.NumGoroutine())

	wg.Wait()
}
```

### Data race

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100
	// declaring a WaitGroup to synchronize the goroutines with the main function.
	var wg sync.WaitGroup

	// adding 200 goroutines to the WaitGroup
	wg.Add(gr * 2)

	// declaring a shared value
	var n int = 0

	// starting 200 goroutines
	for range gr {

		// each goroutine is an anonymous function
		go func() {
			// notifying the WaitGroup that the goroutine is done
			defer wg.Done()
			// introducing some artificial time
			time.Sleep(time.Second / 10)
			// increment the shared value
			n++
		}()

		// goroutine that decrements the shared value
		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 10)
			n--
		}()

	}
	// waiting for the goroutines to terminate.
	wg.Wait()

	//  printing the final value of n
	fmt.Println(n) // it will have another value for each program execution -> DATA RACE
}
```

### Race detextor

```sh
go run --race main.go
```

### Mutex

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100
	var wg sync.WaitGroup

	wg.Add(gr * 2)

	var n int = 0

	// starting 200 goroutines with mutex
	var mx sync.Mutex

	for range gr {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 10)

			mx.Lock()
			n++
			mx.Unlock()
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Second / 10)

			mx.Lock()
			defer mx.Unlock()
			n--
		}()

	}
	wg.Wait()
	fmt.Println(n)
}
```

### Channels

```go
package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println(ch)

	// send operation
	go func() {
		ch <- 42
	}()

	// receive operation
	n := <-ch

	fmt.Println("n =", n)
	// fmt.Println("n =", <-ch)

	// send only
	c1 := make(chan<- string)
	// receive only
	c2 := make(<-chan string)

	fmt.Printf("%T, %T, %T\n", ch, c1, c2)

	// goroutines and channels
	go f1(50, ch)
	n = <-ch
	fmt.Println("n from f1 ch:", n)

	// close channel
	close(ch)
}
```

### Goroutines and channels

```go
package main

import "fmt"

func factorial(n int, ch chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	ch <- f
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go factorial(5, ch)

	// wait from the channel
	fmt.Println(<-ch)

	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		fmt.Println(<-ch)
	}
}
```

### Anonymous fanctions

```go
package main

import "fmt"

func factorial(n int, ch chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	ch <- f
}

func main() {
	ch := make(chan int)

	for i := 5; i <= 15; i++ {
		go func(n int) {
			factorial(n, ch)
		}(i)

		fmt.Printf("%d! = %d\n", i, <-ch)
	}
}
```

### URL checker with channels

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {

		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v  \n", err)

		// Sending the string over the channel.
		// This is a blocking call, so this goroutine will
		// wait for the main goroutine to receive it on the other part of the channel.
		c <- s
	} else {

		s := fmt.Sprintf("Status Code: %d  \n", resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			file := strings.Split(url, "//")[1]
			file += ".txt"

			s += fmt.Sprintf("Writing response Body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				s += "Error writing to file!\n"
				// sending s over the channel
				c <- s
			}
		}
		s += fmt.Sprintf("%s is UP\n", url)

		// sending s over the channel
		c <- s
	}
}

func main() {
	urls := []string{"https://www.golang.org", "https://www.google.com", "https://www.medium.com"}

	// Declaring a new channel
	c := make(chan string)

	for _, url := range urls {
		// Launching the goroutines
		go checkAndSaveBody(url, c)
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => 4

	// Receiving the messages from the channel
	for range urls {
		fmt.Println(<-c)
	}
}
```

### URL checker with range

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkURL(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v  \n", err)
		fmt.Println(s)
		c <- url
	} else {
		s := fmt.Sprintf("Status Code: %d  \n", resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)
		fmt.Println(s)
		c <- url
	}
}

func main() {

	urls := []string{"https://www.golang.org", "https://www.google.com", "https://www.medium.com"}

	c := make(chan string)

	for _, url := range urls {
		// the goroutine is sending the URL over channel for the next goroutine to use it.
		go checkURL(url, c)
	}

	time.Sleep(time.Second * 2)

	// Spawning goroutines continously

	// 1.
	// for {
	//  go checkUrl(<-c, c)

	// OR

	// 2.
	// for url := range c { // It means "wait for the channel to return some values"
	//  fmt.Println(strings.Repeat("#", 30))
	//  time.Sleep(2 * time.Second) // pause for 2 seconds
	//  go checkUrl(url, c)
	// }

	// OR

	// 3.
	for url := range c {
		go func(u string) {
			time.Sleep(2 * time.Second) // pause for 2 seconds
			checkURL(u, c)
		}(url)
	}
}
```

### Unbuffered channel

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered
	ch1 := make(chan int)
	// buffered
	ch2 := make(chan int, 3)
	_ = ch2

	go func(c chan int) {
		fmt.Println("Start goroutine")
		fmt.Println("After sending to channel")
		c <- 10 // blocks main until receving from channel
		fmt.Println("End goroutine")
	}(ch1)

	fmt.Println("Main sleeps")
	time.Sleep(time.Second)

	n := <-ch1
	fmt.Println("Main from channel:", n)
}
```

### Buffered channels

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	go func(c chan int) {
		for i := range 5 {
			fmt.Printf("Start goroutine %d\n", i+1)
			fmt.Println("After sending to channel")
			c <- i + 1
			fmt.Printf("End goroutine %d\n", i+1)
		}
		close(c) // all data sent
	}(ch)

	//fmt.Println("Main sleeps")
	//time.Sleep(time.Second)

	// receive fron data sending channel
	for v := range ch { // v:=<-c
		fmt.Println("Main received:", v)
	}

	// zero value after closing
	fmt.Println("Main closed ch:", <-ch)

	// c <-10 panic!
}
```

### Select

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "Ciao"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "Hello"
	}()

	time.Sleep(time.Second)

	// wait from multiple channels
	for range 2 {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			fmt.Println("No activity")
		}
	}

	fmt.Printf("Exec time: %v\n", time.Since(start))
}
```

## Concurrency exercises

### Wait groups and mutex exercises

####  Wait groups and mutex ex1

```go
package main

import (
	"fmt"
	"sync"
)

func sayHello(n string) {
	fmt.Printf("Hello, %s!\n", n)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello("Mr. Wick")
	}()

	wg.Wait()
}
```

####  Wait groups and mutex ex2

```go
package main

import (
	"fmt"
	"sync"
)

// 1. Create a function called sum() that calculates and then prints out  the sum of 2 float numbers it receives as arguments.
// Format the result with 2 decimal points.
func sum(a, b float64, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
}

func main() {

	// 2. From main launch 3 goroutines that execute the function you have just created (sum)
	// 3. Synchronize the goroutines and the main function using WaitGroups

	var wg sync.WaitGroup

	wg.Add(3)

	go sum(1, 2, &wg)
	go sum(4, 5, &wg)
	go sum(0, -1, &wg)

	wg.Wait()
}
```

####  Wait groups and mutex ex3

```go
package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	//	1. Create an anonymous function that calculates and prints out the square root of a float value it receives as argument.
	go func(f float64, wg *sync.WaitGroup) {
		defer wg.Done()
		x := math.Sqrt(f)
		fmt.Printf("Square root of %.2f is %.4f\n", f, x)
	}(16.1, &wg)

	// 2. Launch the function as a goroutine and synchronize it with main using WaitGroups
	wg.Wait()
}
```

####  Wait groups and mutex ex4

```go
package main

import (
	"fmt"
	"math"
	"sync"
)

// Change the code from Exercise #3 and launch 50 goroutines that calculate concurrently
// the square root of all the numbers between 100 and 149 (both included).

func main() {
	var wg sync.WaitGroup
	wg.Add(50)

	for i := 100; i <= 149; i++ {
		go func(n int, wg *sync.WaitGroup) {
			defer wg.Done()
			x := math.Sqrt(float64(n))
			fmt.Printf("Square root of %d is %.4f\n", n, x)
		}(i, &wg)

	}
	wg.Wait()
}
```

####  Wait groups and mutex ex5

```go
package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()

	mx.Lock()
	defer mx.Unlock()
	*b += n
	//wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()

	mx.Lock()
	defer mx.Unlock()
	*b -= n
	// wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	wg.Add(200)

	balance := 100

	for i := range 100 {
		go deposit(&balance, i, &wg, &mx)
		go withdraw(&balance, i, &wg, &mx)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
```

### Goroutines and channels exercises

#### Goroutines and channels ex1

```go
package main

import (
	"fmt"
	"sync"
)

func sayHello(n string) {
	fmt.Printf("Hello, %s!\n", n)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello("Mr. Wick")
	}()

	wg.Wait()
}
```

#### Goroutines and channels ex2

```go
package main

import "fmt"

// Create a function literal (a.k.a. anonymous function)
// that sends the string value if receives as argument
// to main func using a channel.
func main() {
	ch := make(chan string)

	go func() {
		ch <- "Ciao"
	}()

	fmt.Println(<-ch)
}
```

#### Goroutines and channels ex3

```go
package main

import (
	"fmt"
)

func main() {
	// c := make(<-chan int)
	c := make(chan int)

	go func(n int) {
		c <- n
	}(100)

	fmt.Println(<-c)
}
```

#### Goroutines and channels ex4

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Create a goroutine named power() that has one parameter of type int,
// calculates the square value of its parameter
// and then sends  the result into a channel.
func power(n int, ch chan<- int) {
	ch <- n * n
}

func main() {
	const jobs = 50
	ch := make(chan int, jobs)

	var wg sync.WaitGroup
	wg.Add(jobs)

	// In the main function launch 50 goroutines that calculate
	// the square values of all numbers between 1 and 50 included.
	for i := 1; i <= jobs; i++ {
		go func(n int) {
			defer wg.Done()
			power(n, ch)
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println(runtime.NumGoroutine())

	// Print out the square values.
	for val := range ch {
		fmt.Println(val)
	}
}
```

#### Goroutines and channels ex5

```go
package main

import "fmt"

// Change the program from Exercise #4 and calculate
// the square of all values between 1 and 50 included
// using an anonymous function.

func main() {
	ch := make(chan int)

	for i := 1; i <= 100; i++ {
		go func(x int) {
			ch <- x * x
		}(i)
	}

	for i := 1; i <= 100; i++ {
		fmt.Println(<-ch)
	}
}
```

## Packages and Modules

### Create a Package
