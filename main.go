package main

import (
	"fmt"
)

func main() {
	fmt.Println("Inside the go cli")
	todos := Todos{}

	storage := NewStorage[Todos]("todos.json")

	storage.Load(&todos)

	todos.delete(3)

	todos.print()

	storage.Save(todos)

}
