package main

import "fmt"

func main() {
	grade := 60
	if grade >= 90 {
		fmt.Println("A grade")
	} else if grade >= 80 {
		fmt.Println("B grade")
	} else if grade >= 70 {
		fmt.Println("C grade")
	} else if grade >= 65 {
		fmt.Println("D grade")
	} else {
		fmt.Println("Failing grade")
	}
}
