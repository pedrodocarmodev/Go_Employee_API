package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type employee struct {
	ID		string	`json:"id"`
	Name	string	`json:"name"`
	Email	string	`json:"email"`
	Salary	int		`json:"salary"`
	Active	bool	`json:"active"` 
}

var employees = []employee{
	{ID: "1", Name: "Jonas", Email: "jonas@example.com", Salary: 1600, Active: true},
	{ID: "2", Name: "Erik", Email: "erik@example.com", Salary: 3000, Active: true},
	{ID: "3", Name: "Amélia", Email: "amelia@example.com", Salary: 8000, Active: true},
	{ID: "4", Name: "Cleiton", Email: "cleiton@example.com", Salary: 0, Active: false},
	{ID: "5", Name: "Osmar", Email: "osmar@example.com", Salary: 0, Active: false},
}

func employeeById(c *gin.Context) {
	id := c.Param("id")
	employee, err := getEmployeeById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "employee not found"})
		return 
	}

	c.IndentedJSON(http.StatusOK, employee)
}

func getEmployeeById(id string) (*employee, error) {

	for i, b := range employees {
		if b.ID == id {
			return &employees[i], nil
		}
	}

	return nil, errors.New("employee not found")	
}

func fire(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Id parameter not found"})
		return
	}

	employee, err := getEmployeeById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "employee not found"})
		return
	}

	if !employee.Active {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "this employee is already inactive"})
		return
	}

	employee.Active = false
	c.IndentedJSON(http.StatusOK, employee)
}

func employ(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Id parameter not found"})
		return
	}

	employee, err := getEmployeeById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "user not found"})
		return
	}

	if employee.Active {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "this employee is already active"})
		return
	}

	employee.Active = true
	c.IndentedJSON(http.StatusOK, employee)
}

func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

func addEmployee(c *gin.Context) {
	var newEmployee employee

	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "could not create the employee"})
		return
	}

	employees = append(employees, newEmployee)
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

func main() {
	router := gin.Default()

	router.GET("/employees", getEmployees)
	router.GET("/employees/:id", employeeById)
	router.POST("/employees", addEmployee)
	
	router.PATCH("/fire", fire)
	router.PATCH("/employ", employ)

	router.Run("localhost:8080")
}