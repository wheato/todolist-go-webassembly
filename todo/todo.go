package main

import (
	"fmt"
	"syscall/js"
	"strconv"
	"bytes"
	"html/template"
)

type TodoList struct{
	Items []*TodoItem
}

type TodoItem struct{
	Compeleted bool
	Text string
	ID int32
}

const templ = `
{{range .Items}}
	{{ if .Compeleted }}
		<li class='compeleted' data-index='{{.ID}}'>{{.Text}}</li>
	{{ else }}
		<li data-index='{{.ID}}'>{{.Text}}</li>
	{{ end }}
{{end}}
`
var (
	tl TodoList
)

func main() {
	c := make(chan struct{}, 0)
	addBtn := js.Global().Get("document").Call("querySelector", "#js_addTodo")
	listDom := js.Global().Get("document").Call("querySelector", "#j_todoList")

	addHandlerCb := js.NewCallback(jsAddHandler)
	toggleHandlerCb := js.NewCallback(jsToggleHandler)
	addBtn.Call("addEventListener", "click", addHandlerCb)
	listDom.Call("addEventListener", "click", toggleHandlerCb)
	<- c
}

func jsAddHandler(args []js.Value) {
	inputDom := js.Global().Get("document").Call("querySelector", "#j_textInput")
	listDom := js.Global().Get("document").Call("querySelector", "#j_todoList")
	text := inputDom.Get("value").String()
	if len(text) > 0 {
		tl.AddTodoItem(text)
		inputDom.Set("value", "")
	}
	html := tl.Render()
	listDom.Set("innerHTML", html)
}

func jsToggleHandler(args []js.Value) {
	el := args[0].Get("target")
	elNodeName := el.Get("nodeName").String()
	if elNodeName == "LI" {
		listDom := js.Global().Get("document").Call("querySelector", "#j_todoList")
		id, _ := strconv.ParseInt(el.Get("dataset").Get("index").String(), 10, 32)
		tl.ToggleStatus(int32(id))
		html := tl.Render()
		listDom.Set("innerHTML", html)
	}
}

// NewTodo will Create A Todo
func NewTodo(text string, id int32) TodoItem {
	return TodoItem{ Compeleted: false, Text: text, ID: id }
}

// AddTodoItem could add a todoItem to TodoList
func (tl *TodoList) AddTodoItem(text string) TodoItem{
	newItem := NewTodo(text, int32(len(tl.Items)))
	tl.Items = append(tl.Items, &newItem)
	return newItem
}

// GetByID could find a todoItem by Id
func (tl TodoList) GetByID(id int32) *TodoItem {
	for _, item := range tl.Items{
		if item.ID == id {
			return item
		}
	}
	return &TodoItem{}
}

// ToggleStatus switch the todoItem compeleted
func (tl TodoList) ToggleStatus(id int32) {
	target := tl.GetByID(id)
	target.Compeleted = !target.Compeleted
	fmt.Printf("%v - %s\n", target.Compeleted, target.Text)
}

// Render compile template to string
func (tl *TodoList) Render() string{
	var html bytes.Buffer
	todoList, _ := template.New("TodoList").Parse(templ)
	todoList.Execute(&html, tl)
	return html.String()
}

