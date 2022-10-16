package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	getQuestion()
}
func getQuestion() {
	r := bufio.NewReader(os.Stdin)
	q, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	question := strings.TrimSpace(q)
	fmt.Println("Question: ", question)
}
