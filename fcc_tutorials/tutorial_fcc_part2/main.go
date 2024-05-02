package main

import (
	"fmt"
	// "log"
	// "io"
	"math"
	// "net/http"
	"rsc.io/quote"
)

func main() {
    myMap := map[string]int{
        "apple": 1,
        "banana": 2,
        "cherry": 3,
    }

	if pop , ok := myMap["apple"]; ok {
		fmt.Println("Value:", pop)
	}

	//guessing game
	number := 50
	guess := 70

	if guess < number {
		fmt.Printf("Guess less than number\n")
	} else if guess > number {
	    fmt.Printf("Guess more than number\n")
	} else if guess == number {
	    fmt.Printf("You guessed correctly\n")
	}

	myNum := 0.1
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Printf("Equal\n")   
	} else {
		fmt.Printf("Not Equal\n")
	}

	//switch
	switch 5 {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		default:
			fmt.Println("none")
		}

	var i interface{} = "a" //interface can take any value
	switch i.(type) {
	    case int:
	        fmt.Println("i is an int")
	    case float64:
	        fmt.Println("i is a float64")
	    case string:
	        fmt.Println("i is a string")
	    default:
	        fmt.Println("i is unknown type")
	}

	//looping
	for i := 0; i < 5; i++ {
	    fmt.Println(i)
	}

	//special case replacing while loops
	 j := 0
	 for ; j <5 ; {
		fmt.Printf("Print j values: %v\n ", j)
		j++
	 }

	 // loops with an undefined end statement
	 k := 0
	 for {
		fmt.Println(k)
		k++
		if k == 5 {
		    break
		}
	 }

	 //for range loops for slices
	 s := []int{1,2,3}
	 for k, v := range s {      // key and value in the collection
        fmt.Println(k,v)    // keys will be the indexes
	 }  

	 //looping through maps
	 m := map[string]int{
	     "apple": 1,
	     "banana": 2,
	     "cherry": 3,
	 }
	 for k, v := range m {
	     fmt.Println(k, v)
	 }

	 //extracting just values
	 for _, v := range m {  //k is replaced with _, to extract only keys, remove v
	     fmt.Println(v)
	 }

	 //defer
	//  res, err := http.Get("http://www.google.com/robots.txt")
	//  if err != nil{
    //     log.Fatal(err)   
	//  }
	//  defer res.Body.Close()  //deferal can be aded after the resource opening to make sure you do not forget
	//  robots, err := io.ReadAll(res.Body)
	 
	//  if err != nil {
	// 	log.Fatal(err)
	//  }
	//  fmt.Printf("%s", robots)

	//  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//      w.Write([]byte("Hello Go!"))
	//  })
	//  err := http.ListenAndServe(":8080", nil)
	//  if err != nil {
	//      panic(err.Error())
	//  }

	 // functions
	 greeting := "Hello"
	 name := "Go"
	 sayHello(greeting, name)
	 fmt.Println(name)

	 sum := sum(1, 2, 3, 4 ,5)
	 fmt.Println("The sum is ",sum)

	 d, err := divide(10.0, 0.0)
	 if err != nil {
	     fmt.Println(err)
		 return
	 }
	 fmt.Println("Divide result is", d)

	 fmt.Println(quote.Go())

}

func sayHello (greeting, name string){
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}

//variatic parameter, should be the last in the parameters
func sum(value ...int) int {  //take in all of the last arguments that are passed in, and wrap them up into a slice with name value
	fmt.Println(value)
	total := 0
	for _, v := range value {
		total += v
	}
 return total
}

func divide (a, b float64) (float64, error) {
	if b == 0.0 {
	    return 0.0, fmt.Errorf("Cannot divide by zero")
	}
return a/b, nil
}