/*
Example explained

Line 1: In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.

Line 2: import ("fmt") lets us import files included in the fmt package.

Line 3: A blank line. Go ignores white space. Having white spaces in code makes it more readable.

Line 4: func main() {} is a function. Any code inside its curly brackets {} will be executed.

Line 5: fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".
*/
package main

import "fmt"

func main() {
	fmt.Print("Hello");

	// Go Variable Types
	
	// will guess the type
	// Note: It is not possible to declare a variable using ":=" without assigning a value to it.
	// var is global
	// := is local
	age := 0
	var fly = false

	// defined type
	var age2 int = 2
	var money float32 = 2.2
	var name string = "Raziel"
	var adult bool = true

	// declared without value
	var experience int

	// assingig after delcaration
	experience = 10

	// you can also declare multiple variables in multiple lines
	var rice, beans = true, false

	// different types
	var salad, drink = "Caesar", 10

	// inside blocks
	var (
		meal int = 10
		xp float32 = 10.4
	)

	fmt.Println(age, fly, age2, money, name, adult, experience, rice, beans, salad, drink, meal, xp)

/* 	Go variable naming rules:

	A variable name must start with a letter or an underscore character (_)
	A variable name cannot start with a digit
	A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
	Variable names are case-sensitive (age, Age and AGE are three different variables)
	There is no limit on the length of the variable name
	A variable name cannot contain spaces
	The variable name cannot be any Go keywords */


	// CONSTANTS

	// declared
	const PI float32 = 3.14

	// will guess
	const CEP = 111

	const (
		A string = "A"
		B int = 2
	)

	fmt.Print(PI)
	fmt.Printf("PI %v and %T" , PI, PI )

	fmt.Print(A, ' ', B)

}