package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	q := getQuestion()
	sendReq(q)
}
func getQuestion() string {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a question...")
	q, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	question := strings.TrimSpace(q)
	// fmt.Println("Question: ", question)
	return question
}
func sendReq(q string) {
	baseUrl := "https://8ball.delegator.com/magic/JSON/"
	url := fmt.Sprintf("%s%s", baseUrl, q)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Response: ", res)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Body: ", body)
	strBody := string(body)
	fmt.Println(strBody)
}
