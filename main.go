package main

import "fmt"

func main() {
	assignInt()
	assignFloat()
	assignString()
	assignComplex()
	assignPointer()
	assignArray()
	assignSlice()
}

func assignSlice() {
	panic("unimplemented")
}

func assignArray() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
}

func assignPointer() {
	var name *string = new(string)
	*name = "Hey"
	fmt.Println(*name)

}

func assignComplex() {
	fmt.Println("Assigning Complex")
	c := complex(3, 4)
	fmt.Println(c)
}

func assignString() {
	fmt.Println("Assigning String")
	name := "Name"
	fmt.Println(name)
}

func assignInt() {
	fmt.Println("Assigning Int")
	id := 50
	fmt.Println(id)
}

func assignFloat() {
	fmt.Println("Assigning Float")
	f := 4.21
	fmt.Println(f)
}
