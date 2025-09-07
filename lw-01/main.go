package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Greet() {
	fmt.Printf("Привет, меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

func (h Human) Walk() {
	fmt.Println(h.Name, "идет.")
}

type Action struct {
	Human
	ActionDescription string
}

func (a Action) PerformAction() {
	fmt.Printf("%s выполняет действие: %s\n", a.Human.Name, a.ActionDescription)
}

func main() {
	person := Human{Name: "Алиса", Age: 30}
	person.Greet()
	person.Walk()

	fmt.Println("--------------------")

	task := Action{
		Human:             Human{Name: "Боб", Age: 25},
		ActionDescription: "кодирование на Go",
	}

	task.Greet()
	task.Walk()

	task.PerformAction()

	fmt.Println("--------------------")

	anotherTask := Action{
		Human:             person,
		ActionDescription: "чтение книги",
	}
	anotherTask.Greet()
	anotherTask.PerformAction()
}
