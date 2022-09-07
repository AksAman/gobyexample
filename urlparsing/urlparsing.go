package urlparsing

import (
	"fmt"
	"net/url"
)

func Run() {
	rawURL := "postgres://user:pass@host.com:5432/path?k=v&u=1#first"
	showURLInfo(rawURL)

	rawURL = "https://gobyexample.com/url-parsing"
	showURLInfo(rawURL)
}

func showURLInfo(rawURL string) {
	fmt.Println("\n----- using url:", rawURL)
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("u.Scheme: %v\n", u.Scheme)
	fmt.Printf("u.User.Username(): %v\n", u.User.Username())

	password, hasPassword := u.User.Password()
	fmt.Printf("password: %v, hasPassword: %v\n", password, hasPassword)

	fmt.Printf("u.Fragment: %v\n", u.Fragment)
	fmt.Printf("u.Host: %v\n", u.Host)

	fmt.Printf("u.Path: %v\n", u.Path)

	fmt.Printf("u.Query(): %v\n", u.Query())
	fmt.Printf("u.RawQuery: %v\n", u.RawQuery)
	fmt.Println(url.ParseQuery(u.RawQuery))

}
