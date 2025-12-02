package main

import "fmt"

func main() {
	// 1. Declare variable with explicit type
	var x int
	var name string
	fmt.Println("x:", x)       // 0 (zero value)
	fmt.Println("name:", name) // "" (empty string)

	// 2. Declare and initialize
	var age int = 25
	var msg string = "Hello"
	fmt.Println("age:", age)
	fmt.Println("msg:", msg)

	// 3. Let Go infer the type
	var pi = 3.14
	var ok = true
	fmt.Println("pi:", pi)
	fmt.Println("ok:", ok)

	// 4. Short variable declaration (:=)
	score := 100
	valid := false
	fmt.Println("score:", score)
	fmt.Println("valid:", valid)

	// 5. Declare multiple variables at once
	var a, b, c int = 1, 2, 3
	fmt.Println("a, b, c:", a, b, c)

	name2, age2 := "Alice", 30
	fmt.Println("name2, age2:", name2, age2)

	// 6. Multi-variable block declaration
	var (
		city    string = "New York"
		country string = "USA"
		active  bool
	)
	fmt.Println("city:", city, "country:", country, "active:", active)

	// 7. Constants
	const PI = 3.14159
	fmt.Println("Constant PI:", PI)

	// 8. Blank identifier to ignore values
	_, err := fmt.Println("This prints, error ignored")
	fmt.Println("Returned error ignored:", err) // always nil here
}
