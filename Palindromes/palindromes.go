package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var phrase_invert []string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the phrase: ")
	text, _ := reader.ReadString('\n')
	_input := strings.TrimSpace(text)

	for i := len(_input) - 1; i > -1; i-- {
		phrase_invert = append(phrase_invert[:], string(_input[i]))
	}
	string_phrase := strings.Join(phrase_invert, "")

	if _input == string_phrase {
		fmt.Println("Esto es un palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}
