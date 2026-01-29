// todo.go
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string){
     todo := Todo{
        Title: title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
	 }
	 *todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos){
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) Delete(index int) error{
	 if err := todos.validateIndex(index); err != nil{
		return err
	 }
	 *todos = append((*todos)[:index], (*todos)[index + 1:]...)
	 return nil
}

func (todos *Todos) Toggle(index int) error{
	if err := todos.validateIndex(index); err != nil{
		return err
	}

	todo := &(*todos)[index]

	if !todo.Completed {
		now := time.Now()
		todo.CompletedAt = &now
	}else{
		todo.CompletedAt = nil
	}

	todo.Completed = !todo.Completed
    
	return nil
}

func (todos *Todos) Edit(index int, title string,) error{
    if err := todos.validateIndex(index); err != nil{
		return err
	}

	(*todos)[index].Title = title
    
	return nil
}


func (todos *Todos) Print(){
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title",  "Completed At", "Created At", "Completed",)
	for index, row := range *todos{
		completed := "❌"
		completedAt := ""
		if row.Completed {
			completed = "✅ 🤌"
			completedAt = row.CompletedAt.Format(time.RFC1123)
		}
    	table.AddRow(strconv.Itoa(index), row.Title, completedAt, row.CreatedAt.Format(time.RFC1123), completed,)
	}

	table.Render()
}