package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	EmployeeId   string `json:"employeeid"`
	EmployeeName string `json:"employeename"`
	Company      string `json:"company"`
	Salary       int    `json:"salary"`
}

var employees = []Employee{
	{
		EmployeeId: "001", EmployeeName: "Jay", Company: "Xyz", Salary: 2400,
	},
	{
		EmployeeId: "002", EmployeeName: "Krishna", Company: "zey", Salary: 2800,
	},
	{
		EmployeeId: "003", EmployeeName: "Reddy", Company: "zwd", Salary: 3400,
	},
}

func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

func main() {
	router := gin.Default()
	router.GET("/employees", getEmployees)
	router.Run("localhost:8080")
}
