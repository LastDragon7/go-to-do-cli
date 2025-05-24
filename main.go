package main

import "fmt"

func main() {
	fmt.Println("Inside the go cli")
	todos := Todos{}

	todos.add("Buy Milk")
	todos.add("Buy Toast")

	todos.toggle(0)

	todos.print()

	todos.delete(0)

	todos.toggle(0)

	todos.print()

}
