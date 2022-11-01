package models

type Employee struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Salary float32 `json: "salary"`
	Age    int     `json : "age"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}
