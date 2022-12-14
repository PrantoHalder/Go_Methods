package main

import "fmt"

type person struct {
	Name   string
	age    int
	salary int
}

func method1() {
	var per1 person
	per1.Name = "pranto"
	per1.salary = 2000
	per1.age = 25
	per1.printPersonSalary()
}

func (p person) printPersonSalary() {
	fmt.Printf("the salary of %s is %d\n",p.Name,p.salary)
}
