package main

import "fmt"

func main() {
	type Person struct {
		name     string
		lastname string
		age      int
		fuck     []int //Slice
		suck     map[string]int
	}

	p1 := Person{
		name:     "Foo",
		lastname: "Bar",
		age:      20,
		fuck:     []int{2, 1},
		suck:     map[string]int{},
	}

	fmt.Println(p1)
	p1.suck["first"] = 92

	for _, value := range p1.fuck {
		fmt.Println(value)
	}

	for key, value := range p1.suck {
		fmt.Println(key, value)
	}
}
