package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

// func (e *Engineer) GetName() string {
// 	return "Employee Name : " + e.Name
// }

func GetName(e *Engineer) string {
	return "Employee Name : " + e.Name
}

func PrintDetails(e Employee) {
	fmt.Print(e.GetName())
}

func main() {
	engineer := &Engineer{Name: "Gaurav Shinde"}

	PrintDetails(engineer)
}
