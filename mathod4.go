package main

import "fmt"

type Subject struct { //declaring a struck
	Sub   []string
	Marks []int
}

type Student struct { //declaring the struckt
	Name    string
	Roll    int
	Subject //composite value
}

func stu() {
	students := []Student{
		{
			Name: "Mahmud",
			Roll: 271,
			Subject: Subject{                            //declaring the value
				Sub:   []string{"Bangla", "English"},
				Marks: []int{99, 80},
			},
		},
		{
			Name: "Enamul",                      //declaring the value
			Roll: 260,
			Subject: Subject{
				Sub:   []string{"Bangla", "English"},
				Marks: []int{70, 78}},
		},
		{
			Name: "Pranto",                           //declaring the value
			Roll: 150,
			Subject: Subject{
				Sub:   []string{"Bangla", "English"},
				Marks: []int{67, 58}},
		},
	}
	for _, Student := range students {  
		Student.PrintStudent()             //calling the method 
	}
}

func (s Student) PrintStudent() {                        //decalring a method
	fmt.Printf("Name : %v\n", s.Name)
	fmt.Println("Roll : ", s.Roll)
	fmt.Println("--------------------------------")
	fmt.Println("Sub	Mark 	Grade	GPA")
	fmt.Println("--------------------------------")
	for i := 0; i < len(s.Sub); i++ {
		
		fmt.Printf("%v ", s.Sub[i])
		value := s.Marks[i]
		fmt.Printf("  %v", value)
		if value <= 100 && value >= 80 {              //Grade calculation
			GPA := 4.00
			fmt.Printf("	A+\n	%v",GPA)
		} else if value <= 79 && value >= 70 {
			GPA := 3.5
			fmt.Printf("	A\n		%v",GPA)
		} else if value <= 69 && value >= 60 {
			GPA := 3.00
			fmt.Printf("	A-\n")
		} else if value <= 59 && value >= 50 {
			fmt.Printf("	B\n")
		} else if value <= 49 && value >= 40 {
			fmt.Printf("	C\n")
		} else {
			fmt.Printf("    F\n")
		}
	}
	fmt.Println("-------------------------------")
	
}
