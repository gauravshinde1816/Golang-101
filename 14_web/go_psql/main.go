package main

import (
	"fmt"
	"go_psql/configs"
	"go_psql/models"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	// connect to db
	db := configs.ConnectDB()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// routes

	// create employee
	e.POST("/employees", func(c echo.Context) error {
		u := new(models.Employee)
		if err := c.Bind(u); err != nil {
			return err
		}
		// generate id
		charset := "abcdefghijklmnopqrstuvwxyz"

		num := rand.Intn(len(charset))

		id := charset[num]

		sqlStatement := "INSERT INTO employees (id,name,salary,age)VALUES ($1,$2,$3,$4)"
		res, err := db.Query(sqlStatement, id, u.Name, u.Salary, u.Age)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.JSON(http.StatusOK, res)
	})

	// Get all employees

	e.GET("/employees", func(c echo.Context) error {
		sqlStatement := "SELECT id, name, salary, age FROM employees"
		rows, err := db.Query(sqlStatement)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		result := models.Employees{}

		for rows.Next() {
			employee := models.Employee{}
			err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
			// Exit if we get an error
			if err2 != nil {
				return err2
			}
			result.Employees = append(result.Employees, employee)
		}
		return c.JSON(http.StatusCreated, result)
	})

	e.GET("/employees/:empid", func(c echo.Context) error {
		empid := c.Param("empid")
		sqlStatement := "SELECT id, name, salary, age FROM employees WHERE id=$1"

		res, err := db.Query(sqlStatement, empid)

		if err != nil {
			fmt.Println(err)
		}

		return c.JSON(http.StatusOK, res)

	})

	e.Logger.Fatal(e.Start(":1323"))

}
