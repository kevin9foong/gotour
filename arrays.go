package main 

import ( 
	"fmt"
)

// arrays cannot be resized, the size is part of the type
var a [10]int // array of 10 integers

func main() { 
	for i := 0; i < len(a); i++ { 
		a[i] = i 
	} 

	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // same { } notation as struct init 
	fmt.Println(primes)

	// slices are dynamic in size, they can grow and shrink
	// slices are based the array -> does not store data itself, just describes a section of underlying array
	// reference the same address as the underlying array
	// used much more commonly than arrays
	// array[low inclusive:high exclusive]
	s := a[1:4] // slice of primes from index 1 to 3
	fmt.Println(s)
	s2 := a[1:4]

	s[0] = 100 // this changes the underlying array
	fmt.Println(s2)
	fmt.Println(a)

	// array literals 
	arr_literal := [3]bool{true, false, true}
	fmt.Println(arr_literal)

	slice_literal := []int{1, 2, 3} // creates an underlying array, then creates a slice reference to it
	fmt.Println(slice_literal)

	// can make slice of structs 
	slice_of_squares := []struct { 
		side int 
	}{
		{1},
		{2},
		{3},
	}
	fmt.Println(slice_of_squares)

	// if dont provide the lower and upper bound of slice, lower = 0, upper = len(array)
	s3 := a[:] // slice of all elements
	fmt.Println(s3)

	// slice has a length and a capacity
	// length is the number of elements in the slice
	// capacity is the number of elements in the underlying array, starting from the first element in the slice. 
	// note that capacity cannot change since arrays are fixed size 
	
	fmt.Println(len(s3), cap(s))
	fmt.Println("Setting s3 to a[1:4]")
	s3 = a[1:4]
	fmt.Println(len(s3), cap(s3))
	// we can see that the length of s3 changed -> hence slice is dynamic in size

	var slice []int // has a zero value of nil which has len and capacity of 0 and does not have an underlying array
}