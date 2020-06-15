package main

import "fmt"

type Serve struct {
	Name  string `json: "name"`
	Class string `json: "class"`
}

func main() {
	WriteFile()
	s := "admin"
	serves := readFile()
	for i := 0; i < len(serves); i++ {
		fmt.Println(i)
		fmt.Println("Name: ", serves[i].Name)
		fmt.Println("Class: ", serves[i].Class)
		fmt.Println("__")
	}

	for i := 0; i < len(serves); i++ {
		fmt.Println("__")
		checkString(s, serves[i])
	}

	item := Serve{Name: "fileCustome", Class: "org.cofax.cds.FileServlet.Custome"}
	serves = append(serves, item)

	fmt.Println("____")
	fmt.Println("After append")
	for i := 0; i < len(serves); i++ {
		fmt.Println(i)
		fmt.Println("Name: ", serves[i].Name)
		fmt.Println("Class: ", serves[i].Class)
		fmt.Println("__")
		fmt.Printf("address of serves[%d] is %p  \n", i, &serves[i])
	}
}
