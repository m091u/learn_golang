package main

import (
	"fmt"
	"math/rand"
)

func double (x int) int {
	return x * 2
}

func add (lhs, rhs int)int{
	return lhs+rhs
}

func greet(){
	fmt.Println("hello")
}

func greetPerson(name string) string{
	return fmt.Sprintln("hello", name)
}

func anyMessage() string{
	return "any message"
}

func sumOfThrees(a, b,c int) int{
	return a+b+c
}

func anyNumber() int{
	return 5
}

func anyTwo() (int, int){
	return 2, 2
}
func addThree(a, b, c int)(int){
	return a+b+c
}

func average(a, b, c int) float32 {
	return float32(a + b+ c) /3
}

func roll(sides int) int{
	return rand.Intn(sides) + 1
}

func main(){

	//dice roller
	dice, sides := 2, 6
	rolls := 2
	
	for r := 1; r <= rolls; r++ {
	    sum := 0

		for d :=1; d <= dice; d++{
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll #",r, ",Dice #", d, ":", rolled)
		}

		fmt.Println("Total rolled: ", sum)
		switch sum := sum;{
			case sum ==2 && dice == 2:
				fmt.Println("Snake eyes!")
			case sum ==7:
				fmt.Println("Lucky 7!")
			case sum % 2 == 0:
				fmt.Println("Even!")
			case sum % 2 == 1:
				fmt.Println("Odd!")
		}
	}


	// loops
	sum := 0
	fmt.Println(sum)

	for i := 0; i<= 11; i++{
		sum +=1
		fmt.Println(sum)
	}
	for sum >= 10 {
		sum -= 5
		fmt.Println("Decremented sum is",sum)
	}

	for i:= 1; i <=10 ; i++{
		if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0{
			fmt.Println("Buzz")
		} else if i % 3 == 0 && i % 5 == 0{
		    fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
	

    //switch
	switch age := 14;{
	case age == 0:
		fmt.Println("newborn")
	case age >= 1 && age <=3:
		fmt.Println("toddler")
	case age > 4 && age <=12:
		fmt.Println("child")
	case age >=13 && age <= 17:
		fmt.Println("teenager")
	case age > 18:
		fmt.Println("adult")
	default:
		fmt.Println("unknown")
	}

	switch p := price();{
	case p < 2:
		fmt.Println("cheap item")
	case p < 10:
		fmt.Println("moderately priced")
	default:
		fmt.Println("expensive item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("First Class seating")
	default:
		fmt.Println("Unknown ticket type")
	}

    //if-else
	today, role := Monday, Guest

	if role == Manager || role == Admin {
		accessGranted()
	} else if role == Contractor && !weekday(today){
		accessGranted()
	} else if role == Member && weekday(today){
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
	    accessGranted()
	} else {
		accessDenied()
	}

	accessGranted()

	//if-else
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1> quiz2 {
		fmt.Println("quiz1 is greater")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz2 is greater")
	} else{
		fmt.Println("quiz1 and quiz2 are equal")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}


	greet()
	fmt.Println(greetPerson("john"))
	fmt.Println(anyMessage())
	sumOfThree := sumOfThrees(3,4,5)
	fmt.Println(sumOfThree)
	anyNumber()
	anyTwo()
	a , b := anyTwo()
	answer := addThree(a, b, anyNumber())
	fmt.Println(answer)

	dozen := double(6)
	fmt.Println("the double of 6 is",dozen)

	bakerDozen := add (dozen, 1)
	fmt.Println("the baker dozen is", bakerDozen)

	anotherBakersDozen := add (double(6),1)
	fmt.Println("the another bakers dozen is", anotherBakersDozen)


	favColor := "green"
	fmt.Println("my fav color is",favColor) //green
	birthYear, ageInYears  := 1985, 39
	fmt.Println("Birth Year:", birthYear, "age", ageInYears) 

	var (
	   firstName = `M`
	   lastname = `T`
	)

	fmt.Println("my initials are", firstName, lastname)

	var ageinDays int
	ageinDays = ageInYears *365

	fmt.Println("I am age in days", ageinDays) //7005
}