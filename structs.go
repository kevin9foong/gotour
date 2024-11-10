package main 

import ( 
	"fmt"
)

var p *int // p is a pointer to an int, zero value is nil
var i int = 42 // i is an int

// using structs to create a linked list node
type Node struct { 
	value int 
	next *Node
}

type LinkedList struct { 
	head *Node
}

func (l *LinkedList) add(value int) { 
	newNode := &Node{value: value}
	if l.head == nil { 
		l.head = newNode // access struct fields with dot notation
	} else { 
		current := l.head
		for current.next != nil { 
			current = current.next
		}
		current.next = newNode
	}
}

type Rectangle struct { 
	width int 
	height int 
}

func main() {
	// & is the address operator, it returns the address of the variable
	p = &i // p now points to i 
	fmt.Println(p) // prints the address of i

	// since p is a pointer aka reference to address of i 
	// we can de-reference it using * 
	
	fmt.Println(*p) // prints the value of i
	*p = 21 // this changes the value of i
	fmt.Println(i) // prints the new value of i

	// note that cannot do arithmetic with pointers in Go unlike C/C++ 
	// eg, cannot do p++ or p = p + 7  -> in C, p++ increments the pointer address by the size of the type of the pointer

	// given a pointer to a struct 
	node := Node{value: 10, next: nil}
	address := &node 

	// we can access the fields via 
	fmt.Println(node.value)
	fmt.Println((*address).value) // note that * has lower precedence than ., so we need to use parentheses
	fmt.Println(address.value) // or the shorthand which is the same as above without explicit dereference 

	r1 := Rectangle{10, 8}
	r2 := Rectangle{5, 0}
	r3 := Rectangle{height: 5}
	rp := &Rectangle{1, 2} // returns a pointer to the struct

	fmt.Println(r1, r2, r3, rp)
}
