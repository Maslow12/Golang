package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduce el numero ")
	text, _ := reader.ReadString('\n')
	_input := strings.TrimSpace(text)
	number, _ := strconv.Atoi(_input)

	for i := 1; i < number+1; i++ {
		if i%2 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz", i)
		} else if i%2 == 0 {
			fmt.Println("fizz", i)
		} else if i%3 == 0 {
			fmt.Println("buzz", i)
		}
	}
	return
}
