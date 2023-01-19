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
	fmt.Print("Enter the number: ")
	text, _ := reader.ReadString('\n')
	_text := strings.TrimSpace(text)
	number, err := strconv.Atoi(_text)
	println(number, err)
	return
}
