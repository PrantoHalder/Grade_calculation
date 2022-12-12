package main

import "fmt"


type Person struct {
	Name string
	Roll int
	subject
}
type subject struct {
	sub []string
	marks []int
}

func grade (){
	
	person := []Person{
		{
			Name: "Pranto",
			Roll : 1,
			subject: subject{sub : []string{"CSE","EEE","ICT"},
			                 marks: []int{60,80,50},},
		},
		{
			Name: "Halder",
			Roll : 2,
			subject: subject{sub : []string{"CSE","EEE","ICT"},
			                 marks: []int{90,60,40},},
		},
		{
			Name: "Shovon",
			Roll : 3,
			subject: subject{sub : []string{"CSE","EEE","ICT"},
			                 marks: []int{70,90,50},},
		},}

	
	
	for _,value := range person{
	fmt.Println("Name : ",value.Name)
	fmt.Println("Roll : ",value.Roll)
	for i , value := range person{
	fmt.Printf("Subject : %s  Marks : %d",value.sub[i],value.marks[i])

	if value.marks[i] <=100 && value.marks[i] >= 80 {
		fmt.Println(" Grade : A+")
	}else if value.marks[i]<=79 && value.marks[i] >= 70{
		fmt.Println(" Grade : A")
	}else if value.marks[i]<=69 && value.marks[i]>=60 {
        fmt.Println(" Grade : A-")
	}else if value.marks[i]<=59 && value.marks[i]>=50 {
        fmt.Println(" Grade : B")
	}else if value.marks[i]<=49 && value.marks[i]>=40 {
        fmt.Println(" Grade : C")
	}else{
		fmt.Println(" Grade : F")
	}
	}
	
	fmt.Println("---------------")}

	
}
