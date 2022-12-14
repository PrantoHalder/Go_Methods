package main

import "fmt"

type Person struct{
	Name string
	age int
	Job
}
type Job struct{
	Title string
	salary int
}

func method2 () {
var per2 Person
var per3 Person

per2.Name = "pranto"
per2.age = 25
per2.salary = 1200
per2.Title = "Junior Softwate Engineer"

per3.age = 26
per3.salary = 1500
per3.Name = "Shovon"
per3.Title = "Software Engineer"

per2.printPersonJobTitle()
per3.printPersonJobTitle()

}

func (p Person) printPersonJobTitle(){
	fmt.Printf("%s's job title is %s\n",p.Name,p.Title)
}