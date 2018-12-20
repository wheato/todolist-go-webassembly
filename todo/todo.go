package main

import (
	"fmt"
	"syscall/js"
)

type TodoList struct{
	Items []TodoItem
}

type TodoItem struct{
	Compeleted bool
	Text string
	ID int
}

var (
	tl TodoList
)

func main() {
	c := make(chan struct{}, 0)
	addBtn := js.Global().Get("document").Call("querySelector", "#js_addTodo")
	addHandlerCb := js.NewCallback(jsAddHandler)
	addBtn.Call("addEventListener", "click", addHandlerCb)
	<- c
}

func jsAddHandler(args []js.Value) {
	inputDom := js.Global().Get("document").Call("querySelector", "#j_textInput")
	text := inputDom.Get("value").String()
	if len(text) > 0 {
		tl.AddTodoItem(text)
	}
	fmt.Printf("%d\n", len(tl.Items))
}

func NewTodo(text string, id int) TodoItem {
	return TodoItem{ Compeleted: false, Text: text, ID: id }
}

func SetId(id int, item TodoItem) {
	item.ID = id
}

func (tl *TodoList) AddTodoItem(text string) TodoItem{
	newItem := NewTodo(text, len(tl.Items))
	tl.Items = append(tl.Items, newItem)
	return newItem
}

func (tl TodoList) GetById(id int) TodoItem {
	for _, item := range tl.Items{
		if item.ID == id {
			return item
		}
	}
	return TodoItem{}
}

func (tl TodoList) ToggleStatus(id int) {
	target := tl.GetById(id)
	if target.Text != "" {
		target.Compeleted = !target.Compeleted
	}
}

func (tl TodoList) Render() string{
	html := ""
	for i, item := tl.Items {
		html += item.Text
	}
	return html
}

