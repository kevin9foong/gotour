package main 

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const ( 
	apple_cost = 1.00
	orange_cost = 1.50
)

func count_to_n(n int) { 
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

func count_to_n_2(n int) {
	count := 1 
	for count <= n {  // or for ; count <= n; { -> basically a while loop (there are no while loops in go)
		fmt.Println(count)
		count++
	}
}

func infinite_loop() { 
	for { // concise way of writing an infinite loop
		fmt.Println("infinite loop")
	}
}

func is_balance_sufficient(cost, balance float64) bool { 
	if (cost > balance) {
		return false
	}
	return true
}

func yes_or_no(value bool) string { 
	if yes, no, maybe := "yes", "no", "maybe"; value { // just like a for loop, you can include a statement in the first clause (most used to declare a variable) in the if statement
		return yes
	} else if 1 != 1 {
		return maybe
	} else { 
		return no // the variables declared in the if statement is also available in the else if and else blocks 
	}
}

// newtons method for finding square root
func sqrt (x float64) float64 { 
	z := 1.0 
	for i, prevZ := 0, z; i < 10; i++ { // 10 iterations with prevZ to track the previous value
		z = z - ((z*z - x) / (2 * z))
		if math.Abs(z-prevZ) < 1e-6 { // terminate if the change is small
			break
		}
		prevZ = z
		fmt.Printf("Guess number %d: z = %f\n", i, z)
	}
	return z
}

func get_mysterious_item_name() string {
	if (rand.Intn(2) == 0) {
		return "banana"
	}
	return "grape"
}

func get_price (item string) { 
	// fall through does not happen in go, unlike in other languages. break; is implicit
	switch item { 
	case "apple": 
		fmt.Println("$", apple_cost)
	case "orange": 
		fmt.Println("$", orange_cost)
	case get_mysterious_item_name(): // switch statements do not have to be constants unlike in other languages. each case statement is only executed if the above cases are not matched. 
		fallthrough // fallthrough is used to execute the next case block and must be explicitly stated 
	case "seasonal_item": 
		fmt.Println("seasonal item or mysterious item")
	default: 
		fmt.Println("Unknown item")
	}
	return 
}

func print_occasion() {
	switch { // switch statements can also be used without a variable. this is same as switch true 
		// very useful for long if-else chains 
	case time.Now().Month() == time.March:
		fmt.Println("It's March!")
	case time.Now().Month() == time.October:
		fmt.Println("It's October!")
	default:
		fmt.Println("It's neither March nor October. It is", time.Now().Month())
	}
}

func hello_world() {
	defer fmt.Println("world!") // the arguments are evaluated immediately but the function call is deferred until the surrounding function returns
	fmt.Printf("hello ")
}

// when to use defer? 
// note that defer uses a stack. defer functions are executed in FILO (first in last out) order
func write_files() {
	// open file A and read it 
	// defer fileA.Close()
	// open file B and write to it 
	// defer fileB.Close()

	// this ensures that the files are closed even if an error occurs eg, after file A is opened but file B failure occurs
}

func main() { 
	fmt.Println("count_to_n")
	count_to_n(10)
	fmt.Println("count_to_n_2")
	count_to_n_2(10)

	fmt.Printf("Can i afford an apple with costs $%0.2f with balance of $%0.2f? The answer is %s\n", apple_cost, 1.00, yes_or_no(is_balance_sufficient(apple_cost, 1.00)))

	fmt.Printf("The square root of 100 is %0.8f\n", sqrt(100))

	get_price("apple")
	get_price("orange")
	get_price("banana")

	print_occasion()

	hello_world()
}
