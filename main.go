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
	assignSliceFromSlice()
	assignMap()
	assignStruct()
}

func assignStruct() {

	type user struct {
		Id        int
		FirstName string
		LastName  string
	}

	var u user
	u.Id = 1
	u.FirstName = "Hamza"
	u.LastName = "Hasan"
	fmt.Println(u)

	u1 := user{Id: 2, FirstName: "Hamza", LastName: "Hasan"}
	fmt.Println(u1)
}

func assignMap() {
	m := map[string]int{"Hey": 1}
	fmt.Println(m)
	fmt.Println(m["Hey"])

	m["hello"] = 2
	fmt.Println(m)

	delete(m, "Hey")
	fmt.Println(m)
}

func assignSliceFromSlice() {
	slc := []int{1, 2, 3, 4, 5, 6}

	slc1 := slc[2:]  // Specify starting index, no end index
	slc2 := slc[:1]  // Specify end index, no start index
	slc3 := slc[3:4] // Specify start & end index.

	fmt.Println(slc, slc1, slc2, slc3)
}

func assignSlice() {
	arr := [10]int{1, 2, 3, 4, 5, 6}
	slc := arr[:]
	fmt.Println(slc)

	fmt.Println("Appending to Slice")
	slc = append(slc, 49)

	fmt.Println(slc)
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
