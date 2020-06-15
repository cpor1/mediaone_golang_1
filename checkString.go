package main

import (
	"fmt"
	"strings"
)

func checkString(s string, serve Serve) {
	nameLow := strings.ToLower(serve.Name)
	classLow := strings.ToLower(serve.Class)
	if strings.Contains(nameLow, s) || strings.Contains(classLow, s) {
		fmt.Println("Serve contain: ", s)
		fmt.Println("Name: ", serve.Name)
		fmt.Println("Class: ", serve.Class)
	}
}
