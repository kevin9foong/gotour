package main // every go program starts with a package declaration, it is a way to group and reuse related files

import ( 
	"fmt"
	"math"
	"math/rand" // by convention, the last word is the package name eg, 'math/rand' has the package name 'rand'
) // this is a factored import statement -> it is a good practice to use it

// we can factor variable declarations into a block also 
var (
	fruit string = "apple" 
	count int = 10 
)

var ( 
	zero int 
	empty_string string
	false_boolean bool
)

const ( 
	pi = 3.14 // untyped constant
	name string = "John" // typed constant
)

var pi_float float64 = pi // based on the context, the untype constant has its type inferred 

// kinda looks like typescript 
func add(a int, b int) int { 
	return a + b
}

func subtract (a, b int) int {
	return a - b
}

func add_and_subtract(a, b int) (addition int, subtraction int) {
	var i = 0 // this is a local variable, it is not a return value
	addition = add(a, b) + i
	subtraction = subtract(a, b) + i
	return // naked return value -> returns the named return values, should be avoided for complex functions
}

func int_to_float (a int) float64 {
	return float64(a)
}

func main() {
	fmt.Println("Hello, World! I am a random number: ", rand.Intn(10))
	sum, diff := add_and_subtract(1, 2) // declares a variable using short variable declaration syntax which is only available inside functions
	// outside functions, you need to declare a variable using the var keyword
	fmt.Println("The sum and difference of 1 and 2 are: ", sum, ",", diff)
	fmt.Println("Zero values declared: ", zero, ",", empty_string, ",", false_boolean)
	fmt.Println("The float value of 2 is: ", math.Sqrt(int_to_float(int(2)))) // using Pow to calculate square root
	fmt.Printf("I am of type %T\n", 1.0) // the type depends on the precision -> 1 is int, 1.0 is float64, 1 + 0.0i is complex128
	fmt.Printf("I am of type %T\n", "a") // string 
	var myChar rune = 'A' // Declaring a rune variable - rune is a synonym for int32
	fmt.Printf("I am of type %T\n", myChar) // int32
}

// go run helloworld.go -> this is how you run a go program
// go build helloworld.go -> this is how you build a go program

// supported basic types: 
// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// note: int follows the platform's word size -> use int unless you have a specific reason to use a smaller or larger integer type

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128 

// if not initialized, the variable is set to its zero value
// the zero value is: 
// 0 for numeric types
// false for boolean type
// "" for strings

// Type conversions must be explicit
// eg, var i int = 10 -> this is an int
// var j float64 = float64(i) -> this is a float64

// http://127.0.0.1:3999/flowcontrol/1