package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	q := getQuestion()
	bytArr := sendReq(q)
	r := processResponse(bytArr)
	fmt.Println("response", r)
}
func getQuestion() string {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a question...")
	q, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	question := strings.TrimSpace(q)
	return question
}
func sendReq(question string) []byte {
	baseUrl := "https://8ball.delegator.com/magic/JSON/"
	url := fmt.Sprintf("%s%s", baseUrl, question)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	bodyByt, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return bodyByt
}
func processResponse(b []byte) Response {
	var body Response
	err := json.Unmarshal(b, &body)
	if err != nil {
		log.Println(err)
	}
	return body
}
