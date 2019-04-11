package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	words := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for _, a := range words {
		for _, b := range words {
			fmt.Printf("%s%s\n%s %s\n", a, b, a, b)
		}
	}
}
