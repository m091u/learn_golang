package main

import (
	"fmt"
	"strings"
)

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

//method pointing to teh Car object
func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Driving(){
	fmt.Println("Driving the car")
}

func (c Car) getName() string{
	return c.Name
}

//interfaces
type Vehicle interface{
	Stop()
	Drive()
}
type Truck struct{
    TruckModel string
}
type Dacia struct{
	   DaciaModel string
}

func (t Truck) Drive(){
	fmt.Println("Driving the truck")
	fmt.Println(t.TruckModel)
}

func (d Dacia) Stop(){
	fmt.Println("Stopping the dacia")
	fmt.Println(d.DaciaModel)
}


func main() {

	// control structures: if-else, for

	arr1 := []string{"mirabela", "marinela","tiugan"}
	
	for i:= range arr1{
		fmt.Println(arr1[i])
	}

	newMap:= make(map[string]interface{})
	fmt.Println("New map is",newMap)
	newMap["key1"] = "value1"
	newMap["key2"] = 2
	newMap["key3"] = 3.14

	for k, v := range newMap{
		fmt.Println(k, v)
	}

	//interfaces defining a structure/ guidelines that methods should follow
	t := Truck{"Truck1"}
	t.Drive()
	d := Dacia{"Dacia1"}
	d.Stop()

	//used with maps as a blank value

	 mapExample := make(map[string]interface{}) //key is a tring but the value can be anything
	 mapExample["key1"] = "value1"
	 mapExample["key2"] = 123
	 mapExample["key3"] = true
	 fmt.Println(mapExample)

	//structures abstract data type that groups together logically related data
	// c := Car{"dacia", 2, 123}
	c := Car{
		Name:    "dacia",
		Age:     2,
		ModelNo: 123,
	}
	c.Print()
	c.Driving()
	fmt.Println(c.getName())

	//string
	phrase := "I am home and not out"
	stringInPhrase := "home"

	fmt.Println(strings.Contains(phrase, stringInPhrase))

	fmt.Println(strings.Index(phrase, stringInPhrase))
	fmt.Println((strings.Replace(phrase, "am", "am not", 1)))
	fmt.Println((strings.Replace(phrase, "o", "O", -1)))
	fmt.Println((strings.Split(stringInPhrase, ",")))
	fmt.Println((strings.ToLower(phrase)))
	fmt.Println((strings.Title(phrase)))

	str := "hello"
	chars := strings.Split(str, "")

	result := strings.Join(chars, ",")
	fmt.Println(result)

	//arrays
	// var arr []int
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr = append(arr, 6, 7, 8)
	fmt.Println(arr)

	//pointer
	m1 := 2
	prt := &m1 // & is the referencing operator
	fmt.Println(prt)
	fmt.Println(*prt) // * used to get the value

}
