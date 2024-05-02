package main

import (
	"fmt"
	"unicode/utf8"
	"errors"
)

func main() {
	fmt.Println("Hello, World!")
	var intNum int64 // int64 can store larger numbers, the number is an indicator of how much memory is needed
	fmt.Println(intNum)

	var unitNum uint = 6 // positive numbers
	fmt.Println(unitNum)

	var floatNum float64 = 123456.7 // store numbers with decimal points
	fmt.Println(floatNum)

	// operations with mixed types not possible
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2

	var sum float32 = floatNum32 + float32(intNum32)
	fmt.Println(sum)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // result will be an int runded down
	fmt.Println(intNum1 % intNum2) // modulo to get remainder as well

	//String type
	var myString string = "Hello \nworld" // \n will insert a line break when using ""
	fmt.Println(myString)

	var myString2 string = `Hello
	    world` // ` ` will allow for multiple lines in a string
	fmt.Println(myString2)

	var concatString string = "Hello" + " " + "world"
	fmt.Println(concatString)

	//lengths of strings with chars outside vanilla ASCII
	fmt.Println(utf8.RuneCountInString("y"))

	// type can be skipped if it can be inferred from the right side
	var myString3 = "yellow"
	fmt.Println(myString3)

	//short-hand removing the var keyword and type
	myString4 := "green"
	fmt.Println(myString4)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	//constants cannot change value once created and need to be initialized
	const myConst string = "constant"
	fmt.Println(myConst)

	myFunction()

	var a int = 10
	var b int = 3
	var result , remainder, err  = intDivision(a, b)
	//if logic
	// if err!= nil{
	//     fmt.Printf(err.Error())
	// }else if remainder==0{
	// 	fmt.Printf("the result of the integer division is %v", result)
	// }else{fmt.Printf("the result of the integer division is %v and the remainder is %v", result, remainder)}

	//switch logic
	switch {
	case err!=nil:
		fmt.Printf(err.Error())
	case remainder==0:
		fmt.Printf("the result of the integer division is %v", result)
	default:
		fmt.Printf("the result of the integer division is %v and the remainder is %v", result, remainder)
	}
	
	switch remainder{
	case 0:
		fmt.Printf("\nThe division was exact")
	case 1,2:
		fmt.Printf("\nThe division was close")
	default:
		fmt.Printf("\nThe division was not close")
	}

	//arrays initialization
	

	var myArray [3]int = [3]int{1,2,3}
	myArray[0] = 10
	//reasign values
	myArray = [3]int{10,7,3}
	//short hand initialization
	newArray := [5]int{4,5,6}
	fmt.Print("\nthe array is:", myArray, "second array is:", newArray)

	//using slice and spread operator
	intArr := []int32{7,8,9}
	fmt.Println("\nthe slice is:", intArr)

	var intSlice []int32 = []int32{9,8,7}
	fmt.Printf("\nthe length is %v and the capacity is %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 10)
	fmt.Printf("\nthe length is %v and the capacity is %v and the array is %v", len(intSlice), cap(intSlice), intSlice)

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("\nthe length is %v and the capacity is %v and the array is %v", len(intSlice), cap(intSlice), intSlice)

	//using make
	makeSlice := make([]int32, 3, 5)
	fmt.Println(makeSlice)

	//map key: value pairs
	var myMap map[string]string = make(map[string]string)
	myMap["name"] = "mary"
	myMap["age"] = "25"
	myMap["hobby"] = "reading"
	fmt.Print("\nthis is the map:", myMap)

	var name , ok = myMap["name"]
	if ok{
		fmt.Printf("\nthe value of the key is %v", name)
	}else {
		fmt.Printf("\nthe key was not found")
	}

	delete(myMap, "age")
	fmt.Print("\nthis is the map after deletion:", myMap)

	//loops
	for attribute, value := range myMap{
	    fmt.Printf("\nDetails: %v is %v ", attribute, value)
	}

	for i, v := range intSlice{
	    fmt.Printf("\nindex: %v value: %v", i, v)
	}

	for i:=0; i<10; i++{
		fmt.Printf("\nindex: %v value: %v", i, i*i)
	}


}

func myFunction() {
	var printValue string = "Learning go functions"
	printMe(printValue)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(a, b int) (int , int, error){
	var err error
	if b==0{
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
    var result int = a/b
	var remainder int = a%b
	return result, remainder, err
}