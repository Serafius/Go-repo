package main

import "fmt"

func main() {
	var num int = 5
	var myStr string = "Word"
	// bool, float, rune, byte, ...
	fmt.Println(num, myStr)
	age := 19
	name := "Ahmad"

	fmt.Print("Hello, ")  // Doesn't ends a line unless adding '\n'
	fmt.Println("Ahmad.") // ends the line.

	str1 := "This sentence"
	str2 := "doesn't have whitespaces added!"
	fmt.Print(str1, str2, "\n")

	str3 := "This sentence"
	str4 := "has whitespaces added!"
	fmt.Println(str3, str4)

	fmt.Printf("my age is %v and my name is %v \n", age, name) // %v add the value in default format
	fmt.Printf("my age is %q and my name is %q \n", age, name) // %q adds quotes to the variable // int numbers gets printed as hexidecimal by default
	fmt.Printf("age is of type %T \n", age)                    // %T add the type  age is an int.
	fmt.Printf("you have %f \n", 225.55)                       // %f = float format 6 numbers after point by default
	fmt.Printf("you have %0.1f\n", 225.55)                     //Specify how many decimal numbers to print.

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name) // return a formatted string and saves it
	fmt.Sprintln("New string is", str)
}
