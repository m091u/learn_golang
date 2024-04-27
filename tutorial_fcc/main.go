package main

import (
	"fmt"
	"strconv"
)

//inferring not working at a package level
var a  int = 42

// var block
var(
	age = 30
	name = "Naveen"
)

const (
	_ = iota  //generates a value but is not used so not saved in memory
	catDoc 
	dogDoc
	fishDoc
)

func main() {
    fmt.Println("Hello, Go!")

	var i int = 75
	i = 46
	fmt.Println(i)

	// initializing by inferring
	j := 56
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", age, age)

	//strings
	var l string
	l = string(i) //looks for what Unicode character is set at the value of i
	l = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", l, l)

	// bit shifting
	a := 8  // is 2^3
	fmt.Println(a << 3)  // 2^3  * 2^3 * 2^6  = 64
	fmt.Println(a >>3)  // 2^3 / 2 ^3 = 2 ^ 0 = 1

	//complex numbers
	var c complex64 = 5 + 12i  //parser understands i literal as an imaginary number
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))

	//strings
	s := "hello"
	fmt.Printf("%v, %T\n", s[2], s[2])  //
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	b := []byte(s)  //conversion to bytes, a collection of ASCII values, UTF value of each character
	fmt.Printf("%v, %T\n", b, b)


	//const bloc
	var docType int
	fmt.Printf("%v, %T\n", docType == catDoc, docType)
	fmt.Printf("%v, %T\n", docType == dogDoc, docType)
	fmt.Printf("%v, %T\n", dogDoc, dogDoc)

	//arrays
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v, %T\n", arr, arr)

	arr2 := [...]int{4, 5, 6} // can be used to skip spefifying the number of elements passed and that will be implied from the elements added in the {}
	fmt.Printf("%v, %T\n", arr2, arr2)
	fmt.Printf("Number of elements: %v\n", len(arr2))

	//array of arrays
	var arr3 [2][3]int = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	fmt.Printf("%v, %T\n", arr3, arr3)

	//slices
	//ways of creating slices
	var slice1 []int = []int{1, 2, 3}
	e := []int{1, 2, 3}
	f := e[:] // slice of all elements
	g := e[1:] // slice of all elements starting from index 1
	d := e[:2] // slice of first 2 elements
	fmt.Printf("%v, %v\n %v\n %v\n %v\n", slice1, e, f, g, d)

	e = append(e , 100)
	fmt.Printf("Value after appending %v\n", e)

	//removing elements from slices
	e = append(e[:2], e[3:]...)  //spread operator used to be able to concatenate slices
	fmt.Printf("Value after removing elements %v\n", e)

	//maps

	m := make(map[string]int)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Printf("%v, %T\n", m, m)
	_, ok := n["foo"]
	fmt.Printf("%v, %T\n", ok, ok)

	//struc
	type person struct {
		name string
		age int
	}
	p := person{
		name: "Naveen",
		age: 30,
	}
	fmt.Printf("%v, %T\n", p, p)

	//struct embedding
	type employee struct {
	    person
		company string
		role string
	}
	    
	emp := employee{
	    person: person{
	        name: "Naveen",
	        age: 30,
	    },
	    company: "Go tutorial",
	}
	emp.role= "author"
	fmt.Printf("%v, %T\n", emp, emp)
    fmt.Printf("%v, %T\n", emp.name, emp.name)
    
}