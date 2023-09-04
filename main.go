package main

import "fmt"

var steps int

func main() {

	ll := LinkedList{}

	fmt.Println("empty  : ", ll.isEmpty())

	ll.append(3)
	ll.append(1)
	ll.append(4)
	ll.append(9)
	ll.prepend(22)

	fmt.Println("count  : ", ll.count())
	fmt.Println("first  : ", ll.first())
	fmt.Println("last   : ", ll.last())
}
