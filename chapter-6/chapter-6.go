package main

import (
	"fmt"
)
/**
1、
if condition {  
}
2、
if condition {  
} else if condition {
} else {
}
*/
func main()  {
	fmt.Println("if-else")
	f1()
	f2()
	f3()
	f4()
}

func f1()  {
	num := 10
    if num % 2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    }  else {
        fmt.Println("the number is odd")
    }
}
func f2()  {
    if num := 10; num % 2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    }  else {
        fmt.Println("the number is odd")
    }
}
func f3()  {
    num := 99
    if num <= 50 {
        fmt.Println("number is less than or equal to 50")
    } else if num >= 51 && num <= 100 {
        fmt.Println("number is between 51 and 100")
    } else {
        fmt.Println("number is greater than 100")
    }
}
//由于 if{…} else {…} 是一个单独的语句，它的中间不应该出现分号。因此，需要将 else 语句放置在 } 之后处于同一行中。
func f4()  {
	num := 10
    if num % 2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    }  else {
    // else {
        fmt.Println("the number is odd")
    }
}
