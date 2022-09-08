package httpclients

import (
	"bufio"
	"fmt"
	"net/http"
)

func Run() {
	rawUrl := "https://gobyexample.com"
	resp, err := http.Get(rawUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("resp.Status: %v\n", resp.Status)
	fmt.Printf("resp.StatusCode: %v\n", resp.StatusCode)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Printf("scanner.Text(): %v\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
