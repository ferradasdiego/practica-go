package main

import (
	"fmt"
	"strconv"
)

// No devuelve nada, pinta por consola el string y el int que le pasas
func sayHello(name string, years int) {
	var yearsToS string = strconv.Itoa(years)
	fmt.Println("Hello, " + name + " you have " + yearsToS + " years old")
}

// Devuelve un string pasando un string y un int
func sayHello2(name string, years int) string {
	return ("Your name is: " + name + " and you have " + strconv.Itoa(years) + " old")
}

// Devuelve un string y un int
func sayHello3() (string, int) {
	return ("hola"), (5)
}

func main() {
	sayHello("diego", 27)
	fmt.Println(sayHello2("diego", 27))
	fmt.Println(sayHello3())
}
