// main.go
package main

import (
	"fmt"
)


func main(){
	var todos Todos
	fmt.Println(todos)
	todos.Add("Create homework")
	todos.Add("wash groceries")
	fmt.Printf("list of todos : %v\n\n", todos)
	todos.Toggle(0)
	fmt.Printf("mark todo as done : %v\n", todos)
	todos.Delete(1)
	fmt.Printf("delet to %v\n", todos)
}