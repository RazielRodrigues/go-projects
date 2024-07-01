/*
Example explained

Line 1: In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.

Line 2: import ("fmt") lets us import files included in the fmt package.

Line 3: A blank line. Go ignores white space. Having white spaces in code makes it more readable.

Line 4: func main() {} is a function. Any code inside its curly brackets {} will be executed.

Line 5: fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".
*/

// tutorial by: https://www.w3schools.com/go/go_maps.php

package main

import "fmt"

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

/* Multiple Return Values
 */

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func main() {
	fmt.Print("Hello")

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
		meal int     = 10
		xp   float32 = 10.4
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
		B int    = 2
	)

	fmt.Print(PI)
	/* 	Verb	Description
	%v	Prints the value in the default format
	%#v	Prints the value in Go-syntax format
	%T	Prints the type of the value
	%%	Prints the % sign
	*/
	fmt.Printf("PI %v and %T", PI, PI)

	fmt.Print(A, ' ', B)

	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	/*
		%b	Base 2
		%d	Base 10
		%+d	Base 10 and always show sign
		%o	Base 8
		%O	Base 8, with leading 0o
		%x	Base 16, lowercase
		%X	Base 16, uppercase
		%#x	Base 16, with leading 0x
		%4d	Pad with spaces (width 4, right justified)
		%-4d	Pad with spaces (width 4, left justified)
		%04d	Pad with zeroes (width 4 */
	var i2 = 15

	fmt.Printf("%b\n", i2)
	fmt.Printf("%d\n", i2)
	fmt.Printf("%+d\n", i2)
	fmt.Printf("%o\n", i2)
	fmt.Printf("%O\n", i2)
	fmt.Printf("%x\n", i2)
	fmt.Printf("%X\n", i2)
	fmt.Printf("%#x\n", i2)
	fmt.Printf("%4d\n", i2)
	fmt.Printf("%-4d\n", i2)
	fmt.Printf("%04d\n", i2)

	/*
	     %s	Prints the value as plain string
	   %q	Prints the value as a double-quoted string
	   %8s	Prints the value as plain string (width 8, right justified)
	   %-8s	Prints the value as plain string (width 8, left justified)
	   %x	Prints the value as hex dump of byte values
	   % x	Prints the value as hex dump with spaces */
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	/* Verb	Description
	   %e	Scientific notation with 'e' as exponent
	   %f	Decimal point, no exponent
	   %.2f	Default width, precision 2
	   %6.2f	Width 6, precision 2
	   %g	Exponent as needed, only necessary digits */
	fmt.Printf("%e\n", i)
	fmt.Printf("%f\n", i)
	fmt.Printf("%.2f\n", i)
	fmt.Printf("%6.2f\n", i)
	fmt.Printf("%g\n", i)

	/*   Unsigned Integers
	Unsigned integers, declared with one of the uint keywords, can only store non-negative values:
	*/
	/* 	var uit uint = 10
	 */
	/* 	Go has five keywords/types of signed integers: */

	/* Type	Size	Range
	   int	Depends on platform:
	   32 bits in 32 bit systems and
	   64 bit in 64 bit systems	-2147483648 to 2147483647 in 32 bit systems and
	   -9223372036854775808 to 9223372036854775807 in 64 bit systems
	   int8	8 bits/1 byte	-128 to 127
	   int16	16 bits/2 byte	-32768 to 32767
	   int32	32 bits/4 byte	-2147483648 to 2147483647
	   int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775 */

	var x float32 = 123.78
	var y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)

	var arrayDefinido = [5]int{1, 2, 3, 4, 5}

	var arrayIndefinido = [...]string{"Raziel", "Jaiara"}
	var arrayMapeado = [3]int{1: 10}

	outraForm := [1]int{1}
	outraform2 := [...]int{12, 4, 332}

	fmt.Println(arrayDefinido, arrayIndefinido, arrayMapeado, outraForm, outraform2)
	fmt.Println(arrayDefinido[4])
	arrayDefinido[4] = 12
	fmt.Println(arrayDefinido[4])
	fmt.Println(len(arrayDefinido))

	pizzaSlice := []int{1, 2}

	fmt.Println(pizzaSlice)
	fmt.Println(len(pizzaSlice))
	fmt.Println(cap(pizzaSlice))

	var cake = [4]int{1, 2, 4, 5}

	sliceOfCake := cake[0:2]

	fmt.Println(cake)
	fmt.Println(len(cake))
	fmt.Println(cap(cake))

	fmt.Println(sliceOfCake)

	makeSlice := make([]int, 10)

	fmt.Println(makeSlice)

	makeSlice2 := make([]int, 5)

	fmt.Println(makeSlice2)

	makeSlice[2] = 1

	makeSlice = append(makeSlice, 10, 2)
	fmt.Println(makeSlice)

	// Note: The '...' after slice2 is necessary when appending the elements of one slice to another.
	makeSlice2 = append(makeSlice, makeSlice2...) // Change length by appending items

	pizzaSlice = pizzaSlice[1:2] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", pizzaSlice)
	fmt.Printf("length = %d\n", len(pizzaSlice))
	fmt.Printf("capacity = %d\n", cap(pizzaSlice))

	/* 	Memory Efficiency
	When using slices, Go loads all the underlying elements into the memory.

	If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.

	The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.

	*/

	var numbers = [5]int{1, 2, 3, 4, 5}

	neededNumbers := numbers[:len(numbers)-2]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	/* Arithmetic Operators */
	/* Arithmetic operators are used to perform common mathematica l operations.*/
	/*  */
	/* Operator	Name	Description	Example	Try it */
	/* +	Addition	Adds together two values	x + y	 */
	/* -	Subtraction	Subtracts one value from another	x - y	 */
	/* *	Multiplication	Multiplies two values	x * y	 */
	/* /	Division	Divides one value by another	x / y	 */
	/* %	Modulus	Returns the division remainder	x % y	 */
	/* ++	Increment	Increases the value of a variable by 1	x++ */
	/* --	Decrement	Decreases the value of a variable by 1	x-- */

	/* A list of all assignment operators:

	   Operator	Example	Same As	Try it
	   =	x = 5	x = 5
	   +=	x += 3	x = x + 3
	   -=	x -= 3	x = x - 3
	   *=	x *= 3	x = x * 3
	   /=	x /= 3	x = x / 3
	   %=	x %= 3	x = x % 3
	   &=	x &= 3	x = x & 3
	   |=	x |= 3	x = x | 3
	   ^=	x ^= 3	x = x ^ 3
	   >>=	x >>= 3	x = x >> 3
	   <<=	x <<= 3	x = x << 3
	*/

	/*  A list of all comparison operators:

	Operator	Name	Example	Try it
	==	Equal to	x == y
	!=	Not equal	x != y
	>	Greater than	x > y
	<	Less than	x < y
	>=	Greater than or equal to	x >= y
	<=	Less than or equal to	x <= y */

	/* Logical Operators
	   Logical operators are used to determine the logic between variables or values:

	   Operator	Name	Description	Example	Try it
	   && 	Logical and	Returns true if both statements are true	x < 5 &&  x < 10
	   || 	Logical or	Returns true if one of the statements is true	x < 5 || x < 4
	   !	Logical not	Reverse the result, returns false if the result is true	!(x < 5 && x < 10) */
	/*
	   Bitwise Operators
	   Bitwise operators are used on (binary) numbers:

	   Operator	Name	Description	Example	Try it
	   & 	AND	Sets each bit to 1 if both bits are 1	x & y
	   |	OR	Sets each bit to 1 if one of two bits is 1	x | y
	    ^	XOR	Sets each bit to 1 if only one of two bits is 1	x ^ b
	   <<	Zero fill left shift	Shift left by pushing zeros in from the right	x << 2
	   >>	Signed right shift	Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off */

	if 20 > 30 {
		fmt.Println("ok!")
	}

	time := 20
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}

	a := 3

	switch a {
	case 1:
		fmt.Println("a is one")
		/*   case "b":
		     fmt.Println("a is b") */
	}

	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Println(i)

	}

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	/*   Tip: To only show the value or the index, you can omit the other output using an underscore (_).
	 */
	/*   fruits := [3]string{"apple", "orange", "banana"}
	     for _, val := range fruits {
	        fmt.Printf("%v\n", val)
	     } */

	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)

	fmt.Println(myFunction(5, "Hello"))

	/* If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value.
	 */
	_, b := myFunction(5, "Hello")
	fmt.Println(b)

	testcount(1)

	type Computer struct {
		id   int
		name string
	}

	var computer Computer

	computer.id = 1
	computer.name = "Windows"

	fmt.Println(computer.name)
	fmt.Println(computer.id)

	/* 	Note: The order of the map elements defined in the code is different from the way that they are stored. The data are stored in a way to have efficient data retrieval from the map.
	 */

	var mapa = map[int]string{1: "Raziel", 2: "Jaiara"}
	fmt.Println(mapa)

	var carro = make(map[string]string) // The map is empty now
	carro["brand"] = "Ford"
	carro["model"] = "Mustang"
	carro["year"] = "1964"

	fmt.Println(carro)
	fmt.Printf(carro["brand"])
	carro["color"] = "red" // Adding an element
	carro["year"] = "1970" // Updating an element
	delete(carro, "year")

	val1, ok1 := carro["brand"] // Checking for existing key and its value
	val2, ok2 := carro["color"] // Checking for non-existing key and its value
	val3, ok3 := carro["day"]   // Checking for existing key and its value
	_, ok4 := carro["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
	// iterating
	for k, v := range carro {
		fmt.Printf("%v : %v, ", k, v)
	}

}
